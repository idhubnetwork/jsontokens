package did

import (
	"errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var urls map[string]string
var address = ""
var net = "infuraRopsten"

// var err error

func init() {
	urls = make(map[string]string, 5)
	urls["infuraMainnet"] = "https://mainnet.infura.io"
	urls["infuraRopsten"] = "https://ropsten.infura.io"
	urls["infuraRinkeby"] = "https://rinkeby.infura.io"
}

func defaultDid() (*Did, error) {
	client, err := ethclient.Dial(urls[net])
	if err != nil {
		return nil, err
	}
	contractAddr := common.HexToAddress(address)
	contract, err := NewDid(contractAddr, client)
	if err != nil {
		return nil, err
	}

	return contract, nil
}

func GetDid(s ...string) (*Did, error) {
	switch len(s) {
	case 0:
		return defaultDid()
	case 1:
		if len(s[0]) == 42 {
			address = s[0]
		}
		if len(s[0]) != 42 {
			net = s[0]
		}
	default:
		net = s[0]
		address = s[1]
	}
	net_url, ok := urls[net]
	if !ok {
		return nil, errors.New("invalid ethereum net key")
	}
	client, err := ethclient.Dial(net_url)
	if err != nil {
		return nil, err
	}
	contractAddr := common.HexToAddress(address)
	contract, err := NewDid(contractAddr, client)
	if err != nil {
		return nil, err
	}

	return contract, nil
}
