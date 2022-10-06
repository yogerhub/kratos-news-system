package service

import (
	"context"
	"fmt"
	consul "github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"
	consulAPI "github.com/hashicorp/consul/api"
	"github.com/segmentio/kafka-go"
	pb "github.com/yogerhub/kratos-news-system/api/filter/v1"
	"github.com/yogerhub/kratos-news-system/app/filter/service/internal/biz"
	"github.com/yogerhub/kratos-news-system/app/filter/service/internal/conf"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewFilterService, NewRegistrar, NewKafkaConsumerJobServer)

type FilterService struct {
	pb.UnimplementedFilterServer
	log    *log.Helper
	filter *biz.FilterUsecase
}

func NewFilterService(filter *biz.FilterUsecase, logger log.Logger) *FilterService {
	return &FilterService{
		filter: filter,
		log:    log.NewHelper(logger),
	}
}

func NewRegistrar(conf *conf.Registry) registry.Registrar {
	c := consulAPI.DefaultConfig()
	c.Address = conf.Consul.Address
	c.Scheme = conf.Consul.Scheme
	cli, err := consulAPI.NewClient(c)
	if err != nil {
		panic(err)
	}
	log.Info("consul register")
	r := consul.New(cli, consul.WithHealthCheck(true))
	return r
}

type KafkaConsumerJobServer struct {
	kc *kafka.Reader
}

func NewKafkaConsumerJobServer(conf *conf.Data) *KafkaConsumerJobServer {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  conf.Kafka.Addrs,
		GroupID:  "group-a",
		Topic:    "kns-create-comment",
		MinBytes: 1,    // 1
		MaxBytes: 10e6, // 10MB
	})
	return &KafkaConsumerJobServer{kc: r}
}

func (s *KafkaConsumerJobServer) Start(ctx context.Context) error {
	for {
		m, err := s.kc.ReadMessage(ctx)
		if err != nil {
			break
		}
		//敏感词过滤 todo
		log.Info(fmt.Sprintf("message at offset %d: %s = %s", m.Offset, string(m.Key), string(m.Value)))
	}
	if err := s.kc.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}
	return nil
}

func (s *KafkaConsumerJobServer) Stop(ctx context.Context) error {
	ctx.Done()
	return nil
}
