package helper

import (
	"crypto/md5"
	"encoding/hex"
)

// HashPassword 函数用于加密密码
func HashPassword(pass string) string {
	hasher := md5.New()
	hasher.Write([]byte(pass))
	firstHash := hex.EncodeToString(hasher.Sum(nil))
	hasher.Reset()
	hasher.Write([]byte(firstHash + "Qwert!@#456"))
	return hex.EncodeToString(hasher.Sum(nil))
}
