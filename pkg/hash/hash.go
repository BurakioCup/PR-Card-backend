package hash

import(
	"crypto/sha256"
)

func CreateHashString(str string) string {
	sum := sha256.Sum256([]byte(str))
	return string(sum[:])
}