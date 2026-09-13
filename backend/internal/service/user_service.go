package service

import (
	"context"

	"github.com/sachingunawardhana10/armalora-cloud-native/internal/model"
)

type UserService interface {
	Register(ctx context.Context, user *model.User) error
	FindByEmail(ctx context.Context, email string) (*model.User, error)
	FindByID(ctx context.Context, id int64) (*model.User, error)
}
