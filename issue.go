package identita

import (
	"github.com/imdario/identita/base40"
	"io/ioutil"
	"math/big"
)

func IssueFromFile(keyfile, file, password string) (id []byte, err error) {
	key, err := ReadKey(keyfile)
	if err != nil {
		return
	}
	pubKey, err := ReadPublicKey(keyfile)
	if err != nil {
		return
	}
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return
	}
	return Issue(key, pubKey, password, data)
}

func Issue(key, pubKey []byte, password string, data []byte) (id []byte, err error) {
	/*
	 * 1 - Serialize to Identita Binary Format
	 */
	salt := []byte{} // Use National Identification Number
	ciphertext, err := Encrypt([]byte(password), pubKey, salt, data)
	if err != nil {
		return
	}
	signature, err := Sign(key, ciphertext)
	if err != nil {
		return
	}
	persona := new(big.Int)
	persona.SetBytes(append(ciphertext, signature...))
	id = base40.EncodeBig(nil, persona)
	return
}
