package service

import (
	"testing"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

func TestGconvStruct(t *testing.T) {
	type User struct {
		Uid  int
		Name string
	}
	params := g.Map{
		"uid":  1,
		"name": "john",
	}
	var user, user1 *User
	if err := gconv.Struct(params, &user); err != nil {
		panic(err)
	}
	if err := gconv.Struct(&user, &user1); err != nil {
		panic(err)
	}
	g.Dump(user)
	g.Dump("----------")
	g.Dump(user1)
}
