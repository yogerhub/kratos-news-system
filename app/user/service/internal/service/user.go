package service

import (
	"context"
	"google.golang.org/protobuf/types/known/timestamppb"
	"kratos-news-system/app/user/service/internal/biz"
	"time"

	pb "kratos-news-system/api/user/v1"
)

func (s *NewsService) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.UserReply, error) {
	user := &biz.User{
		Username:  req.Username,
		Phone:     req.Phone,
		Password:  req.Password,
		CreatedAt: time.Now(),
	}
	res, err := s.user.Register(ctx, user)
	if err != nil {
		return nil, err
	}

	return &pb.UserReply{User: &pb.UserInfo{
		Id:        res.ID,
		Username:  res.Username,
		Phone:     res.Phone,
		Password:  res.Password,
		CreatedAt: timestamppb.New(res.CreatedAt),
	}}, err
}
func (s *NewsService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.UserReply, error) {
	user := &biz.User{
		Username: req.Username,
		Password: req.Password,
	}

	res, err := s.user.Login(ctx, user)
	if err != nil {
		return nil, err
	}

	return &pb.UserReply{User: &pb.UserInfo{
		Id:        res.ID,
		Username:  res.Username,
		Phone:     res.Phone,
		Password:  res.Password,
		CreatedAt: timestamppb.New(res.CreatedAt),
	}}, err

}
func (s *NewsService) GetUserByPhone(ctx context.Context, req *pb.GetUserByPhoneRequest) (*pb.UserReply, error) {

	res, err := s.user.GetUserByPhone(ctx, req.Phone)
	if err != nil {
		return nil, err
	}

	return &pb.UserReply{User: &pb.UserInfo{
		Id:        res.ID,
		Username:  res.Username,
		Phone:     res.Phone,
		Password:  res.Password,
		CreatedAt: timestamppb.New(res.CreatedAt),
	}}, nil
}
func (s *NewsService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UserReply, error) {
	user := &biz.User{
		Username:  req.Username,
		Phone:     req.Phone,
		Password:  req.Password,
		UpdatedAt: time.Now(),
	}
	res, err := s.user.UpdateUser(ctx, req.Id, user)
	if err != nil {
		return nil, err
	}

	return &pb.UserReply{User: &pb.UserInfo{
		Id:        res.ID,
		Username:  res.Username,
		Phone:     res.Phone,
		Password:  res.Password,
		CreatedAt: timestamppb.New(res.CreatedAt),
	}}, err

}
