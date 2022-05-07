package controller

import (
	"context"

	v1 "github.com/2720851545/realworld-golang-gf/api/v1"
)

var (
	AuthArticleController = authArticleController{}
)

type authArticleController struct {
}

func (t *authArticleController) CreateArticle(ctx context.Context, req *v1.CreateArticleReq) (res *v1.CreateArticleRes, err error) {
	res, err = articleService.CreateArticle(ctx, req)
	return
}
