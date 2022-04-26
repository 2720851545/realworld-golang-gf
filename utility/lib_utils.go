package utility

import (
	"github.com/gogf/gf/v2/crypto/gmd5"
)

func EntryPassword(password string) string {
	return gmd5.MustEncryptString(password + salt)
}
