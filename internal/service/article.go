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
		assignmentArticleDetailInfo(ctx, &res.Articles[i])
	}
	return
}

func (t *articleService) CreateArticle(ctx context.Context, req *v1.CreateArticleReq) (res *v1.CreateArticleRes, err error) {
	var articleDo do.Article
	if articleDo, err = validateSaveOrUpdateArticle(ctx, req.Article); err != nil {
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
	assignmentArticleDetailInfo(ctx, &res.Article)

	return
}

func (t *articleService) UpdateArticle(ctx context.Context, req *v1.UpdateArticleReq) (res *v1.UpdateArticleRes, err error) {
	var articleDo do.Article
	if articleDo, err = validateSaveOrUpdateArticle(ctx, req.Article); err != nil {
		return
	}

	if !g.IsEmpty(articleDo.Title) {
		articleDo.Slug = slug.Make(gconv.String(articleDo.Title))
	}
	articleDo.UpdatedAt = gtime.Now()

	if _, err = dao.Article.Ctx(ctx).Where("slug = ", req.Slug).OmitEmpty().Update(articleDo); err != nil {
		return
	}

	res = new(v1.UpdateArticleRes)
	err = dao.Article.Ctx(ctx).Where("slug = ?", req.Slug).Scan(&res.Article)
	assignmentArticleDetailInfo(ctx, &res.Article)
	return
}

func validateSaveOrUpdateArticle(ctx context.Context, reqArticle interface{}) (articleDo do.Article, err error) {
	if err = gconv.Struct(reqArticle, &articleDo); err != nil {
		return
	}

	return
}

func (t *articleService) FeedArticle(ctx context.Context, req *v1.FeedArticleReq) (res *v1.FeedArticleRes, err error) {
	res = new(v1.FeedArticleRes)
	err = dao.Article.Ctx(ctx).Where("author_id in (select following_id from follow where followed_by_id = ? )",
		GetUserId(ctx)).Scan(&res.Articles)
	for i := range res.Articles {
		assignmentArticleDetailInfo(ctx, &res.Articles[i])
	}
	return
}

func (t *articleService) SignalArticleBySulg(ctx context.Context, req *v1.SignalArticleBySulgReq) (res *v1.SignalArticleBySulgRes, err error) {
	res = new(v1.SignalArticleBySulgRes)
	err = dao.Article.Ctx(ctx).Where("slug = ?", req.Slug).Scan(&res.Article)
	assignmentArticleDetailInfo(ctx, &res.Article)
	return
}

func (t *articleService) FavoriteArticle(ctx context.Context, req *v1.FavoriteArticleReq) (res *v1.FavoriteArticleRes, err error) {
	res = new(v1.FavoriteArticleRes)
	var id gdb.Value
	if id, err = dao.Article.Ctx(ctx).Where("slug = ?", req.Slug).Value("id"); err != nil {
		return
	}

	userId := GetUserId(ctx)
	if i, err := dao.Favorite.Ctx(ctx).Count(
		"favorite_id = ? and favorite_by_id = ?", id, userId); i == 0 && err == nil {
		dao.Favorite.Ctx(ctx).Insert(do.Favorite{
			CreatedAt:    gtime.Now(),
			FavoriteId:   id,
			FavoriteById: userId,
		})
	}
	err = dao.Article.Ctx(ctx).Where("id = ?", id).Scan(&res.Article)
	assignmentArticleDetailInfo(ctx, &res.Article)
	return
}

func (t *articleService) UnfavoriteArticle(ctx context.Context, req *v1.UnfavoriteArticleReq) (res *v1.UnfavoriteArticleRes, err error) {
	res = new(v1.UnfavoriteArticleRes)
	var id gdb.Value
	if id, err = dao.Article.Ctx(ctx).Where("slug = ?", req.Slug).Value("id"); err != nil {
		return
	}
	dao.Favorite.Ctx(ctx).Delete("favorite_id = ? and favorite_by_id = ?", id, GetUserId(ctx))

	err = dao.Article.Ctx(ctx).Where("id = ?", id).Scan(&res.Article)
	assignmentArticleDetailInfo(ctx, &res.Article)
	return
}

func (t *articleService) CreateCommentArticle(ctx context.Context, req *v1.CreateCommentArticleReq) (res *v1.CreateCommentArticleRes, err error) {
	res = new(v1.CreateCommentArticleRes)

	var articleId gdb.Value
	if articleId, err = dao.Article.Ctx(ctx).Where("slug = ?", req.Slug).Value("id"); err != nil {
		return
	}
	commentId, err := dao.Comment.Ctx(ctx).InsertAndGetId(do.Comment{
		CreatedAt: gtime.Now(),
		UpdatedAt: gtime.Now(),
		ArticleId: articleId,
		AuthorId:  GetUserId(ctx),
		Body:      req.Comment.Body,
	})
	if err != nil {
		return
	}

	err = dao.Comment.Ctx(ctx).Where("id = ?", commentId).Scan(&res.Comment)
	if err != nil {
		return
	}
	err = dao.User.Ctx(ctx).Where("id = ?", res.Comment.AuthorId).Scan(&res.Comment.Author)
	if err != nil {
		return
	}
	res.Comment.Author.Following = isFollowingUser(ctx, res.Comment.AuthorId, gconv.Int(GetUserId(ctx)))
	return
}

func (t *articleService) DeleteCommentForArticle(ctx context.Context, req *v1.DeleteCommentForArticleReq) (err error) {
	dao.Comment.Ctx(ctx).Delete("id = ?", req.CommentId)
	return
}

func (t *articleService) DeleteArticle(ctx context.Context, req *v1.DeleteArticleReq) (err error) {
	dao.Article.Ctx(ctx).Delete("slug = ?", req.Slug)
	return
}

func (t *articleService) AllCommentForArticle(ctx context.Context, req *v1.AllCommentForArticleReq) (res *v1.AllCommentForArticleRes, err error) {
	res = new(v1.AllCommentForArticleRes)
	articleId, err := dao.Article.Ctx(ctx).Where("slug = ?", req.Slug).Value("id")
	dao.Comment.Ctx(ctx).Scan(&res.Comments, g.Map{
		"article_id": articleId,
	})

	for i := range res.Comments {
		comment := &res.Comments[i]
		err = dao.User.Ctx(ctx).Where("id = ?", comment.AuthorId).Scan(&comment.Author)
		comment.Author.Following = isFollowingUser(ctx, comment.AuthorId, gconv.Int(GetUserId(ctx)))
	}
	return
}

func assignmentArticleDetailInfo(ctx context.Context, article *v1.AllArticleResArticle) {
	dao.User.Ctx(ctx).Where("id = ?", article.AuthorId).Scan(&article.Author)
	article.Author.Following = isFollowingUser(ctx, article.AuthorId, gconv.Int(GetUserId(ctx)))
	article.FavoritesCount, _ = favoritesCount(ctx, gconv.Int(article.Id))
	article.TagList = getTags(ctx, gconv.Int(article.Id))
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
