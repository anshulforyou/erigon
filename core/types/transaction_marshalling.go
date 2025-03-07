package types

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/holiman/uint256"
	. "github.com/protolambda/ztyp/view"
	"github.com/valyala/fastjson"

	libcommon "github.com/ledgerwatch/erigon-lib/common"
	"github.com/ledgerwatch/erigon-lib/common/hexutility"
	types2 "github.com/ledgerwatch/erigon-lib/types"

	"github.com/ledgerwatch/erigon/common/hexutil"
)

// txJSON is the JSON representation of transactions.
type txJSON struct {
	Type hexutil.Uint64 `json:"type"`

	// Common transaction fields:
	Nonce    *hexutil.Uint64    `json:"nonce"`
	GasPrice *hexutil.Big       `json:"gasPrice"`
	FeeCap   *hexutil.Big       `json:"maxFeePerGas"`
	Tip      *hexutil.Big       `json:"maxPriorityFeePerGas"`
	Gas      *hexutil.Uint64    `json:"gas"`
	Value    *hexutil.Big       `json:"value"`
	Data     *hexutility.Bytes  `json:"input"`
	V        *hexutil.Big       `json:"v"`
	R        *hexutil.Big       `json:"r"`
	S        *hexutil.Big       `json:"s"`
	To       *libcommon.Address `json:"to"`

	// Access list transaction fields:
	ChainID    *hexutil.Big       `json:"chainId,omitempty"`
	AccessList *types2.AccessList `json:"accessList,omitempty"`

	// Blob transaction fields:
	MaxFeePerDataGas    *hexutil.Big     `json:"maxFeePerDataGas,omitempty"`
	BlobVersionedHashes []libcommon.Hash `json:"blobVersionedHashes,omitempty"`
	// Blob wrapper fields:
	Blobs       Blobs     `json:"blobs,omitempty"`
	Commitments BlobKzgs  `json:"commitments,omitempty"`
	Proofs      KZGProofs `json:"proofs,omitempty"`

	// Only used for encoding:
	Hash libcommon.Hash `json:"hash"`
}

func (tx LegacyTx) MarshalJSON() ([]byte, error) {
	var enc txJSON
	// These are set for all tx types.
	enc.Hash = tx.Hash()
	enc.Type = hexutil.Uint64(tx.Type())
	enc.Nonce = (*hexutil.Uint64)(&tx.Nonce)
	enc.Gas = (*hexutil.Uint64)(&tx.Gas)
	enc.GasPrice = (*hexutil.Big)(tx.GasPrice.ToBig())
	enc.Value = (*hexutil.Big)(tx.Value.ToBig())
	enc.Data = (*hexutility.Bytes)(&tx.Data)
	enc.To = tx.To
	enc.V = (*hexutil.Big)(tx.V.ToBig())
	enc.R = (*hexutil.Big)(tx.R.ToBig())
	enc.S = (*hexutil.Big)(tx.S.ToBig())
	return json.Marshal(&enc)
}

func (tx AccessListTx) MarshalJSON() ([]byte, error) {
	var enc txJSON
	// These are set for all tx types.
	enc.Hash = tx.Hash()
	enc.Type = hexutil.Uint64(tx.Type())
	enc.ChainID = (*hexutil.Big)(tx.ChainID.ToBig())
	enc.AccessList = &tx.AccessList
	enc.Nonce = (*hexutil.Uint64)(&tx.Nonce)
	enc.Gas = (*hexutil.Uint64)(&tx.Gas)
	enc.GasPrice = (*hexutil.Big)(tx.GasPrice.ToBig())
	enc.Value = (*hexutil.Big)(tx.Value.ToBig())
	enc.Data = (*hexutility.Bytes)(&tx.Data)
	enc.To = tx.To
	enc.V = (*hexutil.Big)(tx.V.ToBig())
	enc.R = (*hexutil.Big)(tx.R.ToBig())
	enc.S = (*hexutil.Big)(tx.S.ToBig())
	return json.Marshal(&enc)
}

func (tx DynamicFeeTransaction) MarshalJSON() ([]byte, error) {
	var enc txJSON
	// These are set for all tx types.
	enc.Hash = tx.Hash()
	enc.Type = hexutil.Uint64(tx.Type())
	enc.ChainID = (*hexutil.Big)(tx.ChainID.ToBig())
	enc.AccessList = &tx.AccessList
	enc.Nonce = (*hexutil.Uint64)(&tx.Nonce)
	enc.Gas = (*hexutil.Uint64)(&tx.Gas)
	enc.FeeCap = (*hexutil.Big)(tx.FeeCap.ToBig())
	enc.Tip = (*hexutil.Big)(tx.Tip.ToBig())
	enc.Value = (*hexutil.Big)(tx.Value.ToBig())
	enc.Data = (*hexutility.Bytes)(&tx.Data)
	enc.To = tx.To
	enc.V = (*hexutil.Big)(tx.V.ToBig())
	enc.R = (*hexutil.Big)(tx.R.ToBig())
	enc.S = (*hexutil.Big)(tx.S.ToBig())
	return json.Marshal(&enc)
}

func toSignedBlobTxJSON(tx *SignedBlobTx) *txJSON {
	var enc txJSON
	// These are set for all tx types.
	enc.Hash = tx.Hash()
	enc.Type = hexutil.Uint64(tx.Type())
	enc.ChainID = (*hexutil.Big)(tx.GetChainID().ToBig())
	accessList := tx.GetAccessList()
	enc.AccessList = &accessList
	nonce := tx.GetNonce()
	enc.Nonce = (*hexutil.Uint64)(&nonce)
	gas := tx.GetGas()
	enc.Gas = (*hexutil.Uint64)(&gas)
	enc.FeeCap = (*hexutil.Big)(tx.GetFeeCap().ToBig())
	enc.Tip = (*hexutil.Big)(tx.GetTip().ToBig())
	enc.Value = (*hexutil.Big)(tx.GetValue().ToBig())
	enc.Data = (*hexutility.Bytes)(&tx.Message.Data)
	enc.To = tx.GetTo()
	enc.V = (*hexutil.Big)(tx.Signature.GetV().ToBig())
	enc.R = (*hexutil.Big)(tx.Signature.GetR().ToBig())
	enc.S = (*hexutil.Big)(tx.Signature.GetS().ToBig())
	enc.MaxFeePerDataGas = (*hexutil.Big)(tx.GetMaxFeePerDataGas().ToBig())
	enc.BlobVersionedHashes = tx.GetDataHashes()
	return &enc
}

func (tx SignedBlobTx) MarshalJSON() ([]byte, error) {
	return json.Marshal(toSignedBlobTxJSON(&tx))
}

func (tx BlobTxWrapper) MarshalJSON() ([]byte, error) {
	enc := toSignedBlobTxJSON(&tx.Tx)

	enc.Blobs = tx.Blobs
	enc.Commitments = tx.Commitments
	enc.Proofs = tx.Proofs

	return json.Marshal(enc)
}

func UnmarshalTransactionFromJSON(input []byte) (Transaction, error) {
	var p fastjson.Parser
	v, err := p.ParseBytes(input)
	if err != nil {
		return nil, fmt.Errorf("parse transaction json: %w", err)
	}
	// check the type
	txTypeHex := v.GetStringBytes("type")
	var txType hexutil.Uint64 = LegacyTxType
	if txTypeHex != nil {
		if err = txType.UnmarshalText(txTypeHex); err != nil {
			return nil, err
		}
	}
	switch byte(txType) {
	case LegacyTxType:
		tx := &LegacyTx{}
		if err = tx.UnmarshalJSON(input); err != nil {
			return nil, err
		}
		return tx, nil
	case AccessListTxType:
		tx := &AccessListTx{}
		if err = tx.UnmarshalJSON(input); err != nil {
			return nil, err
		}
		return tx, nil
	case DynamicFeeTxType:
		tx := &DynamicFeeTransaction{}
		if err = tx.UnmarshalJSON(input); err != nil {
			return nil, err
		}
		return tx, nil
	case BlobTxType:
		tx, err := UnmarshalBlobTxJSON(input)
		if err != nil {
			return nil, err
		}
		return tx, nil
	default:
		return nil, fmt.Errorf("unknown transaction type: %v", txType)
	}
}

func (tx *LegacyTx) UnmarshalJSON(input []byte) error {
	var dec txJSON
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.To != nil {
		tx.To = dec.To
	}
	if dec.Nonce == nil {
		return errors.New("missing required field 'nonce' in transaction")
	}
	tx.Nonce = uint64(*dec.Nonce)
	if dec.GasPrice == nil {
		return errors.New("missing required field 'gasPrice' in transaction")
	}
	var overflow bool
	tx.GasPrice, overflow = uint256.FromBig(dec.GasPrice.ToInt())
	if overflow {
		return errors.New("'gasPrice' in transaction does not fit in 256 bits")
	}
	if dec.Gas == nil {
		return errors.New("missing required field 'gas' in transaction")
	}
	tx.Gas = uint64(*dec.Gas)
	if dec.Value == nil {
		return errors.New("missing required field 'value' in transaction")
	}
	tx.Value, overflow = uint256.FromBig(dec.Value.ToInt())
	if overflow {
		return errors.New("'value' in transaction does not fit in 256 bits")
	}
	if dec.Data == nil {
		return errors.New("missing required field 'input' in transaction")
	}
	tx.Data = *dec.Data
	if dec.V == nil {
		return errors.New("missing required field 'v' in transaction")
	}
	overflow = tx.V.SetFromBig(dec.V.ToInt())
	if overflow {
		return fmt.Errorf("dec.V higher than 2^256-1")
	}
	if dec.R == nil {
		return errors.New("missing required field 'r' in transaction")
	}
	overflow = tx.R.SetFromBig(dec.R.ToInt())
	if overflow {
		return fmt.Errorf("dec.R higher than 2^256-1")
	}
	if dec.S == nil {
		return errors.New("missing required field 's' in transaction")
	}
	overflow = tx.S.SetFromBig(dec.S.ToInt())
	if overflow {
		return fmt.Errorf("dec.S higher than 2^256-1")
	}
	if overflow {
		return errors.New("'s' in transaction does not fit in 256 bits")
	}
	withSignature := !tx.V.IsZero() || !tx.R.IsZero() || !tx.S.IsZero()
	if withSignature {
		if err := sanityCheckSignature(&tx.V, &tx.R, &tx.S, true); err != nil {
			return err
		}
	}
	return nil
}

func (tx *AccessListTx) UnmarshalJSON(input []byte) error {
	var dec txJSON
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	// Access list is optional for now.
	if dec.AccessList != nil {
		tx.AccessList = *dec.AccessList
	}
	if dec.ChainID == nil {
		return errors.New("missing required field 'chainId' in transaction")
	}
	var overflow bool
	tx.ChainID, overflow = uint256.FromBig(dec.ChainID.ToInt())
	if overflow {
		return errors.New("'chainId' in transaction does not fit in 256 bits")
	}
	if dec.To != nil {
		tx.To = dec.To
	}
	if dec.Nonce == nil {
		return errors.New("missing required field 'nonce' in transaction")
	}
	tx.Nonce = uint64(*dec.Nonce)
	if dec.GasPrice == nil {
		return errors.New("missing required field 'gasPrice' in transaction")
	}
	tx.GasPrice, overflow = uint256.FromBig(dec.GasPrice.ToInt())
	if overflow {
		return errors.New("'gasPrice' in transaction does not fit in 256 bits")
	}
	if dec.Gas == nil {
		return errors.New("missing required field 'gas' in transaction")
	}
	tx.Gas = uint64(*dec.Gas)
	if dec.Value == nil {
		return errors.New("missing required field 'value' in transaction")
	}
	tx.Value, overflow = uint256.FromBig(dec.Value.ToInt())
	if overflow {
		return errors.New("'value' in transaction does not fit in 256 bits")
	}
	if dec.Data == nil {
		return errors.New("missing required field 'input' in transaction")
	}
	tx.Data = *dec.Data
	if dec.V == nil {
		return errors.New("missing required field 'v' in transaction")
	}
	overflow = tx.V.SetFromBig(dec.V.ToInt())
	if overflow {
		return fmt.Errorf("dec.V higher than 2^256-1")
	}
	if dec.R == nil {
		return errors.New("missing required field 'r' in transaction")
	}
	overflow = tx.R.SetFromBig(dec.R.ToInt())
	if overflow {
		return fmt.Errorf("dec.R higher than 2^256-1")
	}
	if dec.S == nil {
		return errors.New("missing required field 's' in transaction")
	}
	overflow = tx.S.SetFromBig(dec.S.ToInt())
	if overflow {
		return fmt.Errorf("dec.S higher than 2^256-1")
	}
	withSignature := !tx.V.IsZero() || !tx.R.IsZero() || !tx.S.IsZero()
	if withSignature {
		if err := sanityCheckSignature(&tx.V, &tx.R, &tx.S, false); err != nil {
			return err
		}
	}
	return nil
}

func (tx *DynamicFeeTransaction) UnmarshalJSON(input []byte) error {
	var dec txJSON
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	// Access list is optional for now.
	if dec.AccessList != nil {
		tx.AccessList = *dec.AccessList
	}
	if dec.ChainID == nil {
		return errors.New("missing required field 'chainId' in transaction")
	}
	var overflow bool
	tx.ChainID, overflow = uint256.FromBig(dec.ChainID.ToInt())
	if overflow {
		return errors.New("'chainId' in transaction does not fit in 256 bits")
	}
	if dec.To != nil {
		tx.To = dec.To
	}
	if dec.Nonce == nil {
		return errors.New("missing required field 'nonce' in transaction")
	}
	tx.Nonce = uint64(*dec.Nonce)
	if dec.GasPrice == nil {
		return errors.New("missing required field 'gasPrice' in transaction")
	}
	tx.Tip, overflow = uint256.FromBig(dec.Tip.ToInt())
	if overflow {
		return errors.New("'tip' in transaction does not fit in 256 bits")
	}
	tx.FeeCap, overflow = uint256.FromBig(dec.FeeCap.ToInt())
	if overflow {
		return errors.New("'feeCap' in transaction does not fit in 256 bits")
	}
	if dec.Gas == nil {
		return errors.New("missing required field 'gas' in transaction")
	}
	tx.Gas = uint64(*dec.Gas)
	if dec.Value == nil {
		return errors.New("missing required field 'value' in transaction")
	}
	tx.Value, overflow = uint256.FromBig(dec.Value.ToInt())
	if overflow {
		return errors.New("'value' in transaction does not fit in 256 bits")
	}
	if dec.Data == nil {
		return errors.New("missing required field 'input' in transaction")
	}
	tx.Data = *dec.Data
	if dec.V == nil {
		return errors.New("missing required field 'v' in transaction")
	}
	overflow = tx.V.SetFromBig(dec.V.ToInt())
	if overflow {
		return fmt.Errorf("dec.V higher than 2^256-1")
	}
	if dec.R == nil {
		return errors.New("missing required field 'r' in transaction")
	}
	overflow = tx.R.SetFromBig(dec.R.ToInt())
	if overflow {
		return fmt.Errorf("dec.R higher than 2^256-1")
	}
	if dec.S == nil {
		return errors.New("missing required field 's' in transaction")
	}
	overflow = tx.S.SetFromBig(dec.S.ToInt())
	if overflow {
		return fmt.Errorf("dec.S higher than 2^256-1")
	}
	if overflow {
		return errors.New("'s' in transaction does not fit in 256 bits")
	}
	withSignature := !tx.V.IsZero() || !tx.R.IsZero() || !tx.S.IsZero()
	if withSignature {
		if err := sanityCheckSignature(&tx.V, &tx.R, &tx.S, false); err != nil {
			return err
		}
	}
	return nil
}

func UnmarshalBlobTxJSON(input []byte) (Transaction, error) {
	var dec txJSON
	if err := json.Unmarshal(input, &dec); err != nil {
		return nil, err
	}
	tx := SignedBlobTx{}
	if dec.AccessList != nil {
		tx.Message.AccessList = AccessListView(*dec.AccessList)
	} else {
		tx.Message.AccessList = AccessListView([]types2.AccessTuple{})
	}
	if dec.ChainID == nil {
		return nil, errors.New("missing required field 'chainId' in transaction")
	}
	chainID, overflow := uint256.FromBig(dec.ChainID.ToInt())
	if overflow {
		return nil, errors.New("'chainId' in transaction does not fit in 256 bits")
	}
	tx.Message.ChainID = Uint256View(*chainID)
	if dec.To != nil {
		address := AddressSSZ(*dec.To)
		tx.Message.To = AddressOptionalSSZ{Address: &address}
	}
	if dec.Nonce == nil {
		return nil, errors.New("missing required field 'nonce' in transaction")
	}
	tx.Message.Nonce = Uint64View(uint64(*dec.Nonce))
	tip, overflow := uint256.FromBig(dec.Tip.ToInt())
	if overflow {
		return nil, errors.New("'tip' in transaction does not fit in 256 bits")
	}
	tx.Message.GasTipCap = Uint256View(*tip)
	feeCap, overflow := uint256.FromBig(dec.FeeCap.ToInt())
	if overflow {
		return nil, errors.New("'feeCap' in transaction does not fit in 256 bits")
	}
	tx.Message.GasFeeCap = Uint256View(*feeCap)
	if dec.Gas == nil {
		return nil, errors.New("missing required field 'gas' in transaction")
	}
	tx.Message.Gas = Uint64View(uint64(*dec.Gas))
	if dec.Value == nil {
		return nil, errors.New("missing required field 'value' in transaction")
	}
	value, overflow := uint256.FromBig(dec.Value.ToInt())
	if overflow {
		return nil, errors.New("'value' in transaction does not fit in 256 bits")
	}
	tx.Message.Value = Uint256View(*value)
	if dec.Data == nil {
		return nil, errors.New("missing required field 'input' in transaction")
	}
	tx.Message.Data = TxDataView(*dec.Data)

	if dec.MaxFeePerDataGas == nil {
		return nil, errors.New("missing required field 'maxFeePerDataGas' in transaction")
	}
	maxFeePerDataGas, overflow := uint256.FromBig(dec.MaxFeePerDataGas.ToInt())
	if overflow {
		return nil, errors.New("'maxFeePerDataGas' in transaction does not fit in 256 bits")
	}
	tx.Message.MaxFeePerDataGas = Uint256View(*maxFeePerDataGas)

	if dec.BlobVersionedHashes != nil {
		tx.Message.BlobVersionedHashes = VersionedHashesView(dec.BlobVersionedHashes)
	} else {
		tx.Message.BlobVersionedHashes = VersionedHashesView([]libcommon.Hash{})
	}

	if dec.V == nil {
		return nil, errors.New("missing required field 'v' in transaction")
	}
	var v uint256.Int
	overflow = v.SetFromBig(dec.V.ToInt())
	if overflow {
		return nil, fmt.Errorf("dec.V higher than 2^256-1")
	}
	if v.Uint64() > 255 {
		return nil, fmt.Errorf("dev.V higher than 2^8 - 1")
	}

	tx.Signature.V = Uint8View(v.Uint64())

	if dec.R == nil {
		return nil, errors.New("missing required field 'r' in transaction")
	}
	var r uint256.Int
	overflow = r.SetFromBig(dec.R.ToInt())
	if overflow {
		return nil, fmt.Errorf("dec.R higher than 2^256-1")
	}
	tx.Signature.R = Uint256View(r)

	if dec.S == nil {
		return nil, errors.New("missing required field 's' in transaction")
	}
	var s uint256.Int
	overflow = s.SetFromBig(dec.S.ToInt())
	if overflow {
		return nil, errors.New("'s' in transaction does not fit in 256 bits")
	}
	tx.Signature.S = Uint256View(s)

	withSignature := !v.IsZero() || !r.IsZero() || !s.IsZero()
	if withSignature {
		if err := sanityCheckSignature(&v, &r, &s, false); err != nil {
			return nil, err
		}
	}

	if len(dec.Blobs) == 0 {
		// if no blobs are specified in the json we assume it is an unwrapped blob tx
		return &tx, nil
	}

	btx := BlobTxWrapper{
		Tx:          tx,
		Commitments: dec.Commitments,
		Blobs:       dec.Blobs,
		Proofs:      dec.Proofs,
	}
	err := btx.ValidateBlobTransactionWrapper()
	if err != nil {
		return nil, err
	}
	return &btx, nil
}
