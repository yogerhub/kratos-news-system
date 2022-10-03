package biz

import (
	"github.com/go-kratos/kratos/v2/log"
)

type DemoRepo interface {
}

type DemoUsecase struct {
	repo DemoRepo
}

func NewDemoUsecase(repo DemoRepo, logger log.Logger) *DemoUsecase {
	return &DemoUsecase{repo: repo}
}
