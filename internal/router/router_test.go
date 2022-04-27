package router

import (
	"testing"
	"time"

	jwt "github.com/gogf/gf-jwt/v2"
)

// "github.com/tiger1103/gfast-token/gftoken"

type User struct {
	UserData string      // 用户数据
	Data     interface{} // 其他需要携带的数据
}

func Test_Token(t *testing.T) {
	auth := jwt.New(&jwt.GfJWTMiddleware{
		Realm:         "realworld-golang-gf",
		Key:           []byte("secret key"),
		Timeout:       time.Minute * 5,
		MaxRefresh:    time.Minute * 5,
		IdentityKey:   "id",
		TokenLookup:   "header: Authorization",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
		// Authenticator:   Authenticator,
		// Unauthorized:    Unauthorized,
		// PayloadFunc:     PayloadFunc,
		// IdentityHandler: IdentityHandler,
	})

	t.Log(auth)
}

// func Test_Token(t *testing.T) {
// 	gt := gftoken.NewGfToken()
// 	keys, err := gt.GenerateToken(nil, "123", User{
// 		UserData: "123",
// 		Data:     "myData",
// 	})
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	t.Log(keys)
// }
