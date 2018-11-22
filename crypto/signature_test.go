package crypto

import (
	"fmt"
	"testing"

	"github.com/btcsuite/btcd/btcec"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

func TestMain(t *testing.T) {
	fmt.Println(btcec.S256())
	fmt.Println(S256())
	key, err := GenerateKey()
	fmt.Println(key.Curve)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	hash := signHash([]byte("zaakin"))
	sig_ecdsa, err := Sign_ECDSA(hash, key)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	sig_eth, err := Sign_ETH(hash, key)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	sig_go, err := Sign_GO(hash, key)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	fmt.Println("sig_ecdsa:" + hexutil.Encode(sig_ecdsa))
	fmt.Println("sig_eth:" + hexutil.Encode(sig_eth))
	fmt.Println("sig_go:" + hexutil.Encode(sig_go))
}
