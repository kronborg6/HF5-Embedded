package middleware

import (
	b64 "encoding/base64"
)

func Encode(password string) string {
	newPass := b64.URLEncoding.EncodeToString([]byte(password))

	return newPass
}

func Dcode(password string) string {
	pass, _ := b64.URLEncoding.DecodeString(password)

	return string(pass)
}
