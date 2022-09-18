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

// ListArticle 获取列表
func (r *articleRepo) ListArticle(ctx context.Context) ([]*biz.Article, error) {
	var articles []Article
	result := r.data.db.Find(&articles)
	if result.Error != nil {
		return nil, result.Error
	}
	rs := make([]*biz.Article, len(articles))

	for i, x := range articles {
		rs[i] = convertArticle(x)
	}
	return rs, nil
}

func (r *articleRepo) GetArticle(ctx context.Context, id int64) (*biz.Article, error) {
	ar := new(Article)
	exist, err := r.ExistsArticleInfo(ctx, id)
	if err != nil || exist == false {
		err := r.data.db.Where("id = ?", id).First(ar).Error
		if err != nil {
			return nil, err
		}
		_, _ = r.SetArticleInfo(ctx, ar)
		return &biz.Article{
			ID:        int64(ar.ID),
			Title:     ar.Title,
			Content:   ar.Content,
			CreatedAt: ar.CreatedAt,
			UpdatedAt: ar.UpdatedAt,
		}, nil
	} else {
		article, err := r.GetArticleInfo(ctx, id)
		if err != nil {
			return nil, err
		}
		return &biz.Article{
			ID:        int64(article.ID),
			Title:     article.Title,
			Content:   article.Content,
			CreatedAt: article.CreatedAt,
			UpdatedAt: article.UpdatedAt,
		}, nil
	}
}

func (r *articleRepo) CreateArticle(cxt context.Context, article *biz.Article) (*biz.Article, error) {
	ar := Article{
		Title:   article.Title,
		Content: article.Content,
	}

	result := r.data.db.Create(&ar) // 通过数据的指针来创建
	if result.Error != nil {
		return nil, result.Error
	}
	_, _ = r.SetArticleInfo(cxt, &ar)
	return &biz.Article{
		ID:        int64(ar.ID),
		Title:     ar.Title,
		Content:   ar.Content,
		CreatedAt: ar.CreatedAt,
		UpdatedAt: ar.UpdatedAt,
	}, nil
}

// UpdateArticle 更新文章信息
func (r *articleRepo) UpdateArticle(ctx context.Context, id int64, article *biz.Article) (*biz.Article, error) {
	var po Article
	err := r.data.db.Model(&po).Updates(article).Error
	_, _ = r.SetArticleInfo(ctx, &po)
	if err != nil {
		return nil, err
	}
	return convertArticle(po), err
}

func (r *articleRepo) DeleteArticle(ctx context.Context, id int64) error {
	rv := r.data.db.Delete(&Article{}, id)

	return rv.Error
}

func convertArticle(x Article) *biz.Article {
	return &biz.Article{
		ID:        int64(x.ID),
		Title:     x.Title,
		Content:   x.Content,
		CreatedAt: x.CreatedAt,
		UpdatedAt: x.UpdatedAt,
	}
}
