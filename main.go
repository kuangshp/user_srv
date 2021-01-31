package main

import (
	"github.com/kuangshp/user_srv/handler"
	"github.com/kuangshp/user_srv/proto/user"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
)

func main() {
	consulRegistry := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string {
			"127.0.0.1:8500",
		}
	})
	srv := micro.NewService(
		micro.Name("go.micro.service.user"),
		micro.Version("latest"),
		micro.Address("127.0.0.1:8083"),
		micro.Registry(consulRegistry),
	)
	srv.Init()
	// Register handler
	err1 := user.RegisterUserHandler(srv.Server(), new(handler.User))
	if err1 != nil {
		logger.Error(err1)
	}
	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
