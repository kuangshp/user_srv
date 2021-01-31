package handler

import (
	"context"
	"fmt"
	"github.com/kuangshp/user_srv/proto/user"
)

type User struct{}

func (u *User) Register(ctx context.Context, in *user.RegisterReq, out *user.RegisterRes) error {
	username := in.Username
	email := in.Email
	password := in.Password
	fmt.Println("接受到的数据:", username, email, password)
	out.Code = 1
	out.Message = "成功"
	return nil
}
