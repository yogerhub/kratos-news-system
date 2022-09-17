package data

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func userKey(id int64) string {
	return fmt.Sprintf("user:%d", id)
}

func (ar *articleRepo) GetUser(ctx context.Context, id int64) (rv int64, err error) {
	get := ar.data.rdb.Get(ctx, userKey(id))
	rv, err = get.Int64()
	if err == redis.Nil {
		return 0, nil
	}
	return
}
