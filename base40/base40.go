package base40

// Based on github.com/tv42/base58

import (
	"math/big"
	"strconv"
)

// The alphabet has been balanced in order to avoid big similar
// characters clusters and generate human-friendly long strings.
// Also, the alphabet has been limited to ASCII range [42,90] in
// order to allow to pack characters in six bits each.
const alphabet = "0ABV 1CD-2EFW3GH/4IJX5KL*6MNY7PQ.8RSZ#9TU@"
const alphabetSize = int64(len(alphabet))

var decodeMap [256]byte

func init() {
	for i := 0; i < len(decodeMap); i++ {
		decodeMap[i] = 0xFF
	}
	for i := 0; i < len(alphabet); i++ {
		decodeMap[alphabet[i]] = byte(i)
	}
}

type CorruptInputError int64

func (e CorruptInputError) Error() string {
	return "illegal base" + strconv.Itoa(int(alphabetSize)) + " data at input byte " + strconv.FormatInt(int64(e), 10)
}

// Decode a big integer from the bytes. Returns an error on corrupt
// input.
func DecodeToBig(src []byte) (*big.Int, error) {
	n := new(big.Int)
	radix := big.NewInt(alphabetSize)
	for i := 0; i < len(src); i++ {
		b := decodeMap[src[i]]
		if b == 0xFF {
			return nil, CorruptInputError(i)
		}
		n.Mul(n, radix)
		n.Add(n, big.NewInt(int64(b)))
	}
	return n, nil
}

// Encode encodes src, appending to dst. Be sure to use the returned
// new value of dst.
func EncodeBig(dst []byte, src *big.Int) []byte {
	start := len(dst)
	n := new(big.Int)
	n.Set(src)
	radix := big.NewInt(alphabetSize)
	zero := big.NewInt(0)

	for n.Cmp(zero) > 0 {
		mod := new(big.Int)
		n.DivMod(n, radix, mod)
		dst = append(dst, alphabet[mod.Int64()])
	}

	for i, j := start, len(dst)-1; i < j; i, j = i+1, j-1 {
		dst[i], dst[j] = dst[j], dst[i]
	}
	return dst
}
