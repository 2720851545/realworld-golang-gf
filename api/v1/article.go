package v1

import (
	"github.com/2720851545/realworld-golang-gf/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

type AllArticlesReq struct {
	g.Meta    `path:"/articles" method:"get" tags:"文章" summary:"所有文章列表"`
	Author    string `p:"author"`
	Favorited string `p:"favorited"`
	Tag       string `p:"tag"`
	Offset    int    `d:"0" p:"offset"`
	Limit     int    `d:"10" p:"limit"`
}

type CreateArticleReq struct {
	g.Meta  `path:"/articles" method:"post" tags:"文章" summary:"创建文章"`
	Article createArticleReqArticle `json:"article"`
}

// todo: 1. 排除authorId字段
// 		 2. 日期格式化
type CreateArticleRes struct {
	g.Meta  `mime:"application/json"`
	Article AllArticleResArticle `json:"article"`
}

type UpdateArticleReq struct {
	g.Meta  `path:"/articles/:slug" method:"put" tags:"文章" summary:"修改文章"`
	Slug    string                  `p:"slug" v:"required#标题不能为空"`
	Article updateArticleReqArticle `json:"article"`
}

type UpdateArticleRes struct {
	g.Meta  `mime:"application/json"`
	Article AllArticleResArticle `json:"article"`
}

type FavoriteArticleReq struct {
	g.Meta  `path:"/articles/:slug/favorite" method:"post" tags:"文章" summary:"收藏文章"`
	Slug    string                  `p:"slug" v:"required#标题不能为空"`
	Article updateArticleReqArticle `json:"article"`
}

type FavoriteArticleRes struct {
	g.Meta  `mime:"application/json"`
	Article AllArticleResArticle `json:"article"`
}

type UnfavoriteArticleReq struct {
	g.Meta `path:"/articles/:slug/favorite" method:"delete" tags:"文章" summary:"取消收藏文章"`
	Slug   string `p:"slug" v:"required#标题不能为空"`
}

type UnfavoriteArticleRes struct {
	g.Meta  `mime:"application/json"`
	Article AllArticleResArticle `json:"article"`
}

type CreateCommentArticleReq struct {
	g.Meta  `path:"/articles/:slug/comments" method:"post" tags:"文章" summary:"创建文章评论"`
	Slug    string                         `p:"slug" v:"required#标题不能为空"`
	Comment createCommentArticleReqComment `json:"comment"`
}

type CreateCommentArticleRes struct {
	g.Meta  `mime:"application/json"`
	Comment createCommentArticleResComment `json:"comment"`
}

type AllCommentForArticleReq struct {
	g.Meta `path:"/articles/:slug/comments" method:"get" tags:"文章" summary:"文章评论列表"`
	Slug   string `p:"slug" v:"required#标题不能为空"`
}

type AllCommentForArticleRes struct {
	g.Meta   `mime:"application/json"`
	Comments []createCommentArticleResComment `json:"comments"`
}

type DeleteCommentForArticleReq struct {
	g.Meta    `path:"/articles/:slug/comments/:commentId" method:"delete" tags:"文章" summary:"删除文章评论"`
	Slug      string `p:"slug" v:"required#标题不能为空"`
	CommentId string `p:"commentId" v:"required#评论id不能为空"`
}

type DeleteCommentForArticleRes struct {
	g.Meta `mime:"application/json"`
}

type DeleteArticleReq struct {
	g.Meta `path:"/articles/:slug" method:"delete" tags:"文章" summary:"删除文章"`
	Slug   string `p:"slug" v:"required#标题不能为空"`
}

type DeleteArticleRes struct {
	g.Meta `mime:"application/json"`
}

type AllArticleRes struct {
	g.Meta `mime:"application/json"`
	// todo 返回日期, 需要引入第三方json解析,自闭了
	Articles      []AllArticleResArticle `json:"articles"`
	ArticlesCount int                    `json:"articlesCount"`
}

type FeedArticleReq struct {
	g.Meta `path:"/articles/feed" method:"get" tags:"文章" summary:"获取用户关注的文章"`
}

type FeedArticleRes struct {
	g.Meta        `mime:"application/json"`
	Articles      []AllArticleResArticle `json:"articles"`
	ArticlesCount int                    `json:"articlesCount"`
}

type SignalArticleBySulgReq struct {
	g.Meta `path:"/articles/:slug" method:"get" tags:"文章" summary:"获取某个文章"`
	Slug   string `p:"slug"`
}

type SignalArticleBySulgRes struct {
	g.Meta  `mime:"application/json"`
	Article AllArticleResArticle `json:"article"`
}

type AllArticleResArticle struct {
	entity.Article
	TagList        []string                   `json:"tagList"`
	Author         allArticleResArticleAuthor `json:"author"`
	Favorited      bool                       `json:"favorited"`
	FavoritesCount int                        `json:"favoritesCount"`
}

type allArticleResArticleAuthor struct {
	Bio       string `json:"bio"`
	Following bool   `json:"following"`
	Image     string `json:"image"`
	Username  string `json:"username"`
}

type createArticleReqArticle struct {
	Title       string   `json:"title"       `
	Description string   `json:"description" `
	Body        string   `json:"body"        `
	TagList     []string `json:"tagList"    `
}

type updateArticleReqArticle struct {
	Id          int      `json:"id"`
	Title       string   `json:"title"       `
	Description string   `json:"description" `
	Body        string   `json:"body"        `
	TagList     []string `json:"tagList"    `
}

type createCommentArticleReqComment struct {
	Body string `json:"body"`
}

type createCommentArticleResComment struct {
	entity.Comment
	Author allArticleResArticleAuthor `json:"author"`
}
