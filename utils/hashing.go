package utils

import (
	"crypto/sha256"
	"fmt"
)

func asSha256(input interface{}) string {
	hash := sha256.New()
	hash.Write([]byte(fmt.Sprintf("%v", input)))
	return fmt.Sprintf("%x", hash.Sum(nil))
}
