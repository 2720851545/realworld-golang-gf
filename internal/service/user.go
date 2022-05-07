package service

import (
	"context"
	"time"

	v1 "github.com/2720851545/realworld-golang-gf/api/v1"
	"github.com/2720851545/realworld-golang-gf/internal/service/internal/dao"
	"github.com/2720851545/realworld-golang-gf/utility"
	jwt "github.com/gogf/gf-jwt/v2"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type IUserService interface {
	// Register(ctx context.Context, req *v1.UserRegisterReq) (i int64, err error)
	Register(ctx context.Context, req *v1.UserRegisterReq) (res *v1.UserRegisterRes, err error)
	CurrentUser(ctx context.Context, req *v1.CurrentUserReq) (res *v1.CurrentUserRes, err error)
	Login(ctx context.Context, req *v1.UserLoginReq) (res *v1.UserLoginRes, err error)
	UpdateUserInfo(ctx context.Context, req *v1.UserUpdateReq) (res *v1.UserUpdateRes, err error)
}

type userImpl struct{}

func UserService() IUserService {
	return IUserService(&userImpl{})
}

func (s *userImpl) CurrentUser(ctx context.Context, req *v1.CurrentUserReq) (res *v1.CurrentUserRes, err error) {
	var (
		token string
		mc    jwt.MapClaims
	)
	mc, token, err = authService.GetClaimsFromJWT(ctx)
	if err != nil {
		return
	}
	if id, ok := mc["id"]; ok {
		res = new(v1.CurrentUserRes)
		res.User.Token = token
		dao.User.Ctx(ctx).Where("id = ?", id).Scan(&res.User)
	} else {
		err = gerror.New("jwt 解析失败, 没有id字段")
	}
	return
}

func (s *userImpl) Register(ctx context.Context, req *v1.UserRegisterReq) (res *v1.UserRegisterRes, err error) {
	req.User.Password = utility.EntryPassword(req.User.Password)
	if g.IsEmpty(req.User.Image) {
		req.User.Image = "https://api.realworld.io/images/smiley-cyrus.jpeg"
	}

	var id int64
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		err = g.Try(func() {
			id, err = dao.User.Ctx(ctx).TX(tx).InsertAndGetId(req.User)
			if err != nil {
				panic(err)
			}
			g.Log().Info(ctx, "新用户的id=", id)
			res = new(v1.UserRegisterRes)
			err = dao.User.Ctx(ctx).Where("id = ?", id).Scan(&res.User)

			if err == nil {
				res.User.Token, _ = getLoginToken(ctx, id, res.User.Username)
			}
		})
		return err
	})

	return res, err
}

func (s *userImpl) UpdateInfo(ctx context.Context, req *v1.UserRegisterReq) (res *v1.UserRegisterRes, err error) {
	req.User.Password = utility.EntryPassword(req.User.Password)
	if g.IsEmpty(req.User.Image) {
		req.User.Image = "https://api.realworld.io/images/smiley-cyrus.jpeg"
	}

	var id int64
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		err = g.Try(func() {
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

	if err == nil {
		res.User.Token, _ = getLoginToken(ctx, id, res.User.Username)
	}
	return res, err
}

func getLoginToken(ctx context.Context, id int64, username string) (token string, expire time.Time) {
	return getAuthToken(ctx, "Register", map[string]interface{}{
		"id":       id,
		"username": username,
	})
}

func getAuthToken(ctx context.Context, model string, data map[string]interface{}) (token string, expire time.Time) {
	r := g.RequestFromCtx(ctx)
	r.SetCtxVar("Model", model)
	r.SetCtxVar("User", data)

	token, expire = authService.LoginHandler(ctx)
	return
}

func (s *userImpl) UpdateUserInfo(ctx context.Context, req *v1.UserUpdateReq) (res *v1.UserUpdateRes, err error) {
	mc, _, err := authService.GetClaimsFromJWT(ctx)
	if err != nil {
		return
	}
	id := gconv.Int64(mc["id"])
	_, err = dao.User.Ctx(ctx).OmitEmptyData().Where("id = ?", id).Update(req.User)
	res = new(v1.UserUpdateRes)
	dao.User.Ctx(ctx).Where("id = ?", id).Scan(&res.User)
	res.User.Token, _ = getLoginToken(ctx, id, res.User.Username)
	return
}

func (s *userImpl) Login(ctx context.Context, req *v1.UserLoginReq) (res *v1.UserLoginRes, err error) {
	req.User.Password = utility.EntryPassword(req.User.Password)
	res = new(v1.UserLoginRes)
	err = dao.User.Ctx(ctx).Where("email = ? and password = ? ", req.User.Email, req.User.Password).Scan(&res.User)
	if err != nil {
		return nil, gerror.New("用户名或密码错误")
	}

	res.User.Token, _ = getLoginToken(ctx, res.User.Id, res.User.Username)
	return res, nil
}

func isFollowingUser(ctx context.Context, following, followedByID int) bool {
	count, _ := dao.Follow.Ctx(ctx).Count("following_id = ? and followed_by_id = ?", following, followedByID)
	return count > 0
}

func GetUserId(ctx context.Context) interface{} {
	return gconv.Map(g.RequestFromCtx(ctx).Get("JWT_PAYLOAD"))["id"]
}
