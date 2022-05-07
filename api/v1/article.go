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

type AllArticleRes struct {
	g.Meta `mime:"application/json"`
	// todo 返回日期, 需要引入第三方json解析,自闭了
	Articles      []allArticleResArticle `json:"articles"`
	ArticlesCount int                    `json:"articlesCount"`
}

type allArticleResArticle struct {
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
	Article     entity.Article `json:"article"`
	Title       string         `json:"title"       `
	Description string         `json:"description" `
	Body        string         `json:"body"        `
	TagList     []string       `json:"tagList"    `
}

// todo: 1. 排除authorId字段
// 		 2. 日期格式化
type CreateArticleRes struct {
	g.Meta  `mime:"application/json"`
	Article allArticleResArticle `json:"article"`
}
