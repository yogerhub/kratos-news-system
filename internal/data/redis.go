package data

import (
	"context"
	"fmt"
)

func userKey(id int64) string {
	return fmt.Sprintf("user:id:%d", id)
}

func (u *userRepo) SetUserInfo(ctx context.Context, user *User) (res bool, err error) {
	args := make([]interface{}, 0)
	args = append(args, "id")
	args = append(args, user.ID)
	args = append(args, "username")
	args = append(args, user.Username)
	args = append(args, "phone")
	args = append(args, user.Phone)
	args = append(args, "updated_at")
	args = append(args, user.UpdatedAt)

	res, err = u.data.rdb.HMSet(ctx, userKey(int64(user.ID)), args...).Result()
	fmt.Println(res, err)
	return
}

func (u *userRepo) ExistsUserInfo(ctx context.Context, id int64) (rv bool, err error) {
	m, err := u.data.rdb.Exists(ctx, userKey(id)).Result()
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

func (u *userRepo) GetUserInfo(ctx context.Context, id int64) (rv int64, err error) {
	m, err := u.data.rdb.HGetAll(ctx, userKey(id)).Result()
	fmt.Println(m, err)
	return
}
