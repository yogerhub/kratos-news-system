package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/yogerhub/kratos-news-system/app/demo/service/internal/biz"
)

type Demo struct {
}

type demoRepo struct {
	data *Data
	log  *log.Helper
}

// NewDemoRepo .
func NewDemoRepo(data *Data, logger log.Logger) biz.DemoRepo {
	return &demoRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
