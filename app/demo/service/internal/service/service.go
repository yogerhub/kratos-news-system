package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	pb "github.com/yogerhub/kratos-news-system/api/demo/v1"
	"github.com/yogerhub/kratos-news-system/app/demo/service/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewDemoService)

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
