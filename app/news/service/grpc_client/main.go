package main

import (
	"context"
	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	consulAPI "github.com/hashicorp/consul/api"
	newsCli "github.com/yogerhub/kratos-news-system/api/news/v1"
	"log"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	kgrpc "github.com/go-kratos/kratos/v2/transport/grpc"

	"google.golang.org/grpc"
)

func main() {

	c := consulAPI.DefaultConfig()
	c.Address = "127.0.0.1:8500"
	c.Scheme = "http"
	cli, err := consulAPI.NewClient(c)
	if err != nil {
		panic(err)
	}
	r := consul.New(cli, consul.WithHealthCheck(false))

	conn, err := kgrpc.DialInsecure(
		context.Background(),
		kgrpc.WithDiscovery(r),
		kgrpc.WithEndpoint("discovery:///kns.news.service"),
		kgrpc.WithMiddleware(
			recovery.Recovery(),
		),
	)

	if err != nil {
		panic(err)
	}
	defer conn.Close()

	callGetArticleGRPC(conn)
	callGetCommentsGRPC(conn)
}

func callGetArticleGRPC(conn *grpc.ClientConn) {
	client := newsCli.NewNewsClient(conn)
	reply, err := client.GetArticle(context.Background(), &newsCli.GetArticleRequest{Id: 1})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("[grpc] GetArticle", reply, err)

	if errors.IsBadRequest(err) {
		log.Printf("[grpc] GetArticle error is invalid argument: %v\n", err)
	}
}

func callGetCommentsGRPC(conn *grpc.ClientConn) {
	client := newsCli.NewNewsClient(conn)
	reply, err := client.GetComments(context.Background(), &newsCli.GetCommentRequest{ArticleId: 1})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("[grpc] GetComments", reply, err)

	if errors.IsBadRequest(err) {
		log.Printf("[grpc] GetComments error is invalid argument: %v\n", err)
	}
}
