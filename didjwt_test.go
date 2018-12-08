package jsontokens

import (
	"bufio"
	"fmt"
	"os"
	"testing"
	"time"
)

// did----0x7EbEE9a8A3530fd1e54017C39592A5a95af99d07
var did_pri_key = "0x06da20d5a2ffdeb3f5b6bca5199a28d63a6efec4b1b8b9c6c493fc532ed324ec"

// var privateKey = "0x6e365748e2a389106b24c241485a5308fb73548d43327c7a9fd4d972ca4cd472"
var jwtTest = "eyJhbGciOiJFUzI1NmsiLCJ0eXAiOiJKV1QifQ.eyJpc3MiOiJ0ZXN0IiwibmFtZSI6ImlkaHViIiwidHlwZSI6InRlc3QifQ.QT24XXc9iDKwfclvd1K-L2A7qBeXpZtXmfRodVPsS1V5SuK7CrQDwf0AHEQnU_BNg9aS4s17ESPeEuBg_rjv_xw"

func TestGetCredentials(t *testing.T) {
	outputFile, outputError := os.OpenFile("test_credential.json", os.O_WRONLY|os.O_CREATE, 0666)
	if outputError != nil {
		fmt.Printf("An error occurred with file opening or creation\n")
		return
	}
	defer outputFile.Close()

	outputWriter := bufio.NewWriter(outputFile)

	jwt := NewJWT()
	for i := 1; i < 13; i++ {
		jwt.Set("iss", "did:idhub:0x7EbEE9a8A3530fd1e54017C39592A5a95af99d07")
		jwt.Set("sub", i)
		jwt.Set("aud", "did:idhub:0x49dBa8f906c745B0a82f4D21E02BAFD7Df1a0be4")
		jwt.Set("exp", time.Now().Unix()+int64(i*5000000))
		jwt.Set("jti", i*58825)
		jwt.Set("net", " eth_ropsten")
		jwt.Set("ipfs", "IPFS:idhubTest")
		switch i % 4 {
		// 1111
		case 0:
			jwt.Set("status", 0xf)
		// 1010
		case 1:
			jwt.Set("status", 0xa)
		// 1100
		case 2:
			jwt.Set("status", 0xc)
		// 0011
		case 3:
			jwt.Set("status", 0x3)
		}
		err := jwt.Sign(did_pri_key)
		if err != nil {
			fmt.Println(err)
			t.Fail()
		}
		token, err := jwt.GetJWT()
		if err != nil {
			fmt.Println(err)
			t.Fail()
		}
		fmt.Printf("CREDENTIAL: %s\n", token)

		outputWriter.WriteString(token + "\n\n")
		jwt.Reset()
	}
	outputWriter.Flush()
}

func TestGetJWT(t *testing.T) {
	jwt := NewJWT()
	jwt.Set("name", "idhub")
	jwt.Set("type", "test")
	jwt.Set("iss", "test")

	fmt.Println(jwt)

	err := jwt.Sign(privateKey)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Println(jwt)
	token, err := jwt.GetJWT()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	fmt.Println(token)
}

func TestSetJWT(t *testing.T) {
	jwt := NewJWT()
	err := jwt.SetJWT(jwtTest)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	fmt.Println(jwt)
	fmt.Println("JWT SET Success")
}

func TestJWTVerify(t *testing.T) {
	jwt := NewJWT()
	err := jwt.SetJWT(jwtTest)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	err = jwt.Verify()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
}
