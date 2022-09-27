//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/yogerhub/kratos-news-system/app/news/service/internal/biz"
	"github.com/yogerhub/kratos-news-system/app/news/service/internal/conf"
	"github.com/yogerhub/kratos-news-system/app/news/service/internal/data"
	"github.com/yogerhub/kratos-news-system/app/news/service/internal/server"
	"github.com/yogerhub/kratos-news-system/app/news/service/internal/service"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Registry, *conf.Data, log.Logger, *tracesdk.TracerProvider) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
