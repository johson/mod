package main

import (
	"github.com/asim/go-micro/v3"
	k8s "github.com/micro/examples/kubernetes/go/micro"
	// k8s "github.com/micro/kubernetes/go/micro"
)

//TODO: 重构配置
func main() {

	// router := gin.Default()
	// API.InitRoute(router)
	if true {
		service := micro.NewService(micro.Name("hardware"), micro.Address(":8080"))
	} else {
		service = k8s.NewService(micro.Name("hardware"))
	}

	// service.Handle("/", router)
	service.Init()
	service.Run()
}
