package jsontokens

import (
	"bytes"
	"fmt"
	"testing"
)

var base64TestData = []struct {
	decoded     string
	encoded     string
	escaped_dec string
	escaped_enc string
}{
	{
		"{\"alg\":\"HS256\",\"typ\":\"JWT\"}",
		"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9",
		"\"{ \"alg\": \"HS256\", \"typ\": \"JWT\"}\"",
		"\"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9\"",
	},
	{
		`{
         "sub": "1234567890",
         "name": "John Doe",
         "iat": 1516239022
        }`,
		"eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ",
		`"{
         "sub": "1234567890",
         "name": "John Doe",
         "iat": 1516239022
        }"`,
		`"eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ"`,
	},
	{
		`{"alg":"HS256","typ":"JWT"}`,
		"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9",
		`"{"alg": "HS256","typ": "JWT"}"`,
		`"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"`,
	},
	{
		`{"key":"value"}`,
		"eyJrZXkiOiJ2YWx1ZSJ9",
		`"{"key": "value"}"`,
		`"eyJrZXkiOiJ2YWx1ZSJ9"`,
	},
}

func TestBase64(t *testing.T) {
	for _, data := range base64TestData {
		encoded := Base64Encode([]byte(data.decoded))
		decoded, err := Base64Decode([]byte(data.encoded))
		if err != nil {
			fmt.Println(err)
			t.Fail()
		}

		escaped_enc := EncodeEscape([]byte(data.escaped_dec))
		escaped_dec, err := DecodeEscaped([]byte(data.escaped_enc))
		if err != nil {
			fmt.Println(err)
			t.Fail()
		}

		if !bytes.Equal(encoded, []byte(data.encoded)) {
			fmt.Printf("Base64Encode is %s\n", string(encoded))
			fmt.Printf("Base64Decode is %s\n\n", string(decoded))
			t.Fail()
		}

		if !bytes.Equal(escaped_dec, []byte(data.decoded)) {
			fmt.Printf("EncodeEscape is %s\n", string(escaped_enc))
			fmt.Printf("DecodeEscaped is %s\n\n", string(escaped_dec))
			t.Fail()
		}
	}
}
