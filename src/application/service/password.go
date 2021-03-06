package service

import (
	"crypto/md5"
	"fmt"
	"io"
)

func EncryptPassword(password string) string {
	tPass := md5.New()
	io.WriteString(tPass, password)
	ePassword := fmt.Sprintf("%x", tPass.Sum(nil))

	return ePassword
}

func IsPasswordMatch(password, userpass string) bool {
	ePassword := EncryptPassword(password)

	if ePassword != userpass {
		return false
	}

	return true
}
