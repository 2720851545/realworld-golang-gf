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
	User   userRegisterReqUser `json:"user"`
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
	User   currentUserResUser `json:"user"`
}

type currentUserResUser struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Bio      string `json:"bio"`
	Image    string `json:"image"`
	Token    string `json:"token"`
}

type userLoginReqUser struct {
	Email    string `p:"email" json:"email" v:"required#邮箱不能为空"`
	Password string `p:"password" json:"password" v:"required#密码不能为空"`
}

type UserLoginReq struct {
	g.Meta `path:"/users/login" method:"post" tags:"用户" summary:"用户登陆"`
	User   userLoginReqUser `json:"user"`
}

type ProfileReq struct {
	g.Meta   `path:"/profiles/:username" method:"get" tags:"用户" summary:"用户个人资料"`
	Username string `json:"username"`
}

type ProfileRes struct {
	g.Meta  `mime:"application/json"`
	Profile ProfileResProfile `json:"profile"`
}

type FollowProfileReq struct {
	g.Meta   `path:"/profiles/:username/follow" method:"post" tags:"用户" summary:"关注用户"`
	Username string               `p:"username" v:"required#标题不能为空"`
	User     FollowProfileReqUser `json:"user"`
}

type FollowProfileRes struct {
	g.Meta  `mime:"application/json"`
	Profile allArticleResArticleAuthor `json:"profile"`
}

type UnFollowProfileReq struct {
	g.Meta   `path:"/profiles/:username/follow" method:"delete" tags:"用户" summary:"取消关注用户"`
	Username string `p:"username" v:"required#标题不能为空"`
}

type UnFollowProfileRes struct {
	g.Meta  `mime:"application/json"`
	Profile allArticleResArticleAuthor `json:"profile"`
}

type FollowProfileReqUser struct {
	Email string `json:"email" v:"required#邮箱不能为空"`
}

type ProfileResProfile struct {
	Id        int    `json:"id"`
	Username  string `json:"username"`
	Bio       string `json:"bio"`
	Image     string `json:"image"`
	Following bool   `json:"following"`
}

type userLoginResUser struct {
	Id       int64  `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Bio      string `json:"bio"`
	Image    string `json:"image"`
	Token    string `json:"token"`
}

type UserLoginRes struct {
	g.Meta `mime:"application/json"`
	User   userLoginResUser `json:"user"`
}

type userUpdateReqUser struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Bio      string `json:"bio"`
	Image    string `json:"image"`
}

type UserUpdateReq struct {
	g.Meta `path:"/user" method:"put" tags:"用户" summary:"更新用户信息"`
	User   userUpdateReqUser `json:"user"`
}

type userUpdateResUser struct {
	Id       int64  `json:"id" `
	Email    string `json:"email"`
	Username string `json:"username"`
	Bio      string `json:"bio"`
	Image    string `json:"image"`
	Token    string `json:"token"`
}

type UserUpdateRes struct {
	g.Meta `mime:"application/json"`
	User   userUpdateResUser `json:"user"`
}
