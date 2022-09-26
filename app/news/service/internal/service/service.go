package service

import (
	consul "github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"
	consulAPI "github.com/hashicorp/consul/api"
	pb "github.com/yogerhub/kratos-news-system/api/news/v1"
	biz2 "github.com/yogerhub/kratos-news-system/app/news/service/internal/biz"
	"github.com/yogerhub/kratos-news-system/app/news/service/internal/conf"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewNewsService, NewRegistrar)

type NewsService struct {
	pb.UnimplementedNewsServer
	log     *log.Helper
	article *biz2.ArticleUsecase
	comment *biz2.CommentUsecase
}

func NewNewsService(article *biz2.ArticleUsecase, comment *biz2.CommentUsecase, logger log.Logger) *NewsService {
	return &NewsService{
		article: article,
		comment: comment,
		log:     log.NewHelper(logger),
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
