package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	pb "kratos-news-system/api/news/v1"
	"kratos-news-system/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewNewsService)

type NewsService struct {
	pb.UnimplementedNewsServer
	log     *log.Helper
	article *biz.ArticleUsecase
}

func NewNewsService(article *biz.ArticleUsecase, logger log.Logger) *NewsService {
	return &NewsService{
		article: article,
		log:     log.NewHelper(logger),
	}
}
