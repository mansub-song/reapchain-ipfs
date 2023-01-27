package encrypt

var key []byte

func SetKey(keyByte []byte) {
	key = keyByte
}

func GetKey() []byte {
	return key
}

// func SetNonce(nonceByte []byte) {
// 	nonce = nonceByte
// }

// func GetNonce() []byte {
// 	return nonce
// }
