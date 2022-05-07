package service

import (
	"context"

	v1 "github.com/2720851545/realworld-golang-gf/api/v1"
	"github.com/2720851545/realworld-golang-gf/internal/service/internal/dao"
	"github.com/2720851545/realworld-golang-gf/internal/service/internal/do"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gosimple/slug"
)

var ArticleService = &articleService{}

type articleService struct {
}

func (t *articleService) AllArticle(ctx context.Context, req *v1.AllArticlesReq) (res *v1.AllArticleRes, err error) {
	res = new(v1.AllArticleRes)
	var model *gdb.Model = dao.Article.Ctx(ctx).Safe(false)
	switch true {
	case !g.IsEmpty(req.Author):
		model.Where("author_id = ?", req.Author)
	case !g.IsEmpty(req.Tag):
		model.Where("id in (SELECT article_model_id FROM article_tags WHERE tag_model_id = (SELECT id FROM tag WHERE tag = ?))", req.Tag)
	case !g.IsEmpty(req.Favorited):
		model.Where("id in (SELECT favorite_id FROM favorite WHERE favorite_by_id = (select id from user where  username =?))",
			req.Favorited)
	}
	res.ArticlesCount, _ = model.Count()
	err = model.Limit(req.Offset, req.Limit).Scan(&res.Articles)

	for i := range res.Articles {
		article := &res.Articles[i]
		dao.User.Ctx(ctx).Where("id = ?", article.AuthorId).Scan(&article.Author)
		article.Author.Following = isFollowingUser(ctx, gconv.Int(article.AuthorId), gconv.Int(GetUserId(ctx)))
		article.FavoritesCount, _ = favoritesCount(ctx, gconv.Int(article.Id))
		article.TagList = getTags(ctx, gconv.Int(article.Id))
	}
	return
}

func (t *articleService) CreateArticle(ctx context.Context, req *v1.CreateArticleReq) (res *v1.CreateArticleRes, err error) {
	articleDo := do.Article{}
	if err = gconv.Struct(req.Article, &articleDo); err != nil {
		return
	}
	createdUserId := gconv.Int(GetUserId(ctx))
	articleDo.Slug = slug.Make(gconv.String(articleDo.Title))
	articleDo.AuthorId = createdUserId
	articleDo.CreatedAt = gtime.Now()
	articleDo.UpdatedAt = gtime.Now()
	lastInsertId, err := dao.Article.Ctx(ctx).InsertAndGetId(articleDo)
	if err != nil {
		return
	}

	err = SaveTags(ctx, lastInsertId, req.Article.TagList)
	if err != nil {
		return
	}

	res = new(v1.CreateArticleRes)
	err = dao.Article.Ctx(ctx).Where("id = ?", lastInsertId).Scan(&res.Article)

	dao.User.Ctx(ctx).Where("id = ?", createdUserId).Scan(&res.Article.Author)
	res.Article.Author.Following = isFollowingUser(ctx, createdUserId, createdUserId)
	res.Article.FavoritesCount, _ = favoritesCount(ctx, gconv.Int(lastInsertId))
	res.Article.TagList = getTags(ctx, gconv.Int(res.Article.Id))
	return
}

func favoritesCount(ctx context.Context, articleId int) (count int, err error) {
	count, err = dao.Favorite.Ctx(ctx).Where("favorite_id = ?", articleId).Count()
	return
}

func getTags(ctx context.Context, articleId int) (tagList []string) {
	tagListInterface, _ := g.DB().GetArray(ctx,
		"SELECT t.tag FROM article_tags at inner join tag t on at.tag_model_id = t.id WHERE at.article_model_id = ?",
		articleId)
	tagList = gconv.Strings(tagListInterface)
	return
}
