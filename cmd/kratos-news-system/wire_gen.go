// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-news-system/internal/biz"
	"kratos-news-system/internal/conf"
	"kratos-news-system/internal/data"
	"kratos-news-system/internal/server"
	"kratos-news-system/internal/service"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	db := data.NewDB(confData)
	dataData, cleanup, err := data.NewData(confData, logger, db)
	if err != nil {
		return nil, nil, err
	}
	articleRepo := data.NewArticleRepo(dataData, logger)
	articleUsecase := biz.NewArticleUsecase(articleRepo, logger)
	userRepo := data.NewUserRepo(dataData, logger)
	userUsecase := biz.NewUserUsecase(userRepo, logger)
	commentRepo := data.NewCommentRepo(dataData, logger)
	commentUsecase := biz.NewCommentUsecase(commentRepo, logger)
	newsService := service.NewNewsService(articleUsecase, userUsecase, commentUsecase, logger)
	grpcServer := server.NewGRPCServer(confServer, newsService, logger)
	httpServer := server.NewHTTPServer(confServer, newsService, logger)
	app := newApp(logger, grpcServer, httpServer)
	return app, func() {
		cleanup()
	}, nil
}
