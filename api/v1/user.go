package v1

import "github.com/gogf/gf/v2/frame/g"

type UserRegisterReq struct {
	g.Meta `path:"/" method:"post" summary:"用户注册"`
	User   struct {
		Username string `p:"username" v:"required#用户名不能为空"`
		Password string `p:"password" v:"required#密码不能为空"`
		Email    string `p:"email" v:"required#邮箱不能为空"`
		Image    string `p:"image"`
	}
}

type UserRegisterRes struct {
	g.Meta `mime:"application/json"`
	User   struct {
		Email    string `json:"email"`
		Username string `json:"username"`
		Bio      string `json:"bio"`
		Image    string `json:"image"`
		Token    string `json:"token"`
	} `json:"user"`
}
