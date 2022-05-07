package controller

import (
	"context"

	v1 "github.com/2720851545/realworld-golang-gf/api/v1"
)

var AuthUserController = authUserController{}

type authUserController struct {
}

func (c *authUserController) CurrentUser(ctx context.Context, req *v1.CurrentUserReq) (res *v1.CurrentUserRes, err error) {
	res, err = userService.CurrentUser(ctx, req)
	return
}

func (c *authUserController) UpdateUserInfo(ctx context.Context, req *v1.UserUpdateReq) (res *v1.UserUpdateRes, err error) {
	res, err = userService.UpdateUserInfo(ctx, req)
	return
}

func (c *authUserController) FollowProfile(ctx context.Context, req *v1.FollowProfileReq) (res *v1.FollowProfileRes, err error) {
	res, err = userService.FollowProfile(ctx, req)
	return
}

func (c *authUserController) UnfollowProfile(ctx context.Context, req *v1.UnFollowProfileReq) (res *v1.UnFollowProfileRes, err error) {
	res, err = userService.UnfollowProfile(ctx, req)
	return
}
