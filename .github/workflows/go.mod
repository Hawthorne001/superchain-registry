module github.com/scripttoken/script

require (
	github.com/aerospike/aerospike-client-go v1.36.0
	github.com/bgentry/speakeasy v0.1.0
	github.com/davecgh/go-spew v1.1.1
	github.com/dgraph-io/badger v1.6.0-rc1
	github.com/fd/go-nat v1.0.0
	github.com/golang/protobuf v1.3.1
	github.com/golang/snappy v0.0.0-20180518054509-2e65f85255db
	github.com/gorilla/context v1.1.1 // indirect
	github.com/gorilla/mux v1.6.2
	github.com/hashicorp/golang-lru v0.5.1
	github.com/herumi/bls-eth-go-binary v0.0.0-20200107021104-147ed25f233e
	github.com/huin/goupnp v1.0.0
	github.com/influxdata/influxdb v1.7.8 // indirect
	github.com/ipfs/go-datastore v0.0.5
	github.com/ipfs/go-ipfs-addr v0.0.1
	github.com/jackpal/gateway v1.0.5
	github.com/jackpal/go-nat-pmp v1.0.1
	github.com/karalabe/hid v0.0.0-20180420081245-2b4488a37358
	github.com/koron/go-ssdp v0.0.0-20180514024734-4a0ed625a78b
	github.com/libp2p/go-libp2p v0.3.0
	github.com/libp2p/go-libp2p-connmgr v0.1.1
	github.com/libp2p/go-libp2p-core v0.2.0
	github.com/libp2p/go-libp2p-crypto v0.1.0
	github.com/libp2p/go-libp2p-kad-dht v0.2.0
	github.com/libp2p/go-libp2p-peerstore v0.1.3
	github.com/libp2p/go-libp2p-pubsub v0.1.1
	github.com/libp2p/go-libp2p-swarm v0.2.0
	github.com/libp2p/go-libp2p-transport v0.1.0
	github.com/libp2p/go-nat v0.0.3
	github.com/libp2p/go-stream-muxer v0.1.0
	github.com/mattn/go-isatty v0.0.12
	github.com/mitchellh/go-homedir v1.1.0
	github.com/mongodb/mongo-go-driver v0.0.17
	github.com/multiformats/go-multiaddr v0.0.4
	github.com/pborman/uuid v0.0.0-20180906182336-adf5a7427709
	github.com/phoreproject/bls v0.0.0-20191016230924-b2e57acce2ed
	github.com/pion/datachannel v1.4.13
	github.com/pion/webrtc/v2 v2.1.12
	github.com/pkg/errors v0.8.1
	github.com/pkg/profile v1.4.0
	github.com/prysmaticlabs/prysm v0.0.0-20191018160938-a05dca18c7f7
	github.com/russross/blackfriday v2.0.0+incompatible // indirect
	github.com/scripttoken/script/common v0.0.0
	github.com/scripttoken/script/rpc/lib/rpc-codec/jsonrpc2 v0.0.0
	github.com/shurcooL/sanitized_anchor_name v1.0.0 // indirect
	github.com/sirupsen/logrus v1.4.2
	github.com/smira/go-statsd v1.3.1
	github.com/spf13/cobra v0.0.5
	github.com/spf13/viper v1.5.0
	github.com/stretchr/testify v1.4.0
	github.com/syndtr/goleveldb v1.0.0
	github.com/tidwall/pretty v1.0.0 // indirect
	github.com/wedeploy/gosocketio v0.0.7-beta
	github.com/xdg/scram v0.0.0-20180814205039-7eeb5667e42c // indirect
	github.com/xdg/stringprep v0.0.0-20180714160509-73f8eece6fdc // indirect
	github.com/ybbus/jsonrpc v1.1.1
	github.com/yuin/gopher-lua v0.0.0-20180827083657-b942cacc89fe // indirect
	go.opencensus.io v0.21.0
	golang.org/x/crypto v0.0.0-20191001170739-f9e2070545dc
	golang.org/x/net v0.0.0-20191021144547-ec77196f6094
	golang.org/x/sys v0.0.0-20220412071739-889880a91fd5
	golang.org/x/xerrors v0.0.0-20191011141410-1b5146add898 // indirect
	gopkg.in/karalabe/cookiejar.v2 v2.0.0-20150724131613-8dcd6a7f4951
	gopkg.in/mgo.v2 v2.0.0-20180705113604-9856a29383ce
)

replace github.com/scripttoken/script/rpc/lib/rpc-codec/jsonrpc2 v0.0.0 => ./rpc/lib/rpc-codec/jsonrpc2/

replace github.com/scripttoken/script/common v0.0.0 => ./common

go 1.13
