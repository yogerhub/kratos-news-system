package data

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"kratos-news-system/app/user/service/internal/biz"
)

type User struct {
	gorm.Model
	Username string `json:"username"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type userRepo struct {
	data *Data
	log  *log.Helper
}

// NewUserRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *userRepo) Register(ctx context.Context, user *biz.User) (*biz.User, error) {
	u := new(User)
	err := r.data.db.Where("username = ?", user.Username).First(u).Error
	if err != nil {
		return nil, err
	}
	if u != nil {
		return nil, errors.New("用户已存在")
	}
	u.Username = user.Username
	u.Phone = user.Phone
	u.Password = user.Password

	result := r.data.db.Create(&u) // 通过数据的指针来创建
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (r *userRepo) Login(ctx context.Context, user *biz.User) (*biz.User, error) {
	u := new(User)
	err := r.data.db.Where("username = ? AND password = ?", user.Username, user.Password).First(u).Error
	if err != nil {
		return nil, err
	}

	_, _ = r.SetUserInfo(ctx, u)

	return user, nil
}

func (r *userRepo) GetUserByPhone(ctx context.Context, phone string) (*biz.User, error) {
	u := new(User)
	err := r.data.db.Where("phone = ?", phone).First(u).Error
	if err != nil {
		return nil, err
	}
	return &biz.User{
		ID:        int64(u.ID),
		Username:  u.Username,
		Phone:     u.Phone,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}, nil
}

func (r *userRepo) UpdateUser(ctx context.Context, id int64, user *biz.User) (*biz.User, error) {
	var u User
	user.ID = id
	err := r.data.db.Model(&u).Updates(user).Error
	return &biz.User{
		ID:        int64(u.ID),
		Username:  u.Username,
		Phone:     u.Phone,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}, err
}
