package crypto

import (
	"encoding/ascii85"
	"errors"
	"github.com/agl/ed25519"
	"io/ioutil"
)

const (
	Ascii85SignatureSize = 80
	SignatureExtension   = ".sig"
)

func readSignature(in string) (signature []byte, err error) {
	a85signature, err := ioutil.ReadFile(in + SignatureExtension)
	if err != nil {
		return
	}
	signature = make([]byte, ed25519.SignatureSize)
	ascii85.Decode(signature, a85signature, true)
	return
}

func writeSignature(out string, signature []byte) (err error) {
	a85signature := make([]byte, Ascii85SignatureSize)
	ascii85.Encode(a85signature, signature)
	return ioutil.WriteFile(out+SignatureExtension, a85signature, defaultPermission)
}

func SignFile(keyfile, file string) (err error) {
	message, err := ioutil.ReadFile(file)
	if err != nil {
		return
	}
	key, err := ReadKey(keyfile)
	if err != nil {
		return
	}
	signature, err := Sign(key, message)
	if err != nil {
		return
	}
	return writeSignature(file, signature)
}

func Sign(key, message []byte) (signature []byte, err error) {
	var vKey [ed25519.PrivateKeySize]byte
	copy(vKey[:], key)
	vSignature := ed25519.Sign(&vKey, message)
	signature = vSignature[:]
	return
}

func VerifyFile(keyfile, file string) (err error) {
	message, err := ioutil.ReadFile(file)
	if err != nil {
		return
	}
	signature, err := readSignature(file)
	if err != nil {
		return
	}
	key, err := ReadPublicKey(keyfile)
	if err != nil {
		return
	}
	return Verify(key, message, signature)
}

func Verify(key, message, signature []byte) (err error) {
	var (
		vKey       [ed25519.PublicKeySize]byte
		vSignature [ed25519.SignatureSize]byte
	)
	copy(vKey[:], key)
	copy(vSignature[:], signature)
	if !ed25519.Verify(&vKey, message, &vSignature) {
		err = errors.New("Signature not valid")
	}
	return
}
