package data

import (
	"context"
	"fmt"
)

func articleKey(id int64) string {
	return fmt.Sprintf("article:id:%d", id)
}

func (u *articleRepo) SetArticleInfo(ctx context.Context, article *Article) (res bool, err error) {
	args := make([]interface{}, 0)
	args = append(args, "id")
	args = append(args, article.ID)
	args = append(args, "title")
	args = append(args, article.Title)
	args = append(args, "content")
	args = append(args, article.Content)
	args = append(args, "created_at")
	args = append(args, article.CreatedAt)

	res, err = u.data.rdb.HMSet(ctx, articleKey(int64(article.ID)), args...).Result()
	fmt.Println(res, err)
	return
}

func (u *articleRepo) ExistsArticleInfo(ctx context.Context, id int64) (rv bool, err error) {
	m, err := u.data.rdb.Exists(ctx, articleKey(id)).Result()
	defer func() {
		fmt.Println(rv, err)
	}()
	if err != nil {
		return false, err
	}
	if m == 1 {
		return true, err
	}
	return
}

func (u *articleRepo) GetArticleInfo(ctx context.Context, id int64) (rv int64, err error) {
	m, err := u.data.rdb.HGetAll(ctx, articleKey(id)).Result()
	fmt.Println(m, err)
	return
}
