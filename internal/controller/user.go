package controller

import (
	"context"

	v1 "github.com/2720851545/realworld-golang-gf/api/v1"
	"github.com/2720851545/realworld-golang-gf/internal/service"
	"github.com/gogf/gf/v2/frame/g"
)

var UserController = userController{}

type userController struct {
}

func (c *userController) Register(ctx context.Context, req *v1.UserRegisterReq) (res *v1.UserRegisterRes, err error) {
	g.Log().Info(ctx, req)
	res, err = service.UserService().Register(ctx, req)
	return
}
