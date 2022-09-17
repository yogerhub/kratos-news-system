package service

import (
	"context"
	"google.golang.org/protobuf/types/known/timestamppb"
	pb "kratos-news-system/api/news/v1"
	"log"
)

func (s *NewsService) CreateArticle(ctx context.Context, req *pb.CreateArticleRequest) (*pb.CreateArticleReply, error) {
	return &pb.CreateArticleReply{}, nil
}
func (s *NewsService) UpdateArticle(ctx context.Context, req *pb.UpdateArticleRequest) (*pb.UpdateArticleReply, error) {
	return &pb.UpdateArticleReply{}, nil
}
func (s *NewsService) DeleteArticle(ctx context.Context, req *pb.DeleteArticleRequest) (*pb.DeleteArticleReply, error) {
	return &pb.DeleteArticleReply{}, nil
}
func (s *NewsService) GetArticle(ctx context.Context, req *pb.GetArticleRequest) (*pb.GetArticleReply, error) {
	p, err := s.article.Get(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	log.Println(p, err)
	return &pb.GetArticleReply{
		Article: &pb.Article{
			Id:        p.ID,
			Title:     p.Title,
			Content:   p.Content,
			CreatedAt: timestamppb.New(p.CreatedAt),
			UpdatedAt: timestamppb.New(p.UpdatedAt)},
	}, nil
}
func (s *NewsService) ListArticle(ctx context.Context, req *pb.ListArticleRequest) (*pb.ListArticleReply, error) {
	return &pb.ListArticleReply{}, nil
}
