package rawdbreset

import (
	"context"
	"fmt"

	"github.com/ledgerwatch/erigon-lib/chain"
	"github.com/ledgerwatch/erigon-lib/common/datadir"
	"github.com/ledgerwatch/erigon-lib/common/hexutility"
	"github.com/ledgerwatch/erigon-lib/kv"
	"github.com/ledgerwatch/erigon-lib/kv/kvcfg"
	"github.com/ledgerwatch/erigon-lib/state"
	"github.com/ledgerwatch/erigon/consensus"
	"github.com/ledgerwatch/erigon/core"
	"github.com/ledgerwatch/erigon/core/rawdb"
	"github.com/ledgerwatch/erigon/eth/stagedsync"
	"github.com/ledgerwatch/erigon/eth/stagedsync/stages"
	"github.com/ledgerwatch/erigon/turbo/backup"
	"github.com/ledgerwatch/erigon/turbo/services"
	"github.com/ledgerwatch/erigon/turbo/snapshotsync"
	"github.com/ledgerwatch/log/v3"
)

func ResetState(db kv.RwDB, ctx context.Context, chain string, tmpDir string) error {
	// don't reset senders here
	if err := Reset(ctx, db, stages.HashState); err != nil {
		return err
	}
	if err := Reset(ctx, db, stages.IntermediateHashes); err != nil {
		return err
	}
	if err := Reset(ctx, db, stages.AccountHistoryIndex, stages.StorageHistoryIndex); err != nil {
		return err
	}
	if err := Reset(ctx, db, stages.LogIndex); err != nil {
		return err
	}
	if err := Reset(ctx, db, stages.CallTraces); err != nil {
		return err
	}
	if err := db.Update(ctx, ResetTxLookup); err != nil {
		return err
	}
	if err := Reset(ctx, db, stages.Finish); err != nil {
		return err
	}

	if err := ResetExec(ctx, db, chain, tmpDir); err != nil {
		return err
	}
	return nil
}

func ResetBlocks(tx kv.RwTx, db kv.RoDB, snapshots *snapshotsync.RoSnapshots, agg *state.AggregatorV3,
	br services.FullBlockReader, dirs datadir.Dirs, cc chain.Config, engine consensus.Engine, logger log.Logger) error {
	// keep Genesis
	if err := rawdb.TruncateBlocks(context.Background(), tx, 1); err != nil {
		return err
	}
	if err := stages.SaveStageProgress(tx, stages.Bodies, 1); err != nil {
		return fmt.Errorf("saving Bodies progress failed: %w", err)
	}
	if err := stages.SaveStageProgress(tx, stages.Headers, 1); err != nil {
		return fmt.Errorf("saving Bodies progress failed: %w", err)
	}
	if err := stages.SaveStageProgress(tx, stages.Snapshots, 0); err != nil {
		return fmt.Errorf("saving Snapshots progress failed: %w", err)
	}

	// remove all canonical markers from this point
	if err := rawdb.TruncateCanonicalHash(tx, 1, false); err != nil {
		return err
	}
	if err := rawdb.TruncateTd(tx, 1); err != nil {
		return err
	}
	hash, err := rawdb.ReadCanonicalHash(tx, 0)
	if err != nil {
		return err
	}
	if err = rawdb.WriteHeadHeaderHash(tx, hash); err != nil {
		return err
	}

	// ensure no garbage records left (it may happen if db is inconsistent)
	if err := tx.ForEach(kv.BlockBody, hexutility.EncodeTs(2), func(k, _ []byte) error { return tx.Delete(kv.BlockBody, k) }); err != nil {
		return err
	}
	ethtx := kv.EthTx
	transactionV3, err := kvcfg.TransactionsV3.Enabled(tx)
	if err != nil {
		panic(err)
	}
	if transactionV3 {
		ethtx = kv.EthTxV3
	}

	if err := backup.ClearTables(context.Background(), db, tx,
		kv.NonCanonicalTxs,
		ethtx,
		kv.MaxTxNum,
	); err != nil {
		return err
	}
	if err := rawdb.ResetSequence(tx, ethtx, 0); err != nil {
		return err
	}
	if err := rawdb.ResetSequence(tx, kv.NonCanonicalTxs, 0); err != nil {
		return err
	}

	if snapshots != nil && snapshots.Cfg().Enabled && snapshots.BlocksAvailable() > 0 {
		if err := stagedsync.FillDBFromSnapshots("fillind_db_from_snapshots", context.Background(), tx, dirs, snapshots, br, cc, engine, agg, logger); err != nil {
			return err
		}
		_ = stages.SaveStageProgress(tx, stages.Snapshots, snapshots.BlocksAvailable())
		_ = stages.SaveStageProgress(tx, stages.Headers, snapshots.BlocksAvailable())
		_ = stages.SaveStageProgress(tx, stages.Bodies, snapshots.BlocksAvailable())
		_ = stages.SaveStageProgress(tx, stages.Senders, snapshots.BlocksAvailable())
	}

	return nil
}
func ResetSenders(ctx context.Context, db kv.RwDB, tx kv.RwTx) error {
	if err := backup.ClearTables(ctx, db, tx, kv.Senders); err != nil {
		return nil
	}
	return clearStageProgress(tx, stages.Senders)
}

func WarmupExec(ctx context.Context, db kv.RwDB) (err error) {
	for _, tbl := range stateBuckets {
		backup.WarmupTable(ctx, db, tbl, log.LvlInfo, backup.ReadAheadThreads)
	}
	historyV3 := kvcfg.HistoryV3.FromDB(db)
	if historyV3 { //hist v2 is too big, if you have so much ram, just use `cat mdbx.dat > /dev/null` to warmup
		for _, tbl := range stateHistoryV3Buckets {
			backup.WarmupTable(ctx, db, tbl, log.LvlInfo, backup.ReadAheadThreads)
		}
	}
	return
}

func ResetExec(ctx context.Context, db kv.RwDB, chain string, tmpDir string) (err error) {
	historyV3 := kvcfg.HistoryV3.FromDB(db)
	if historyV3 {
		stateHistoryBuckets = append(stateHistoryBuckets, stateHistoryV3Buckets...)
		stateHistoryBuckets = append(stateHistoryBuckets, stateHistoryV4Buckets...)
	}

	return db.Update(ctx, func(tx kv.RwTx) error {
		if err := clearStageProgress(tx, stages.Execution, stages.HashState, stages.IntermediateHashes); err != nil {
			return err
		}

		if err := backup.ClearTables(ctx, db, tx, stateBuckets...); err != nil {
			return nil
		}
		for _, b := range stateBuckets {
			if err := tx.ClearBucket(b); err != nil {
				return err
			}
		}

		if err := backup.ClearTables(ctx, db, tx, stateHistoryBuckets...); err != nil {
			return nil
		}
		if !historyV3 {
			genesis := core.GenesisBlockByChainName(chain)
			if _, _, err := core.WriteGenesisState(genesis, tx, tmpDir); err != nil {
				return err
			}
		}

		return nil
	})
}

func ResetTxLookup(tx kv.RwTx) error {
	if err := tx.ClearBucket(kv.TxLookup); err != nil {
		return err
	}
	if err := stages.SaveStageProgress(tx, stages.TxLookup, 0); err != nil {
		return err
	}
	if err := stages.SaveStagePruneProgress(tx, stages.TxLookup, 0); err != nil {
		return err
	}
	return nil
}

var Tables = map[stages.SyncStage][]string{
	stages.HashState:           {kv.HashedAccounts, kv.HashedStorage, kv.ContractCode},
	stages.IntermediateHashes:  {kv.TrieOfAccounts, kv.TrieOfStorage},
	stages.CallTraces:          {kv.CallFromIndex, kv.CallToIndex},
	stages.LogIndex:            {kv.LogAddressIndex, kv.LogTopicIndex},
	stages.AccountHistoryIndex: {kv.AccountsHistory},
	stages.StorageHistoryIndex: {kv.StorageHistory},
	stages.Finish:              {},
}
var stateBuckets = []string{
	kv.PlainState, kv.HashedAccounts, kv.HashedStorage, kv.TrieOfAccounts, kv.TrieOfStorage,
	kv.Epoch, kv.PendingEpoch, kv.BorReceipts,
	kv.Code, kv.PlainContractCode, kv.ContractCode, kv.IncarnationMap,
}
var stateHistoryBuckets = []string{
	kv.AccountChangeSet,
	kv.StorageChangeSet,
	kv.Receipts,
	kv.Log,
	kv.CallTraceSet,
}
var stateHistoryV3Buckets = []string{
	kv.AccountHistoryKeys, kv.AccountIdx, kv.AccountHistoryVals,
	kv.StorageKeys, kv.StorageVals, kv.StorageHistoryKeys, kv.StorageHistoryVals, kv.StorageIdx,
	kv.CodeKeys, kv.CodeVals, kv.CodeHistoryKeys, kv.CodeHistoryVals, kv.CodeIdx,
	kv.AccountHistoryKeys, kv.AccountIdx, kv.AccountHistoryVals,
	kv.StorageHistoryKeys, kv.StorageIdx, kv.StorageHistoryVals,
	kv.CodeHistoryKeys, kv.CodeIdx, kv.CodeHistoryVals,
	kv.LogAddressKeys, kv.LogAddressIdx,
	kv.LogTopicsKeys, kv.LogTopicsIdx,
	kv.TracesFromKeys, kv.TracesFromIdx,
	kv.TracesToKeys, kv.TracesToIdx,
}
var stateHistoryV4Buckets = []string{
	kv.AccountKeys, kv.StorageKeys, kv.CodeKeys,
	kv.CommitmentKeys, kv.CommitmentVals, kv.CommitmentHistoryKeys, kv.CommitmentHistoryVals, kv.CommitmentIdx,
}

func clearStageProgress(tx kv.RwTx, stagesList ...stages.SyncStage) error {
	for _, stage := range stagesList {
		if err := stages.SaveStageProgress(tx, stage, 0); err != nil {
			return err
		}
		if err := stages.SaveStagePruneProgress(tx, stage, 0); err != nil {
			return err
		}
	}
	return nil
}

func Reset(ctx context.Context, db kv.RwDB, stagesList ...stages.SyncStage) error {
	return db.Update(ctx, func(tx kv.RwTx) error {
		for _, st := range stagesList {
			if err := backup.ClearTables(ctx, db, tx, Tables[st]...); err != nil {
				return err
			}
			if err := clearStageProgress(tx, stagesList...); err != nil {
				return err
			}
		}
		return nil
	})
}
func Warmup(ctx context.Context, db kv.RwDB, lvl log.Lvl, stList ...stages.SyncStage) error {
	for _, st := range stList {
		for _, tbl := range Tables[st] {
			backup.WarmupTable(ctx, db, tbl, lvl, backup.ReadAheadThreads)
		}
	}
	return nil
}
