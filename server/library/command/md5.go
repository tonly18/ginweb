package command

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5(keyword string) (string, error) {
	md := md5.New()
	md.Write([]byte(keyword))
	cipherStr := hex.EncodeToString(md.Sum(nil))

	return cipherStr, nil
}
