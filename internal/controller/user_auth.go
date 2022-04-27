package controller

import (
	"context"

	v1 "github.com/2720851545/realworld-golang-gf/api/v1"
)

var AuthUserController = authUserController{}

type authUserController struct {
}

func (c *authUserController) CurrentUser(ctx context.Context, req *v1.CurrentUserReq) (res *v1.CurrentUserRes, err error) {
	// g.Log().Info(ctx, req)
	// res, err = service.UserService().Register(ctx, req)
	return
}
