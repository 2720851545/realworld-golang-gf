package controller

import (
	"context"

	v1 "github.com/2720851545/realworld-golang-gf/api/v1"
	"github.com/2720851545/realworld-golang-gf/internal/service"
)

var (
	ArticleController = articleController{}

	articleService = service.ArticleService
)

type articleController struct {
}

func (t *articleController) AllArticle(ctx context.Context, req *v1.AllArticlesReq) (res *v1.AllArticleRes, err error) {
	res, err = articleService.AllArticle(ctx, req)
	return
}
