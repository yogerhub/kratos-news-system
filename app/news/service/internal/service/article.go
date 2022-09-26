package service

import (
	"context"
	pb "github.com/yogerhub/kratos-news-system/api/news/v1"
	"github.com/yogerhub/kratos-news-system/app/news/service/internal/biz"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"time"
)

func (s *NewsService) CreateArticle(ctx context.Context, req *pb.CreateArticleRequest) (*pb.CreateArticleReply, error) {
	ar := &biz.Article{
		Title:     req.Title,
		Content:   req.Content,
		CreatedAt: time.Now(),
	}
	p, err := s.article.Create(ctx, ar)
	if err != nil {
		return nil, err
	}
	return &pb.CreateArticleReply{
		Article: &pb.Article{
			Id:        p.ID,
			Title:     p.Title,
			Content:   p.Content,
			CreatedAt: timestamppb.New(p.CreatedAt),
			UpdatedAt: timestamppb.New(p.UpdatedAt)},
	}, nil
}
func (s *NewsService) UpdateArticle(ctx context.Context, req *pb.UpdateArticleRequest) (*pb.UpdateArticleReply, error) {
	ar := &biz.Article{
		ID:        req.Id,
		Title:     req.Title,
		Content:   req.Content,
		UpdatedAt: time.Now(),
	}
	p, err := s.article.Update(ctx, req.Id, ar)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateArticleReply{Article: &pb.Article{
		Id:        p.ID,
		Title:     p.Title,
		Content:   p.Content,
		CreatedAt: timestamppb.New(p.CreatedAt),
		UpdatedAt: timestamppb.New(p.UpdatedAt)},
	}, nil
}
func (s *NewsService) DeleteArticle(ctx context.Context, req *pb.DeleteArticleRequest) (*pb.DeleteArticleReply, error) {
	err := s.article.Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}
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
	articles, err := s.article.List(ctx)
	if err != nil {
		return nil, err
	}
	res := &pb.ListArticleReply{}
	for _, article := range articles {
		ar := &pb.Article{
			Id:        article.ID,
			Title:     article.Title,
			Content:   article.Content,
			CreatedAt: timestamppb.New(article.CreatedAt),
			UpdatedAt: timestamppb.New(article.UpdatedAt)}
		res.Results = append(res.Results, ar)
	}
	return res, nil
}
