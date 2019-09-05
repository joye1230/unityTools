package md5

import (
	"crypto/md5"
	"fmt"
	"strings"
)

func GetMd5Value(encryptString string) string {
	h := md5.New()
	h.Write([]byte(encryptString))
	md5Value := strings.ToUpper(fmt.Sprintf("%x", h.Sum(nil)))
	return md5Value
}
