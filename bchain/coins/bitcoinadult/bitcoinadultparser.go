package bitcoinadult

import (
	"github.com/stepollo2/blockbook/bchain"
	"github.com/stepollo2/blockbook/bchain/coins/btc"

	"github.com/martinboehm/btcd/wire"
	"github.com/martinboehm/btcutil/chaincfg"
)

const (
	MainnetMagic wire.BitcoinNet = 0xd8b2a5c8
)

var (
	MainNetParams chaincfg.Params
)

func init() {
	MainNetParams = chaincfg.MainNetParams
	MainNetParams.Net = MainnetMagic

	MainNetParams.PubKeyHashAddrID = []byte{25}
	MainNetParams.ScriptHashAddrID = []byte{30}
}

type BitcoinadultParser struct {
	*btc.BitcoinParser
	baseparser *bchain.BaseParser
}

func NewBitcoinadultParser(params *chaincfg.Params, c *btc.Configuration) *BitcoinadultParser {
	return &BitcoinadultParser{BitcoinParser: btc.NewBitcoinParser(params, c), baseparser: &bchain.BaseParser{}}
}

func GetChainParams(chain string) *chaincfg.Params {
	if !chaincfg.IsRegistered(&MainNetParams) {
		err := chaincfg.Register(&MainNetParams)
		if err != nil {
			panic(err)
		}
	}
	return &MainNetParams
}

func (p *BitcoinadultParser) PackTx(tx *bchain.Tx, height uint32, blockTime int64) ([]byte, error) {
	return p.baseparser.PackTx(tx, height, blockTime)
}

func (p *BitcoinadultParser) UnpackTx(buf []byte) (*bchain.Tx, uint32, error) {
	return p.baseparser.UnpackTx(buf)
}