package data

import (
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"github.com/segmentio/kafka-go"
	"github.com/yogerhub/kratos-news-system/app/news/service/internal/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewDB, NewArticleRepo, NewCommentRepo, NewKafkaProducer)

// Data .
type Data struct {
	db  *gorm.DB
	rdb *redis.Client
	kp  *kafka.Writer
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, db *gorm.DB, kp *kafka.Writer) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:         c.Redis.Addr,
		WriteTimeout: c.Redis.WriteTimeout.AsDuration(),
		ReadTimeout:  c.Redis.ReadTimeout.AsDuration(),
	})

	return &Data{db: db, rdb: rdb, kp: kp}, cleanup, nil
}

func NewDB(c *conf.Data) *gorm.DB {
	db, err := gorm.Open(mysql.Open(c.Database.Source), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic("failed to connect database")
	}
	if err := db.AutoMigrate(
		&Article{},
		&Comment{},
	); err != nil {
		panic(err)
	}
	fmt.Println(err)
	return db
}

func NewKafkaProducer(conf *conf.Data) *kafka.Writer {
	w := &kafka.Writer{
		Topic:    "kns-create-comment",
		Addr:     kafka.TCP(conf.Kafka.Addrs...),
		Balancer: &kafka.LeastBytes{},
	}
	return w
}
