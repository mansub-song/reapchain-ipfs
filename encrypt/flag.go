package encrypt

import (
	"bytes"
	"io"
)

var encryptionFlag = false
var decryptionFlag = false
var keyByte []byte
var nonce []byte

func SetEncryptionFlag(flag bool) {
	encryptionFlag = flag
}

func GetEncryptionFlag() bool {
	return encryptionFlag
}

func SetDecryptionFlag(flag bool) {
	decryptionFlag = flag
}

func GetDecryptionFlag() bool {
	return decryptionFlag
}

func StreamToByte(stream io.Reader) []byte {
	buf := new(bytes.Buffer)
	buf.ReadFrom(stream)
	return buf.Bytes()
}
