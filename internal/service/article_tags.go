package service

import (
	"context"

	"github.com/2720851545/realworld-golang-gf/internal/service/internal/dao"
	"github.com/2720851545/realworld-golang-gf/internal/service/internal/do"
	"github.com/gogf/gf/v2/frame/g"
)

func SaveTags(ctx context.Context, id interface{}, tagList []string) (err error) {
	//  循环插入
	for _, tag := range tagList {
		var tagCount int
		if tagCount, err = dao.Tag.Ctx(ctx).Where("tag = ?", tag).Count(); err != nil {
			return
		} else if tagCount == 0 {
			_, err = dao.Tag.Ctx(ctx).Insert(do.Tag{
				Tag: tag,
			})
		}

		if tagId, err := dao.Tag.Ctx(ctx).Fields("id").Where("tag = ?", tag).Value(); err == nil {
			dao.ArticleTags.Ctx(ctx).Data(g.Map{
				"ArticleModelId": id,
				"TagModelId":     tagId,
			}).Save()
		}
	}
	return
}
