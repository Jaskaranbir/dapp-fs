package kademlia

import (
	"crypto/rand"
	"crypto/rsa"
	"math/bits"

	sha256 "github.com/minio/sha256-simd"
	"github.com/mr-tron/base58"
	"github.com/pkg/errors"
	"golang.org/x/crypto/ssh"
)

// IDSize represents the number of bytes constituting the ID
const IDSize = 32

// ID represents a Node-ID
type ID string

func (id ID) CommonPrefixLen(target ID) int {
	xor := id.XOR(target)
	return ZeroPrefixLen(string(xor))
}

func (id ID) ToDHT() string {
	dhtID := sha256.Sum256([]byte(id))
	return string(dhtID[:])
}

// XOR returns the XOR-result with provided ID.
func (id ID) XOR(xid ID) []byte {
	aBytes := []byte(id)
	bBytes := []byte(xid)

	result := make([]byte, len(aBytes))
	for i := 0; i < len(result); i++ {
		result[i] = aBytes[i] ^ bBytes[i]
	}
	return result
}

// ZeroPrefixLen returns the number of consecutive zeroes in a byte slice.
func ZeroPrefixLen(in string) int {
	inBytes := []byte(in)

	for i, inByte := range inBytes {
		if inByte != 0 {
			return i*8 + bits.LeadingZeros8(uint8(inByte))
		}
	}
	return len(in) * 8
}

// NewID generates new Node-ID using given RSA public-key.
// If nil public-key is provided, a new key is generated.
func NewID(publicKey []byte) (ID, error) {
	var err error
	if publicKey == nil {
		publicKey, err = genPublicKey(2048)
		if err != nil {
			return "", errors.Wrap(err, "error generating public-key")
		}
	}

	pubKeySHASum := sha256.Sum256(publicKey)
	idBuff := encode(pubKeySHASum[:IDSize])
	id := base58.Encode(idBuff)
	
	return ID(id), nil
}

func genPublicKey(bits int) ([]byte, error) {
	if bits < 512 {
		return nil, errors.New("rsa keys must be >= 512 bits to be useful")
	}
	priv, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, errors.Wrap(err, "error genering rsa key")
	}
	publicRSAKey, err := ssh.NewPublicKey(&priv.PublicKey)
	if err != nil {
		return nil, errors.Wrap(err, "error getting public-key from generated rsa")
	}

	pubKeyBytes := ssh.MarshalAuthorizedKey(publicRSAKey)
	return pubKeyBytes, nil
}

// As per Kademlia specs, the byte-string needs to start
// with string-type (ID, in this case) as first byte,
// followed by length of content as second byte,
// followed by actual content.
func encode(buf []byte) []byte {
	// Prepend type and length to byte-string
	encodedBuf := []byte{0, byte(len(buf))}
	encodedBuf = append(encodedBuf, buf...)
	return encodedBuf
}
