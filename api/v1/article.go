package v1

import (
	"github.com/2720851545/realworld-golang-gf/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

type AllArticlesReq struct {
	g.Meta `path:"/articles" method:"get" tags:"文章" summary:"所有文章列表"`
}

type AllArticleRes struct {
	g.Meta   `mime:"application/json"`
	Articles []entity.Article `json:"articles"`
}
