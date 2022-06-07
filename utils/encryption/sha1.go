package encryption

import (
	"fmt"
	"crypto/sha1"
)

//执行sha1加密
func Sha1(data string) string {
	sha1 := sha1.New()
	sha1.Write([]byte(data))
	has := sha1.Sum([]byte(""))
	return fmt.Sprintf("%x", has)
}
