package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/yogerhub/kratos-news-system/app/filter/service/internal/biz"
)

type Filter struct {
}

type filterRepo struct {
	data *Data
	log  *log.Helper
}

// NewFilterRepo .
func NewFilterRepo(data *Data, logger log.Logger) biz.FilterRepo {
	return &filterRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
