package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	pb "github.com/yogerhub/kratos-news-system/api/news/v1"
	biz2 "github.com/yogerhub/kratos-news-system/app/news/service/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewNewsService)

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
