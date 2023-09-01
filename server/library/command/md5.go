package command

import (
	"crypto/md5"
	"encoding/hex"
)

//MD5 md5加密
func MD5(keyword string) (string, error) {
	md := md5.New()
	md.Write([]byte(keyword))
	cipherStr := hex.EncodeToString(md.Sum(nil))

	return cipherStr, nil
}
