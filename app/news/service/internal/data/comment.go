package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"kratos-news-system/app/user/service/internal/biz"
)

type Comment struct {
	gorm.Model
	UserId    int64  `json:"user_id"`
	ArticleId int64  `json:"article_id"`
	Content   string `json:"content"`
}

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
	co := Comment{
		UserId:    comment.UserId,
		ArticleId: comment.ArticleId,
		Content:   comment.Content,
	}

	result := r.data.db.Create(&co) // 通过数据的指针来创建
	if result.Error != nil {
		return nil, result.Error
	}

	return &biz.Comment{
		ID:        int64(co.ID),
		UserId:    co.UserId,
		ArticleId: co.ArticleId,
		Content:   co.Content,
		CreatedAt: co.CreatedAt,
	}, nil
}

// GetComments 获取评论列表
func (r *commentRepo) GetComments(ctx context.Context, articleId int64) ([]*biz.Comment, error) {
	var comments []Comment
	result := r.data.db.Where("article_id = ?", articleId).Find(&comments)
	if result.Error != nil {
		return nil, result.Error
	}
	log.Info("comment list:", comments, result)
	rs := make([]*biz.Comment, len(comments))

	for i, x := range comments {
		rs[i] = &biz.Comment{
			ID:        int64(x.ID),
			UserId:    x.UserId,
			ArticleId: x.ArticleId,
			Content:   x.Content,
			CreatedAt: x.CreatedAt,
		}
	}
	return rs, nil
}

func (r *commentRepo) DeleteComment(ctx context.Context, id int64) error {
	rv := r.data.db.Delete(&Comment{}, id)

	return rv.Error
}
