package data

import (
	"context"
	"kratos-news-system/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type articleRepo struct {
	data *Data
	log  *log.Helper
}

// NewArticleRepo .
func NewArticleRepo(data *Data, logger log.Logger) biz.ArticleRepo {
	return &articleRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *articleRepo) ListArticle(ctx context.Context) ([]*biz.Article, error) {
	rs := make([]*biz.Article, 0)
	return rs, nil
}

func (r *articleRepo) GetArticle(ctx context.Context, id int64) (*biz.Article, error) {
	return &biz.Article{}, nil
}

func (r *articleRepo) CreateArticle(context.Context, *biz.Article) error {
	return nil
}

func (r *articleRepo) UpdateArticle(ctx context.Context, id int64, article *biz.Article) error {
	return nil
}
func (r *articleRepo) DeleteArticle(ctx context.Context, id int64) error {
	return nil
}
