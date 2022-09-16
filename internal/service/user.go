package service

import (
	"context"

	pb "kratos-news-system/api/news/v1"
)

func (s *NewsService) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterRequest, error) {
	return &pb.RegisterRequest{}, nil
}
func (s *NewsService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.UserReply, error) {
	return &pb.UserReply{}, nil
}
func (s *NewsService) GetUserByPhone(ctx context.Context, req *pb.GetUserByPhoneRequest) (*pb.UserReply, error) {
	return &pb.UserReply{}, nil
}
func (s *NewsService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UserReply, error) {
	return &pb.UserReply{}, nil
}
