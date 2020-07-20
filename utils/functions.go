package utils

import (
	"bytes"
	"crypto/md5"
	"fmt"
)

//GeneratePassword 生成密码
func GeneratePassword(src string, salt string) string {
	//密码验证
	encryptP := md5.Sum([]byte(salt))
	strbuf := bytes.NewBufferString(src)
	md5salt := fmt.Sprintf("%x", encryptP)
	_, _ = strbuf.WriteString(md5salt)
	encryptP = md5.Sum(strbuf.Bytes())
	return fmt.Sprintf("%x", encryptP)
}