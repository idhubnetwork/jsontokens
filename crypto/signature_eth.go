// Copyright 2018 The idhub_2vid Authors
// This file is part of the jsontokens library.
//

package crypto

import (
	"crypto/ecdsa"
	"errors"
	"fmt"

	"github.com/btcsuite/btcd/btcec"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// Sign_ECDSA calculates an ECDSA signature.
//
// This function is susceptible to chosen plaintext attacks that can leak
// information about the private key that is used for signing. Callers must
// be aware that the given hash cannot be chosen by an adversery. Common
// solution is to hash any input before calculating the signature.
//
// The produced signature is in the [R || S || V] format where V is 0 or 1.
func Sign_ECDSA(hash []byte, key interface{}) ([]byte, error) {
	prv, ok := key.(*ecdsa.PrivateKey)
	if !ok {
		return nil, errors.New("key is invalid")
	}

	if len(hash) != 32 {
		return nil, fmt.Errorf("hash is required to be exactly 32 bytes (%d)", len(hash))
	}

	if prv.Curve != btcec.S256() {
		return nil, fmt.Errorf("private key curve is not secp256k1")
	}

	sig, err := btcec.SignCompact(btcec.S256(), (*btcec.PrivateKey)(prv), hash, false)
	if err != nil {
		return nil, err
	}
	// Convert to Ethereum signature format with 'recovery id' v at the end.
	v := sig[0] - 27
	copy(sig, sig[1:])
	sig[64] = v
	return sig, nil
}

// Sign_ETH calculates an Ethereum ECDSA signature for:
// keccack256("\x19Ethereum Signed Message:\n" + len(message) + message))
//
// Note, the produced signature conforms to the secp256k1 curve R, S and V values,
// where the V value will be 27 or 28 for legacy reasons.
//
// The key used to calculate the signature is decrypted with the given password.
//
// https://github.com/ethereum/go-ethereum/wiki/Management-APIs#personal_sign
func Sign_ETH(privateKey, msg string) ([]byte, error) {
	data := []byte(msg)
	hash := SignHash(data)
	fmt.Println("hash already")
	prv, err := HexToECDSA(privateKey)
	if err != nil {
		return nil, err
	}
	fmt.Println("prv already")

	if len(hash) != 32 {
		return nil, fmt.Errorf("hash is required to be exactly 32 bytes (%d)", len(hash))
	}
	/*
		if prv.Curve != btcec.S256() {
			return nil, fmt.Errorf("private key curve is not secp256k1")
		}
	*/

	sig, err := btcec.SignCompact(btcec.S256(), (*btcec.PrivateKey)(prv), hash, false)
	if err != nil {
		return nil, err
	}
	// V is 27/28 according to the yellow paper
	v := sig[0]
	copy(sig, sig[1:])
	sig[64] = v
	return sig, nil
}

// Sign calculates an hex Ethereum ECDSA signature for:
// keccack256("\x19Ethereum Signed Message:\n" + len(message) + message))
// with a hex private key.
//
// The V value is 27/28 according to the yellow paper.
func Sign(privateKey, msg string) (sigHex string, err error) {
	signature, err := Sign_ETH(privateKey, msg)
	if err != nil {
		return "", err
	}
	fmt.Println("signature already")
	sigHex = hexutil.Encode(signature)
	return sigHex, nil
}

// Ecrecover returns the uncompressed public key that created the given signature.
func Ecrecover(hash, sig []byte) (publicKey []byte, err error) {
	pub, err := SigToPub(hash, sig)
	if err != nil {
		return nil, err
	}
	bytes := (*btcec.PublicKey)(pub).SerializeUncompressed()
	return bytes, err
}

// EcRecover returns the address for the account that was used to create the signature.
// Note, this function is compatible with eth_sign and personal_sign. As such it recovers
// the address of:
// hash = keccak256("\x19Ethereum Signed Message:\n"${message length}${message})
// addr = ecrecover(hash, signature)
//
// Note, the signature must conform to the secp256k1 curve R, S and V values, where
// the V value must be be 27 or 28 for legacy reasons.
//
// https://github.com/ethereum/go-ethereum/wiki/Management-APIs#personal_ecRecover
func EcRecover(msg, sigHex string) (addr string, err error) {
	defer func() {
		if r := recover(); r != nil {
			addr = common.Address{}.String()
			err = r.(error)
		}
	}()

	data := []byte(msg)
	hash := SignHash(data)

	sig := hexutil.MustDecode(sigHex)
	// https://github.com/ethereum/go-ethereum/blob/55599ee95d4151a2502465e0afc7c47bd1acba77/internal/ethapi/api.go#L442
	if sig[64] != 27 && sig[64] != 28 {
		return common.Address{}.String(), errors.New("nvalid Ethereum signature (V is not 27 or 28)")
	}
	sig[64] -= 27

	sigPiblicKey, err := SigToPub(hash, sig)

	if err != nil {
		panic(err)
	}

	commAddr := PubkeyToAddress(*sigPiblicKey)

	return commAddr.String(), nil
}
