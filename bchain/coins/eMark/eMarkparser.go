package eMark

import (
	"github.com/trezor/blockbook/bchain"
	"github.com/trezor/blockbook/bchain/coins/btc"


	"github.com/martinboehm/btcd/wire"
	"github.com/martinboehm/btcutil/chaincfg"
)

// magic numbers
const (
	MainnetMagic wire.BitcoinNet = 0xe5e9e8e4
)

// chain parameters
var (
	MainNetParams chaincfg.Params
)

func init() {
	MainNetParams = chaincfg.MainNetParams
	MainNetParams.Net = MainnetMagic
	MainNetParams.PubKeyHashAddrID = []byte{53}
	MainNetParams.ScriptHashAddrID = []byte{30}
}

// eMarkParser handle
type eMarkParser struct {
	*btc.BitcoinParser
	baseparser *bchain.BaseParser
}

// NeweMarkParser returns new eMarkParser instance
func NeweMarkParser(params *chaincfg.Params, c *btc.Configuration) *eMarkParser {
	return &eMarkParser{
		BitcoinParser: btc.NewBitcoinParser(params, c),
		baseparser:    &bchain.BaseParser{},
	}
}

// GetChainParams contains network parameters for the main eMark network,
// and the test eMark network
func GetChainParams(chain string) *chaincfg.Params {
	if !chaincfg.IsRegistered(&MainNetParams) {
		err := chaincfg.Register(&MainNetParams)
		if err != nil {
			panic(err)
		}
	}
	switch chain {
	default:
		return &MainNetParams
	}
}

// PackTx packs transaction to byte array using protobuf
func (p *eMarkParser) PackTx(tx *bchain.Tx, height uint32, blockTime int64) ([]byte, error) {
	return p.baseparser.PackTx(tx, height, blockTime)
}

// UnpackTx unpacks transaction from protobuf byte array
func (p *eMarkParser) UnpackTx(buf []byte) (*bchain.Tx, uint32, error) {
	return p.baseparser.UnpackTx(buf)
}
