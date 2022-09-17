package service

import (
	"context"
	"google.golang.org/protobuf/types/known/timestamppb"
	"kratos-news-system/internal/biz"
	"time"

	pb "kratos-news-system/api/news/v1"
)

func (s *NewsService) AddComment(ctx context.Context, req *pb.AddCommentRequest) (*pb.AddCommentReply, error) {
	co := &biz.Comment{
		UserId:    req.UserId,
		ArticleId: req.ArticleId,
		Content:   req.Content,
		CreatedAt: time.Now(),
	}
	p, err := s.comment.AddComment(ctx, co)
	if err != nil {
		return nil, err
	}

	return &pb.AddCommentReply{
		Comment: &pb.Comment{
			Id:        p.ID,
			UserId:    p.UserId,
			ArticleId: p.ArticleId,
			Content:   p.Content,
			CreatedAt: timestamppb.New(p.CreatedAt),
		},
	}, nil
}
func (s *NewsService) GetComments(ctx context.Context, req *pb.GetCommentRequest) (*pb.GetCommentReply, error) {
	comments, err := s.comment.GetComments(ctx, req.ArticleId)
	if err != nil {
		return nil, err
	}
	res := &pb.GetCommentReply{}
	for _, comment := range comments {
		co := &pb.Comment{
			Id:        comment.ID,
			UserId:    comment.UserId,
			ArticleId: comment.ArticleId,
			Content:   comment.Content,
			CreatedAt: timestamppb.New(comment.CreatedAt),
		}
		res.Comment = append(res.Comment, co)
	}
	return res, nil
}
func (s *NewsService) DeleteComment(ctx context.Context, req *pb.DeleteCommentRequest) (*pb.DeleteCommentReply, error) {
	err := s.comment.DeleteComment(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteCommentReply{}, nil
}
