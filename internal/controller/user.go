package controller

import (
	// "context"

	// v1 "github.com/2720851545/realworld-golang-gf/api/v1"
	"context"

	v1 "github.com/2720851545/realworld-golang-gf/api/v1"
	"github.com/2720851545/realworld-golang-gf/internal/service"
)

var (
	NoAuthUserController = noAuthUserController{}
	userService          = service.UserService()
)

type noAuthUserController struct {
}

func (c *noAuthUserController) Register(ctx context.Context, req *v1.UserRegisterReq) (res *v1.UserRegisterRes, err error) {
	res, err = userService.Register(ctx, req)
	return
}

func (c *noAuthUserController) Login(ctx context.Context, req *v1.UserLoginReq) (res *v1.UserLoginRes, err error) {
	res, err = userService.Login(ctx, req)
	return
}

func (c *noAuthUserController) Profile(ctx context.Context, req *v1.ProfileReq) (res *v1.ProfileRes, err error) {
	res, err = userService.Profile(ctx, req)
	return
}
