// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

import (
	libcommon "github.com/ledgerwatch/erigon-lib/common"

	"github.com/ledgerwatch/erigon/params/networkname"
)

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Ethereum network.
var MainnetBootnodes = []string{
	// Ethereum Foundation Go Bootnodes
	"enode://d860a01f9722d78051619d1e2351aba3f43f943f6f00718d1b9baa4101932a1f5011f16bb2b1bb35db20d6fe28fa0bf09636d26a87d31de9ec6203eeedb1f666@18.138.108.67:30303", // bootnode-aws-ap-southeast-1-001
	"enode://22a8232c3abc76a16ae9d6c3b164f98775fe226f0917b0ca871128a74a8e9630b458460865bab457221f1d448dd9791d24c4e5d88786180ac185df813a68d4de@3.209.45.79:30303",   // bootnode-aws-us-east-1-001
	"enode://2b252ab6a1d0f971d9722cb839a42cb81db019ba44c08754628ab4a823487071b5695317c8ccd085219c3a03af063495b2f1da8d18218da2d6a82981b45e6ffc@65.108.70.101:30303", // bootnode-hetzner-hel
	"enode://4aeb4ab6c14b23e2c4cfdce879c04b0748a20d8e9b59e25ded2a08143e265c6c25936e74cbc8e641e3312ca288673d91f2f93f8e277de3cfa444ecdaaf982052@157.90.35.166:30303", // bootnode-hetzner-fsn
}

// SepoliaBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Sepolia test network.
var SepoliaBootnodes = []string{
	// EF DevOps
	"enode://4e5e92199ee224a01932a377160aa432f31d0b351f84ab413a8e0a42f4f36476f8fb1cbe914af0d9aef0d51665c214cf653c651c4bbd9d5550a934f241f1682b@138.197.51.181:30303", // sepolia-bootnode-1-nyc3
	"enode://143e11fb766781d22d92a2e33f8f104cddae4411a122295ed1fdb6638de96a6ce65f5b7c964ba3763bba27961738fef7d3ecc739268f3e5e771fb4c87b6234ba@146.190.1.103:30303",  // sepolia-bootnode-1-sfo3
	"enode://8b61dc2d06c3f96fddcbebb0efb29d60d3598650275dc469c22229d3e5620369b0d3dedafd929835fe7f489618f19f456fe7c0df572bf2d914a9f4e006f783a9@170.64.250.88:30303",  // sepolia-bootnode-1-syd1
	"enode://10d62eff032205fcef19497f35ca8477bea0eadfff6d769a147e895d8b2b8f8ae6341630c645c30f5df6e67547c03494ced3d9c5764e8622a26587b083b028e8@139.59.49.206:30303",  // sepolia-bootnode-1-blr1
	"enode://9e9492e2e8836114cc75f5b929784f4f46c324ad01daf87d956f98b3b6c5fcba95524d6e5cf9861dc96a2c8a171ea7105bb554a197455058de185fa870970c7c@138.68.123.152:30303", // sepolia-bootnode-1-ams3
}

// RinkebyBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network.
var RinkebyBootnodes = []string{
	"enode://a24ac7c5484ef4ed0c5eb2d36620ba4e4aa13b8c84684e1b4aab0cebea2ae45cb4d375b77eab56516d34bfbd3c1a833fc51296ff084b770b94fb9028c4d25ccf@52.169.42.101:30303", // IE
	"enode://343149e4feefa15d882d9fe4ac7d88f885bd05ebb735e547f12e12080a9fa07c8014ca6fd7f373123488102fe5e34111f8509cf0b7de3f5b44339c9f25e87cb8@52.3.158.184:30303",  // INFURA
	"enode://b6b28890b006743680c52e64e0d16db57f28124885595fa03a562be1d2bf0f3a1da297d56b13da25fb992888fd556d4c1a27b1f39d531bde7de1921c90061cc6@159.89.28.211:30303", // AKASHA
}

// GoerliBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Görli test network.
var GoerliBootnodes = []string{
	// Upstream bootnodes
	"enode://011f758e6552d105183b1761c5e2dea0111bc20fd5f6422bc7f91e0fabbec9a6595caf6239b37feb773dddd3f87240d99d859431891e4a642cf2a0a9e6cbb98a@51.141.78.53:30303",
	"enode://176b9417f511d05b6b2cf3e34b756cf0a7096b3094572a8f6ef4cdcb9d1f9d00683bf0f83347eebdf3b81c3521c2332086d9592802230bf528eaf606a1d9677b@13.93.54.137:30303",
	"enode://46add44b9f13965f7b9875ac6b85f016f341012d84f975377573800a863526f4da19ae2c620ec73d11591fa9510e992ecc03ad0751f53cc02f7c7ed6d55c7291@94.237.54.114:30313",
	"enode://b5948a2d3e9d486c4d75bf32713221c2bd6cf86463302339299bd227dc2e276cd5a1c7ca4f43a0e9122fe9af884efed563bd2a1fd28661f3b5f5ad7bf1de5949@18.218.250.66:30303",

	// Ethereum Foundation bootnode
	"enode://a61215641fb8714a373c80edbfa0ea8878243193f57c96eeb44d0bc019ef295abd4e044fd619bfc4c59731a73fb79afe84e9ab6da0c743ceb479cbb6d263fa91@3.11.147.67:30303",

	// Goerli Initiative bootnodes
	"enode://d4f764a48ec2a8ecf883735776fdefe0a3949eb0ca476bd7bc8d0954a9defe8fea15ae5da7d40b5d2d59ce9524a99daedadf6da6283fca492cc80b53689fb3b3@46.4.99.122:32109",
	"enode://d2b720352e8216c9efc470091aa91ddafc53e222b32780f505c817ceef69e01d5b0b0797b69db254c586f493872352f5a022b4d8479a00fc92ec55f9ad46a27e@88.99.70.182:30303",
}

var SepoliaStaticPeers = []string{
	// from https://github.com/ledgerwatch/erigon/issues/6134#issuecomment-1354923418
	"enode://8ae4559db1b1e160be8cc46018d7db123ed6d03fbbfe481da5ec05f71f0aa4d5f4b02ad059127096aa994568706a0d02933984083b87c5e1e3de2b7692444d37@35.161.233.158:46855",
	"enode://d0b3b290422f35ec3e68356f3a4cdf9c661f71a868110670e31441a5021d7abd0440ae8dfb9360aafdd0198f177863361e3a7a7eb5e1a3e26575bf1ac3ef4ab3@162.19.136.65:48264",
	"enode://d64624bda3cdb65d542c90757a4a661cfe9dddf8328bdb1ea97a8d70fad287c360f0101c492d8fd6ab30d79160a3bf148cacfd68f5d2e47eab0b709516419304@51.195.63.10:30040",
	"enode://c7df835939e027325c6bba926220fae5912a33c83d96b3eef8ef445c98083f3191788581c9a0e8f74cadb0b13229b847f5c1ebd315b22bcf11faf6468020eb48@54.163.51.157:30303",
	"enode://da0609bad3afcab9b93175a41a2d621d07aa7ff6c134a00792d4541f0ce8d30d8f3c51bb37a47573508a0bf18865b04066af2a661edf1d3a3d8d133fc1031aa0@88.151.101.14:45192",
	"enode://7a4534d392c59369eae6befa56ac670476d9edc16597cf53c92bbefa6e741b6b0b9e6822cab12afb09123e03ca1131026fbef145adec429fe2e50182dfb650a5@94.130.18.108:31312",
	"enode://db6fa13b63a885440de581ee3fc8df9c6a590326b39fc5ccba7991707ee0cebac306211f7eca5270a350201a3132511f2338481edd81f3dc819c2a1c60419cf2@65.21.89.157:30303",
	"enode://fcf03e9404cace34c60e4eed374ef9a779471014319b3346352fbc2f992a399af6517486e8e65a4ab55f4645fe55420bbea1cddc13a4af4df63b0f731915c6a6@13.125.238.49:46173",
	"enode://8b973816278fdd56966709e4794c7ccce1f256eaa9165a6b013b991a9bdf3886a8f2d23af50ee723a5614a9fe9d197252b803b4455a87ab2468e128f7b06e0ca@172.104.107.145:30303",
	"enode://5a1fb15f826a213d3ef4adb9be47ab58b2240ea05df0d760a244f04762b0847dcb08276b1284f726c22eea30fce0c601cf121b81bac0c151f1b3b4ad00d1482a@34.159.55.147:51262",
	"enode://560928dd14819f88113586726e452b16bbc694ed4144ddadd6290053e7f3fc66bfad13add6889f7d8f37e0c21ccbb6948eb8899c8b30743f4b45a3081f1efed8@34.138.254.5:29888",
	"enode://69a13b575b8c5278431409e9f7db36e7218667ae286bfb65a72dfec9201b2c5bbbe2797a1babbdf17a7bf7ca68fa3fbe1554612637eb1b2425fa975e1bccb54c@35.223.41.3:30303",
	"enode://66158b31eecff939f220b291d2b448edbfe94f1d4c992d9395b5d476e55e54b5abd11d3ee44daf1e18ee27b910ef99cdf6f19775eb4820ebe4f77d7aa948e3b6@51.195.63.10:55198",
	"enode://bf94acbd51170bf075cacb9f149b21ff46354d659ab434a0d40688f776e1e1556bc62be2dc2867ba513844268c0dc8240099a6b60efe1713fbc25da7fdeb6ff1@3.82.105.139:30303",
	"enode://41329e5ceb51cdddbe6a475db00b682505768b71ff8ee37d2d3500ca1b78918f9fad57d6006dd9f79cd418437dbcf87ec2fd58d60710f925cb17da05a51197cf@65.21.34.60:30303",
	"enode://4e5e92199ee224a01932a377160aa432f31d0b351f84ab413a8e0a42f4f36476f8fb1cbe914af0d9aef0d51665c214cf653c651c4bbd9d5550a934f241f1682b@138.197.51.181:30303",
	"enode://143e11fb766781d22d92a2e33f8f104cddae4411a122295ed1fdb6638de96a6ce65f5b7c964ba3763bba27961738fef7d3ecc739268f3e5e771fb4c87b6234ba@146.190.1.103:30303",
	"enode://8b61dc2d06c3f96fddcbebb0efb29d60d3598650275dc469c22229d3e5620369b0d3dedafd929835fe7f489618f19f456fe7c0df572bf2d914a9f4e006f783a9@170.64.250.88:30303",
	"enode://10d62eff032205fcef19497f35ca8477bea0eadfff6d769a147e895d8b2b8f8ae6341630c645c30f5df6e67547c03494ced3d9c5764e8622a26587b083b028e8@139.59.49.206:30303",
	"enode://9e9492e2e8836114cc75f5b929784f4f46c324ad01daf87d956f98b3b6c5fcba95524d6e5cf9861dc96a2c8a171ea7105bb554a197455058de185fa870970c7c@138.68.123.152:30303",
}

var V5Bootnodes = []string{
	// Teku team's bootnode
	"enr:-KG4QOtcP9X1FbIMOe17QNMKqDxCpm14jcX5tiOE4_TyMrFqbmhPZHK_ZPG2Gxb1GE2xdtodOfx9-cgvNtxnRyHEmC0ghGV0aDKQ9aX9QgAAAAD__________4JpZIJ2NIJpcIQDE8KdiXNlY3AyNTZrMaEDhpehBDbZjM_L9ek699Y7vhUJ-eAdMyQW_Fil522Y0fODdGNwgiMog3VkcIIjKA",
	"enr:-KG4QDyytgmE4f7AnvW-ZaUOIi9i79qX4JwjRAiXBZCU65wOfBu-3Nb5I7b_Rmg3KCOcZM_C3y5pg7EBU5XGrcLTduQEhGV0aDKQ9aX9QgAAAAD__________4JpZIJ2NIJpcIQ2_DUbiXNlY3AyNTZrMaEDKnz_-ps3UUOfHWVYaskI5kWYO_vtYMGYCQRAR3gHDouDdGNwgiMog3VkcIIjKA",
	// Prylab team's bootnodes
	"enr:-Ku4QImhMc1z8yCiNJ1TyUxdcfNucje3BGwEHzodEZUan8PherEo4sF7pPHPSIB1NNuSg5fZy7qFsjmUKs2ea1Whi0EBh2F0dG5ldHOIAAAAAAAAAACEZXRoMpD1pf1CAAAAAP__________gmlkgnY0gmlwhBLf22SJc2VjcDI1NmsxoQOVphkDqal4QzPMksc5wnpuC3gvSC8AfbFOnZY_On34wIN1ZHCCIyg",
	"enr:-Ku4QP2xDnEtUXIjzJ_DhlCRN9SN99RYQPJL92TMlSv7U5C1YnYLjwOQHgZIUXw6c-BvRg2Yc2QsZxxoS_pPRVe0yK8Bh2F0dG5ldHOIAAAAAAAAAACEZXRoMpD1pf1CAAAAAP__________gmlkgnY0gmlwhBLf22SJc2VjcDI1NmsxoQMeFF5GrS7UZpAH2Ly84aLK-TyvH-dRo0JM1i8yygH50YN1ZHCCJxA",
	"enr:-Ku4QPp9z1W4tAO8Ber_NQierYaOStqhDqQdOPY3bB3jDgkjcbk6YrEnVYIiCBbTxuar3CzS528d2iE7TdJsrL-dEKoBh2F0dG5ldHOIAAAAAAAAAACEZXRoMpD1pf1CAAAAAP__________gmlkgnY0gmlwhBLf22SJc2VjcDI1NmsxoQMw5fqqkw2hHC4F5HZZDPsNmPdB1Gi8JPQK7pRc9XHh-oN1ZHCCKvg",
	// Lighthouse team's bootnodes
	"enr:-IS4QLkKqDMy_ExrpOEWa59NiClemOnor-krjp4qoeZwIw2QduPC-q7Kz4u1IOWf3DDbdxqQIgC4fejavBOuUPy-HE4BgmlkgnY0gmlwhCLzAHqJc2VjcDI1NmsxoQLQSJfEAHZApkm5edTCZ_4qps_1k_ub2CxHFxi-gr2JMIN1ZHCCIyg",
	"enr:-IS4QDAyibHCzYZmIYZCjXwU9BqpotWmv2BsFlIq1V31BwDDMJPFEbox1ijT5c2Ou3kvieOKejxuaCqIcjxBjJ_3j_cBgmlkgnY0gmlwhAMaHiCJc2VjcDI1NmsxoQJIdpj_foZ02MXz4It8xKD7yUHTBx7lVFn3oeRP21KRV4N1ZHCCIyg",
	// EF bootnodes
	"enr:-Ku4QHqVeJ8PPICcWk1vSn_XcSkjOkNiTg6Fmii5j6vUQgvzMc9L1goFnLKgXqBJspJjIsB91LTOleFmyWWrFVATGngBh2F0dG5ldHOIAAAAAAAAAACEZXRoMpC1MD8qAAAAAP__________gmlkgnY0gmlwhAMRHkWJc2VjcDI1NmsxoQKLVXFOhp2uX6jeT0DvvDpPcU8FWMjQdR4wMuORMhpX24N1ZHCCIyg",
	"enr:-Ku4QG-2_Md3sZIAUebGYT6g0SMskIml77l6yR-M_JXc-UdNHCmHQeOiMLbylPejyJsdAPsTHJyjJB2sYGDLe0dn8uYBh2F0dG5ldHOIAAAAAAAAAACEZXRoMpC1MD8qAAAAAP__________gmlkgnY0gmlwhBLY-NyJc2VjcDI1NmsxoQORcM6e19T1T9gi7jxEZjk_sjVLGFscUNqAY9obgZaxbIN1ZHCCIyg",
	"enr:-Ku4QPn5eVhcoF1opaFEvg1b6JNFD2rqVkHQ8HApOKK61OIcIXD127bKWgAtbwI7pnxx6cDyk_nI88TrZKQaGMZj0q0Bh2F0dG5ldHOIAAAAAAAAAACEZXRoMpC1MD8qAAAAAP__________gmlkgnY0gmlwhDayLMaJc2VjcDI1NmsxoQK2sBOLGcUb4AwuYzFuAVCaNHA-dy24UuEKkeFNgCVCsIN1ZHCCIyg",
	"enr:-Ku4QEWzdnVtXc2Q0ZVigfCGggOVB2Vc1ZCPEc6j21NIFLODSJbvNaef1g4PxhPwl_3kax86YPheFUSLXPRs98vvYsoBh2F0dG5ldHOIAAAAAAAAAACEZXRoMpC1MD8qAAAAAP__________gmlkgnY0gmlwhDZBrP2Jc2VjcDI1NmsxoQM6jr8Rb1ktLEsVcKAPa08wCsKUmvoQ8khiOl_SLozf9IN1ZHCCIyg",
}

var MumbaiBootnodes = []string{
	"enode://320553cda00dfc003f499a3ce9598029f364fbb3ed1222fdc20a94d97dcc4d8ba0cd0bfa996579dcc6d17a534741fb0a5da303a90579431259150de66b597251@54.147.31.250:30303",
	"enode://f0f48a8781629f95ff02606081e6e43e4aebd503f3d07fc931fad7dd5ca1ba52bd849a6f6c3be0e375cf13c9ae04d859c4a9ae3546dc8ed4f10aa5dbb47d4998@34.226.134.117:30303",
}

var BorMainnetBootnodes = []string{
	"enode://0cb82b395094ee4a2915e9714894627de9ed8498fb881cec6db7c65e8b9a5bd7f2f25cc84e71e89d0947e51c76e85d0847de848c7782b13c0255247a6758178c@44.232.55.71:30303",
	"enode://88116f4295f5a31538ae409e4d44ad40d22e44ee9342869e7d68bdec55b0f83c1530355ce8b41fbec0928a7d75a5745d528450d30aec92066ab6ba1ee351d710@159.203.9.164:30303",
}

var GnosisBootnodes = []string{
	"enode://6765fff89db92aa8d923e28c438af626c8ae95a43093cdccbd6f550a7b6ce6ab5d1a3dc60dd79af3e6d2c2e6731bae629f0e54446a0d9da408c4eca7ebcd8485@3.75.159.31:30303",
	"enode://9a7c98e8ee8cdd3199db68092b48868847d4743a471b26afc2ff878bafaa829ed43ee405f9aff58ae13fce53b898f7c2e3c30cb80af8eb111682c3c13f686dbb@18.198.130.54:30303",
	"enode://2c4307831914c237801993eac4f9596d8b2f78e1e76830419b64cb23f0933e52cb1e2bb3009cb4af76454bb5bc296135b36869fd6c13e2c2e536a0780e60fe82@3.64.242.196:30303",
	"enode://074f68e1a7df5b0859314ff721d55b59d9690e93249c941660609a29b302f02864df4f93ee48884f7ede57dc7f7646379d017a43c9745e34baff049749896b50@3.126.169.151:30303",
	"enode://d239697375d7586c7ea1de790401c310b0b1d389326849fa3b7c7005833c7a6b9020e49dfb3b61abfa39135237ffc4ff219cb84ca7653069e8548497527aa432@107.22.4.120:30303",
	"enode://d5852bf415d89b756faa809f4ff3f8beb661dc7d60cfb4a5542f9a5fcdf41e1ed0708a210db64b8c7ca32426e04ef0a50da58974124fdf562a8510314d11e28c@3.26.206.142:30303",
	"enode://01d372392bb22dd8a91f8b10b6bbb8d80d2dbe98d695801e0df9e4bd4825781df84bba88361f24d1b6580a61313f64e6cec82e8d842ad5f1b3d7cf8d6d132da7@15.152.45.82:30303",
	"enode://aee88e803b8e54925081957965b2527961cd90f4d6d14664884580b429da44729678a1258a8b49a42d1582c9c7c5ded05733622f7ab442ad9c6f655545a5ecdd@54.207.220.169:30303",
}

var ChiadoBootnodes = []string{
	"enode://7dd44af6138120f328bb031eb56e00985c149319d4f1e33275b30be7fddadd8ccd9f7b9c3b35a16136a61e85b2b2d1de073f30ec1d0ddf576a33be8ff48d88d0@139.144.26.89:30303",
	"enode://317b9cee65ccf1d747b00e604242bfa3ae367beee8f149e28c5b2b88820f855ea7b5a75eb5327cfc3d8ca97adbf71538468290a46592ed7009f3fb394ec752f1@139.144.26.115:30303",
	"enode://b77ae97906155ebbb83fd32c87ab0aa57372a24abbd8aa4bae679f048b726de4a195709f613be4981e44b24640bc89e4824427d94e9a37afc148da8250c8ab2d@139.144.26.101:30303",
	"enode://69f8abfa3b0221161f8c19014b90857a18742554af27af73fd779c486728750a0ff11b873975f104fc5276a3a7c3b5b68cb3c26c815e9f78462901895d652124@139.144.26.85:30303",
	"enode://ac7fc76f9b2ab343fb2d091365a7f46d17018e525cbedfbf24b247c76657e934ef4df61cc2f6dad6bfcf722425e03e1a8a6e4e4b52743acc2319cb8ebf27d742@170.187.154.239:30303",
	"enode://712144ac396fd2298b3e2559e2930d7f3a36fded3addd66955224958f1845634067717ab9522757ed2948f480fc52add5676487c8378e9011a7e2c0ac2f36cc3@3.71.132.231:30303",
	"enode://595160631241ea41b187b85716f9f9572a266daa940d74edbe3b83477264ce284d69208e61cf50e91641b1b4f9a03fa8e60eb73d435a84cf4616b1c969bc2512@3.69.35.13:30303",
}

const dnsPrefix = "enrtree://AKA3AM6LPBYEUDMVNU3BSVQJ5AD45Y7YPOHJLEF6W26QOE4VTUDPE@"

// KnownDNSNetwork returns the address of a public DNS-based node list for the given
// genesis hash and protocol. See https://github.com/ethereum/discv4-dns-lists for more
// information.
func KnownDNSNetwork(genesis libcommon.Hash, protocol string) string {
	var net string
	switch genesis {
	case MainnetGenesisHash:
		net = "mainnet"
	case RinkebyGenesisHash:
		net = "rinkeby"
	case GoerliGenesisHash:
		net = "goerli"
	default:
		return ""
	}
	return dnsPrefix + protocol + "." + net + ".ethdisco.net"
}

func BootnodeURLsOfChain(chain string) []string {
	switch chain {
	case networkname.MainnetChainName:
		return MainnetBootnodes
	case networkname.SepoliaChainName:
		return SepoliaBootnodes
	case networkname.RinkebyChainName:
		return RinkebyBootnodes
	case networkname.GoerliChainName:
		return GoerliBootnodes
	case networkname.MumbaiChainName:
		return MumbaiBootnodes
	case networkname.BorMainnetChainName:
		return BorMainnetBootnodes
	case networkname.GnosisChainName:
		return GnosisBootnodes
	case networkname.ChiadoChainName:
		return ChiadoBootnodes
	default:
		return []string{}
	}
}

func StaticPeerURLsOfChain(chain string) []string {
	switch chain {
	case networkname.SepoliaChainName:
		return SepoliaStaticPeers
	default:
		return []string{}
	}
}
