package v1

import "github.com/gogf/gf/v2/frame/g"

type userRegisterReqUser struct {
	Username string `p:"username" json:"username" v:"required#用户名不能为空"`
	Password string `p:"password" json:"password" v:"required#密码不能为空"`
	Email    string `p:"email" json:"email" v:"required#邮箱不能为空"`
	Image    string `p:"image" json:"image"`
}

type UserRegisterReq struct {
	g.Meta `path:"/users" method:"post" tags:"用户" summary:"用户注册"`
	User   *userRegisterReqUser `json:"user"`
}

type userRegisterResUser struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Bio      string `json:"bio"`
	Image    string `json:"image"`
	Token    string `json:"token"`
}

type UserRegisterRes struct {
	g.Meta `mime:"application/json"`
	User   userRegisterResUser `json:"user"`
}

type CurrentUserReq struct {
	g.Meta `path:"/user" method:"get" tags:"用户" summary:"当前用户"`
}

type CurrentUserRes struct {
	g.Meta `mime:"application/json"`
}
