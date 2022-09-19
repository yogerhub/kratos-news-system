package main

import (
	"context"
	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	consulAPI "github.com/hashicorp/consul/api"
	v1 "kratos-news-system/api/news/v1"
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
		kgrpc.WithEndpoint("discovery:///news.base.service"),
		kgrpc.WithMiddleware(
			recovery.Recovery(),
		),
	)

	if err != nil {
		panic(err)
	}
	defer conn.Close()

	callGetUserByPhoneGRPC(conn)
	callGetArticleGRPC(conn)
	callGetCommentsGRPC(conn)
}
func callGetUserByPhoneGRPC(conn *grpc.ClientConn) {
	client := v1.NewNewsClient(conn)
	reply, err := client.GetUserByPhone(context.Background(), &v1.GetUserByPhoneRequest{Phone: "13211112222"})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("[grpc] GetUserByPhone", reply, err)

	if errors.IsBadRequest(err) {
		log.Printf("[grpc] GetUserByPhone error is invalid argument: %v\n", err)
	}
}

func callGetArticleGRPC(conn *grpc.ClientConn) {
	client := v1.NewNewsClient(conn)
	reply, err := client.GetArticle(context.Background(), &v1.GetArticleRequest{Id: 1})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("[grpc] GetArticle", reply, err)

	if errors.IsBadRequest(err) {
		log.Printf("[grpc] GetArticle error is invalid argument: %v\n", err)
	}
}

func callGetCommentsGRPC(conn *grpc.ClientConn) {
	client := v1.NewNewsClient(conn)
	reply, err := client.GetComments(context.Background(), &v1.GetCommentRequest{ArticleId: 1})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("[grpc] GetComments", reply, err)

	if errors.IsBadRequest(err) {
		log.Printf("[grpc] GetComments error is invalid argument: %v\n", err)
	}
}
