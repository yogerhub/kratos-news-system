package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	pb "github.com/yogerhub/kratos-news-system/api/user/v1"
	biz2 "github.com/yogerhub/kratos-news-system/app/user/service/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewUserService)

type UserService struct {
	pb.UnimplementedUserServer
	log  *log.Helper
	user *biz2.UserUsecase
}

func NewUserService(user *biz2.UserUsecase, logger log.Logger) *UserService {
	return &UserService{
		user: user,
		log:  log.NewHelper(logger),
	}
}
