package biz

import (
	"github.com/go-kratos/kratos/v2/log"
)

type FilterRepo interface {
}

type FilterUsecase struct {
	repo FilterRepo
}

func NewFilterUsecase(repo FilterRepo, logger log.Logger) *FilterUsecase {
	return &FilterUsecase{repo: repo}
}
