package service

import (
	"context"
	"time"

	jwt "github.com/gogf/gf-jwt/v2"
	"github.com/gogf/gf/v2/frame/g"
)

var authService *jwt.GfJWTMiddleware

func Auth() *jwt.GfJWTMiddleware {
	return authService
}

func init() {
	auth := jwt.New(&jwt.GfJWTMiddleware{
		Realm:           "realworld-golang-gf",
		Key:             []byte("Wnm7UCzloiyN"),
		Timeout:         time.Minute * 60 * 24,
		MaxRefresh:      time.Minute * 30,
		IdentityKey:     "id",
		TokenLookup:     "header: Authorization",
		TokenHeadName:   "Token",
		TimeFunc:        time.Now,
		Authenticator:   Authenticator,
		Unauthorized:    Unauthorized,
		PayloadFunc:     PayloadFunc,
		IdentityHandler: IdentityHandler,
	})
	authService = auth
}

func PayloadFunc(data interface{}) jwt.MapClaims {
	claims := jwt.MapClaims{}
	params := data.(map[string]interface{})
	if len(params) > 0 {
		for k, v := range params {
			claims[k] = v
		}
	}
	return claims
}

func IdentityHandler(ctx context.Context) interface{} {
	claims := jwt.ExtractClaims(ctx)
	return claims[authService.IdentityKey]
}

func Unauthorized(ctx context.Context, code int, message string) {
	r := g.RequestFromCtx(ctx)
	r.Response.Status = 411
	r.Response.WriteJson(g.Map{
		"errors": g.Array{message},
	})
	r.ExitAll()
}

func Authenticator(ctx context.Context) (interface{}, error) {
	ctx = g.RequestFromCtx(ctx).Context()
	switch ctx.Value("Model").(string) {
	case "Register", "Login":
		return ctx.Value("User"), nil
	default:
		return nil, jwt.ErrFailedAuthentication
	}
}
