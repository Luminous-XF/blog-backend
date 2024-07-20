package helper

import (
	"crypto/md5"
	"encoding/hex"
)

// MD5 对文本进行 MD5 加密
func MD5(text string) string {
	data := md5.New()
	data.Write([]byte(text))
	digest := data.Sum(nil)
	return hex.EncodeToString(digest)
}
