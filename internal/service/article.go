package service

import (
	"context"

	v1 "github.com/2720851545/realworld-golang-gf/api/v1"
	"github.com/2720851545/realworld-golang-gf/internal/service/internal/dao"
)

var ArticleService = &articleService{}

type articleService struct {
}

func (t *articleService) AllArticle(ctx context.Context, req *v1.AllArticlesReq) (res *v1.AllArticleRes, err error) {
	res = new(v1.AllArticleRes)
	err = dao.Article.Ctx(ctx).Scan(&res.Articles)
	return
}
