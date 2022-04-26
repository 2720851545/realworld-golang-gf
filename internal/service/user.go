package service

import (
	"context"

	v1 "github.com/2720851545/realworld-golang-gf/api/v1"
	"github.com/2720851545/realworld-golang-gf/internal/service/internal/dao"
	"github.com/2720851545/realworld-golang-gf/utility"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

type IUserService interface {
	// Register(ctx context.Context, req *v1.UserRegisterReq) (i int64, err error)
	Register(ctx context.Context, req *v1.UserRegisterReq) (res *v1.UserRegisterRes, err error)
}

type userImpl struct{}

func UserService() IUserService {
	return IUserService(&userImpl{})
}

func (s *userImpl) Register(ctx context.Context, req *v1.UserRegisterReq) (res *v1.UserRegisterRes, err error) {
	req.User.Password = utility.EntryPassword(req.User.Password)
	if g.IsEmpty(req.User.Image) {
		req.User.Image = "https://api.realworld.io/images/smiley-cyrus.jpeg"
	}

	err = g.DB().Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		err = g.Try(func() {
			var id int64
			id, err = dao.User.Ctx(ctx).TX(tx).InsertAndGetId(req.User)
			if err != nil {
				panic(err)
			}
			g.Log().Info(ctx, "新用户的id=", id)
			res = new(v1.UserRegisterRes)
			err = dao.User.Ctx(ctx).Where("id = ?", id).Scan(&res.User)
		})
		return err
	})
	return res, err
}
