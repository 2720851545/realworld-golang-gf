package controller

import (
	// "context"

	// v1 "github.com/2720851545/realworld-golang-gf/api/v1"
	"context"

	v1 "github.com/2720851545/realworld-golang-gf/api/v1"
	"github.com/2720851545/realworld-golang-gf/internal/service"
)

var NoAuthUserController = noAuthUserController{}

type noAuthUserController struct {
}

func (c *noAuthUserController) Register(ctx context.Context, req *v1.UserRegisterReq) (res *v1.UserRegisterRes, err error) {
	res, err = service.UserService().Register(ctx, req)
	return
}
