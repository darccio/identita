package identita

import (
	"golang.org/x/crypto/scrypt"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"math"
)

const keySize = 16

// Decrypt https://gist.github.com/manishtpatel/8222606

func StretchKey(password, salt []byte, size int) (key []byte, err error) {
	return scrypt.Key(password, salt, int(math.Pow(2, 15)), 16, 2, size)
}

func Encrypt(password, pubKey, salt, data []byte) (ciphertext []byte, err error) {
	salt = append(salt, pubKey...)
	salt = append(salt, password...)
	key, err := StretchKey(password, salt, keySize)
	if err != nil {
		return
	}
	c, err := aes.NewCipher(key)
	if err != nil {
		return
	}
	ivSize := keySize / 4
	ciphertext = make([]byte, ivSize+len(data))
	iv := ciphertext[:ivSize]
	if err = getRandomBytes(iv); err != nil {
		return
	}
	trueIv, err := StretchKey(iv, key, keySize)
	if err != nil {
		return
	}
	stream := cipher.NewCFBEncrypter(c, trueIv)
	stream.XORKeyStream(ciphertext[ivSize:], data)
	return
}

func getRandomBytes(b []byte) (err error) {
	_, err = io.ReadFull(rand.Reader, b)
	return
}
