package service

import "github.com/gogf/gf/g/crypto/gmd5"

const (
	ENCRYPTMD5 = "gadmin"
)

// EncryptPassword 加密密码
//
// createTime:2019年04月25日 10:19:13
// author:hailaz
func EncryptPassword(data string) string {
	res, _ := gmd5.EncryptString(data + ENCRYPTMD5)
	return res
}
