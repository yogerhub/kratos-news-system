package service

import (
	"context"

	pb "kratos-news-system/api/news/v1"
)

func (s *NewsService) AddComment(ctx context.Context, req *pb.AddCommentRequest) (*pb.AddCommentReply, error) {
	return &pb.AddCommentReply{}, nil
}
func (s *NewsService) GetComments(ctx context.Context, req *pb.GetCommentRequest) (*pb.GetCommentReply, error) {
	return &pb.GetCommentReply{}, nil
}
func (s *NewsService) DeleteComment(ctx context.Context, req *pb.DeleteCommentRequest) (*pb.DeleteCommentReply, error) {
	return &pb.DeleteCommentReply{}, nil
}
