package crypto

import (
	"crypto/rand"
	"encoding/ascii85"
	"github.com/agl/ed25519"
	"io/ioutil"
	"strings"
)

const (
	Ascii85PublicKeySize  = 40
	Ascii85PrivateKeySize = 80
	PublicKeyExtension    = ".pub"
)

func ReadKey(in string) (key []byte, err error) {
	a85key, err := ioutil.ReadFile(in)
	if err != nil {
		return
	}
	size := ed25519.PrivateKeySize
	if isPublicKey(in) {
		size = ed25519.PublicKeySize
	}
	key = make([]byte, size)
	ascii85.Decode(key, a85key, true)
	return
}

func isPublicKey(filename string) bool {
	return strings.HasSuffix(filename, PublicKeyExtension)
}

func writeKey(out string, key []byte) (err error) {
	size := Ascii85PrivateKeySize
	if isPublicKey(out) {
		size = Ascii85PublicKeySize
	}
	a85key := make([]byte, size)
	ascii85.Encode(a85key, key)
	return ioutil.WriteFile(out, a85key, defaultPermission)
}

func ReadPublicKey(in string) (key []byte, err error) {
	return ReadKey(in + PublicKeyExtension)
}

func writePublicKey(out string, key []byte) (err error) {
	return writeKey(out+PublicKeyExtension, key)
}

func GenerateKey(out string) (err error) {
	pubKey, key, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return
	}
	if err = writeKey(out, key[:]); err != nil {
		return
	}
	if err = writePublicKey(out, pubKey[:]); err != nil {
		return
	}
	return
}
