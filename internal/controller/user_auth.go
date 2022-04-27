package controller

import (
	"context"

	v1 "github.com/2720851545/realworld-golang-gf/api/v1"
	"github.com/2720851545/realworld-golang-gf/internal/service"
)

var AuthUserController = authUserController{}

type authUserController struct {
}

func (c *authUserController) CurrentUser(ctx context.Context, req *v1.CurrentUserReq) (res *v1.CurrentUserRes, err error) {
	res, err = service.UserService().CurrentUser(ctx, req)
	return
}
