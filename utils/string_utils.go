package utils

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
func Vdd(str string)bool  {
	if strings.Contains(str,".."){
		return true
	}
	return false
}
