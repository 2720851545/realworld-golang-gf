package v1

import (
	"github.com/2720851545/realworld-golang-gf/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

type AllArticlesReq struct {
	g.Meta `path:"/articles" method:"get" tags:"文章" summary:"所有文章列表"`
}

type AllArticleRes struct {
	g.Meta `mime:"application/json"`
	// todo 返回日期, 需要引入第三方json解析,自闭了
	Articles      []allArticleResArticle `json:"articles"`
	ArticlesCount int                    `json:"articlesCount"`
}

type allArticleResArticle struct {
	entity.Article
	TagList   []string                   `json:"tagList"`
	Author    allArticleResArticleAuthor `json:"author"`
	Favorited bool                       `json:"favorited"`
}

type allArticleResArticleAuthor struct {
	Bio       string `json:"bio"`
	Following bool   `json:"following"`
	Image     string `json:"image"`
	Username  string `json:"username"`
}
