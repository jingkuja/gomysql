package handle

import (
	"gotest/domain/respmsg"
	"gotest/domain/user"
	"gotest/log"

	restful "github.com/emicklei/go-restful/v3"
)

var mylog = log.NewFlieLogger("debug", "./", 100*1024*1024)

func PostUser(req *restful.Request, resp *restful.Response) {
	u := &user.User{}
	remsg := &respmsg.Resp{Status: 200, Msg: "success"}
	req.ReadEntity(u)
	_, err := user.Post(u)
	if err != nil {
		mylog.Info(err.Error())
		remsg.Status = 503
		remsg.Msg = err.Error()
		resp.WriteAsJson(remsg)
		return
	}
	resp.WriteAsJson(remsg)
}

func GetUser(req *restful.Request, resp *restful.Response) {
	uid := req.PathParameter("uid")
	res := &respmsg.Resp{Status: 200, Msg: "success"}
	rs, err := user.Get(uid)
	if err != nil {
		mylog.Info(err.Error())
		res.Status = 503
		res.Msg = err.Error()
	}
	remsg := &user.UserResp{Resp: res, Data: rs}
	resp.WriteAsJson(remsg)
}

func GetUsers(req *restful.Request, resp *restful.Response) {
	var pram user.UserPram
	req.ReadEntity(&pram)
	res := &respmsg.Resp{Status: 200, Msg: "success"}
	rs, err := user.Getrows(pram.Offset, pram.Limit)
	if err != nil {
		mylog.Info(err.Error())
		res.Status = 503
		res.Msg = err.Error()
	}
	remsg := &user.UsersResp{Resp: res, Data: rs}
	resp.WriteAsJson(remsg)
}
