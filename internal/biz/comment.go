package biz

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

type Comment struct {
	ID        int64
	UserId    int64
	ArticleId int64
	Content   string
	CreatedAt time.Time
}

type CommentRepo interface {
	AddComment(ctx context.Context, comment *Comment) (*Comment, error)
	GetComments(ctx context.Context, articleId int64) ([]*Comment, error)
	DeleteComment(ctx context.Context, id int64) error
}

type CommentUsecase struct {
	repo CommentRepo
}

func NewCommentUsecase(repo CommentRepo, logger log.Logger) *CommentUsecase {
	return &CommentUsecase{repo: repo}
}

func (uc *CommentUsecase) AddComment(ctx context.Context, comment *Comment) (p *Comment, err error) {
	return uc.repo.AddComment(ctx, comment)
}

func (uc *CommentUsecase) GetComments(ctx context.Context, articleId int64) ([]*Comment, error) {
	log.Info("comment list:", articleId)
	return uc.repo.GetComments(ctx, articleId)
}

func (uc *CommentUsecase) DeleteComment(ctx context.Context, id int64) error {
	return uc.repo.DeleteComment(ctx, id)
}
