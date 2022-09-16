package biz

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

type User struct {
	ID        int64
	Username  string
	Phone     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserRepo interface {
	Register(ctx context.Context, user *User) (*User, error)
	Login(ctx context.Context, user *User) (*User, error)
	GetUserByPhone(ctx context.Context, phone string) (*User, error)
	UpdateUser(ctx context.Context, id int64, user *User) (*User, error)
}

type UserUsecase struct {
	repo UserRepo
}

func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo}
}

func (uc *UserUsecase) Register(ctx context.Context, phone string) (p *User, err error) {
	return uc.repo.GetUserByPhone(ctx, phone)
}

func (uc *UserUsecase) Login(ctx context.Context, user *User) (*User, error) {
	return uc.repo.Login(ctx, user)
}

func (uc *UserUsecase) GetUserByPhone(ctx context.Context, phone string) (*User, error) {
	return uc.repo.GetUserByPhone(ctx, phone)
}

func (uc *UserUsecase) UpdateUser(ctx context.Context, id int64, user *User) (*User, error) {
	return uc.repo.UpdateUser(ctx, id, user)
}
