package controller

import (
	"context"

	v1 "github.com/2720851545/realworld-golang-gf/api/v1"
	"github.com/2720851545/realworld-golang-gf/internal/service"
)

var (
	NoAuthArticleController = noAuthArticleController{}

	articleService = service.ArticleService
)

type noAuthArticleController struct {
}

func (t *noAuthArticleController) AllArticle(ctx context.Context, req *v1.AllArticlesReq) (res *v1.AllArticleRes, err error) {
	res, err = articleService.AllArticle(ctx, req)
	return
}

func (t *noAuthArticleController) SignalArticleBySulg(ctx context.Context, req *v1.SignalArticleBySulgReq) (res *v1.SignalArticleBySulgRes, err error) {
	res, err = articleService.SignalArticleBySulg(ctx, req)
	return
}

func (t *noAuthArticleController) AllCommentForAricle(ctx context.Context, req *v1.AllCommentForArticleReq) (res *v1.AllCommentForArticleRes, err error) {
	res, err = articleService.AllCommentForArticle(ctx, req)
	return
}

func (t *noAuthArticleController) TagsForArticle(ctx context.Context, req *v1.TagsReq) (res *v1.TagsRes, err error) {
	res, err = articleService.TagsForArticle(ctx, req)
	return
}
