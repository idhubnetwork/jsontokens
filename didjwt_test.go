package jsontokens

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"
)

var did_pri_key = "0x06da20d5a2ffdeb3f5b6bca5199a28d63a6efec4b1b8b9c6c493fc532ed324ec"

var jwtTest = "eyJhbGciOiJFUzI1NmsiLCJ0eXAiOiJKV1QifQ.eyJhdWQiOiJkaWQ6aWRodWI6MHg0OWRCYThmOTA2Yzc0NUIwYTgyZjREMjFFMDJCQUZEN0RmMWEwYmU0IiwiZXhwIjoxNTQ5NDMxMTEwLCJpcGZzIjoiSVBGUzppZGh1YlRlc3QiLCJpc3MiOiJkaWQ6aWRodWI6MHg3RWJFRTlhOEEzNTMwZmQxZTU0MDE3QzM5NTkyQTVhOTVhZjk5ZDA3IiwianRpIjo1ODgyNSwibmV0IjoiIGV0aF9yb3BzdGVuIiwic3RhdHVzIjoxMCwic3ViIjoxfQ.9IgAbcwtV2SkTGwbmg-YkSMUyb49eh5ZA0lmm8uhXQc8SjkuTM32meF3hamDiQR2U-S14eg6M3lYcgJ2c9xvvBs"

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
		jwt.Set("sub", strconv.Itoa(i))
		jwt.Set("aud", "did:idhub:0x49dBa8f906c745B0a82f4D21E02BAFD7Df1a0be4")
		jwt.Set("exp", time.Now().Unix()+int64(i*5000000))
		jwt.Set("jti", strconv.Itoa(i*58825))
		jwt.Set("net", "eth_ropsten")
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
	jwt.Set("iss", "did:idhub:0x7EbEE9a8A3530fd1e54017C39592A5a95af99d07")

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
