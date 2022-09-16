package data

import (
	"context"
	"kratos-news-system/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type commentRepo struct {
	data *Data
	log  *log.Helper
}

// NewCommentRepo .
func NewCommentRepo(data *Data, logger log.Logger) biz.CommentRepo {
	return &commentRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *commentRepo) AddComment(ctx context.Context, comment *biz.Comment) (*biz.Comment, error) {
	return comment, nil
}

func (r *commentRepo) GetComments(ctx context.Context, articleId int64) ([]*biz.Comment, error) {
	rs := make([]*biz.Comment, 0)
	return rs, nil
}

func (r *commentRepo) DeleteComment(ctx context.Context, id int64) error {
	return nil
}
