package router

import (
	"github.com/2720851545/realworld-golang-gf/internal/controller"
	"github.com/2720851545/realworld-golang-gf/internal/model/vo"
	"github.com/2720851545/realworld-golang-gf/internal/service"
	jwt "github.com/gogf/gf-jwt/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

var jwtAuth = service.Auth()

func BindController(group *ghttp.RouterGroup) {
	group.Middleware(ghttp.MiddlewareCORS)
	group.Middleware(CustomizeMiddlewareHandlerResponse)

	group.Bind(controller.NoAuthUserController)
	group.Bind(controller.NoAuthArticleController)
	group.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(AuthMiddleware)
		group.Bind(controller.AuthUserController)
		group.Bind(controller.AuthArticleController)
	})

}

func AuthMiddleware(r *ghttp.Request) {
	jwtAuth.MiddlewareFunc()(r)
	r.Middleware.Next()
}

// 参考 ghttp.MiddlewareHandlerResponse
func CustomizeMiddlewareHandlerResponse(r *ghttp.Request) {

	// 所有路由注入登陆用户信息
	claims, _, err := service.Auth().GetClaimsFromJWT(r.GetCtx())
	if err == nil && int64(claims["exp"].(float64)) >= jwtAuth.TimeFunc().Unix() {
		r.SetParam(jwt.PayloadKey, claims)
	}

	r.Middleware.Next()

	// There's custom buffer content, it then exits current handler.
	if r.Response.BufferLength() > 0 {
		return
	}

	var (
		ctx = r.Context()
		res = r.GetHandlerResponse()
	)
	err = r.GetError()
	if err != nil {
		r.Response.Status = 500
		res = vo.Error{
			Errors: g.Map{
				"": g.Array{err},
			},
		}
	}
	internalErr := r.Response.WriteJson(res)
	if internalErr != nil {
		g.Log().Errorf(ctx, `%+v`, internalErr)
	}
}
