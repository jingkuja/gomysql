package container

import (
	"gotest/handle"

	restful "github.com/emicklei/go-restful/v3"
)

func GetmainCon() *restful.Container {
	container := restful.NewContainer()
	ws := new(restful.WebService)
	ws.Path("/rest")
	ws.Route(ws.POST("/user").To(handle.PostUser))
	ws.Route(ws.GET("/user/{uid}").To(handle.GetUser))
	ws.Route(ws.GET("/users").To(handle.GetUsers))
	container.Add(ws)
	return container
}
