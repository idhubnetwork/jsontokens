package jsontokens

func DecodeEscaped(b []byte) ([]byte, error) {
	if len(b) > 1 && b[0] == '"' && b[len(b)-1] == '"' {
		b = b[1 : len(b)-1]
	}
	return b, nil
}
