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

func (t *authArticleController) FeedArticle(ctx context.Context, req *v1.FeedArticleReq) (res *v1.FeedArticleRes, err error) {
	res, err = articleService.FeedArticle(ctx, req)
	return
}

func (t *authArticleController) UpdateArticle(ctx context.Context, req *v1.UpdateArticleReq) (res *v1.UpdateArticleRes, err error) {
	res, err = articleService.UpdateArticle(ctx, req)
	return
}

func (t *authArticleController) FavoriteArticle(ctx context.Context, req *v1.FavoriteArticleReq) (res *v1.FavoriteArticleRes, err error) {
	res, err = articleService.FavoriteArticle(ctx, req)
	return
}

func (t *authArticleController) UnfavoriteArticle(ctx context.Context, req *v1.UnfavoriteArticleReq) (res *v1.UnfavoriteArticleRes, err error) {
	res, err = articleService.UnfavoriteArticle(ctx, req)
	return
}

func (t *authArticleController) CreateCommentArticle(ctx context.Context, req *v1.CreateCommentArticleReq) (res *v1.CreateCommentArticleRes, err error) {
	res, err = articleService.CreateCommentArticle(ctx, req)
	return
}

func (t *authArticleController) DeleteCommentForArticle(ctx context.Context, req *v1.DeleteCommentForArticleReq) (res *v1.DeleteCommentForArticleRes, err error) {
	err = articleService.DeleteCommentForArticle(ctx, req)
	res = new(v1.DeleteCommentForArticleRes)
	return
}

func (t *authArticleController) DeleteArticle(ctx context.Context, req *v1.DeleteArticleReq) (res *v1.DeleteArticleRes, err error) {
	err = articleService.DeleteArticle(ctx, req)
	res = new(v1.DeleteArticleRes)
	return
}
