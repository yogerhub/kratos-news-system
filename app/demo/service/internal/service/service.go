package service

import (
	consul "github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"
	consulAPI "github.com/hashicorp/consul/api"
	pb "github.com/yogerhub/kratos-news-system/api/demo/v1"
	"github.com/yogerhub/kratos-news-system/app/demo/service/internal/biz"
	"github.com/yogerhub/kratos-news-system/app/demo/service/internal/conf"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewDemoService, NewRegistrar)

type DemoService struct {
	pb.UnimplementedDemoServer
	log  *log.Helper
	demo *biz.DemoUsecase
}

func NewDemoService(demo *biz.DemoUsecase, logger log.Logger) *DemoService {
	return &DemoService{
		demo: demo,
		log:  log.NewHelper(logger),
	}
}

func NewRegistrar(conf *conf.Registry) registry.Registrar {
	c := consulAPI.DefaultConfig()
	c.Address = conf.Consul.Address
	c.Scheme = conf.Consul.Scheme
	cli, err := consulAPI.NewClient(c)
	if err != nil {
		panic(err)
	}
	log.Info("consul register")
	r := consul.New(cli, consul.WithHealthCheck(true))
	return r
}