package user

import "gotest/domain/respmsg"

type UserResp struct {
	*respmsg.Resp
	Data User `json:"data"`
}

type UsersResp struct {
	*respmsg.Resp
	Data []User `json:"data"`
}
