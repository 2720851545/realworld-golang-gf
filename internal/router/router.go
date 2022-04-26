package router

import (
	"github.com/2720851545/realworld-golang-gf/internal/controller"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func BindController(group *ghttp.RouterGroup) {
	group.Middleware(ghttp.MiddlewareCORS)
	group.Middleware(CustomizeMiddlewareHandlerResponse)

	group.Group("/users", func(group *ghttp.RouterGroup) {
		group.Bind(controller.UserController)
	})
}

// 参考 ghttp.MiddlewareHandlerResponse
func CustomizeMiddlewareHandlerResponse(r *ghttp.Request) {
	r.Middleware.Next()

	// There's custom buffer content, it then exits current handler.
	if r.Response.BufferLength() > 0 {
		return
	}

	var (
		ctx = r.Context()
		err = r.GetError()
		res = r.GetHandlerResponse()
	)
	if err != nil {
		res = err.Error()
		r.Response.Status = 500
	}
	internalErr := r.Response.WriteJson(res)
	if internalErr != nil {
		g.Log().Errorf(ctx, `%+v`, internalErr)
	}
}
