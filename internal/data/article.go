package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"kratos-news-system/internal/biz"
)

type Article struct {
	gorm.Model
	Title   string `json:"title"`
	Content string `json:"content"`
}

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
	ar := new(Article)
	err := r.data.db.Where("id = ?", id).First(ar).Error
	if err != nil {
		return nil, err
	}
	return &biz.Article{
		ID:        int64(ar.ID),
		Title:     ar.Title,
		Content:   ar.Content,
		CreatedAt: ar.CreatedAt,
		UpdatedAt: ar.UpdatedAt,
	}, nil
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
