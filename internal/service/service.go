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
	user    *biz.UserUsecase
	article *biz.ArticleUsecase
	comment *biz.CommentUsecase
}

func NewNewsService(article *biz.ArticleUsecase, user *biz.UserUsecase, comment *biz.CommentUsecase, logger log.Logger) *NewsService {
	return &NewsService{
		article: article,
		user:    user,
		comment: comment,
		log:     log.NewHelper(logger),
	}
}
