package repositories

import (
	"context"

	"github.com/Aliizi83/sample-golang-project/src/domain/models"
)

type UserRepository interface {
	ExistsMobileNumber(ctx context.Context, mobileNumber string) (bool, error)
	ExistsUsername(ctx context.Context, username string) (bool, error)
	ExistsEmail(ctx context.Context, email string) (bool, error)
	FetchUserInfo(ctx context.Context, username string, password string) (models.User, error)
	GetUserByUsername(ctx context.Context, username string) (models.User, error)
	GetDefaultRole(ctx context.Context) (roleId int, err error)
	CreateUser(ctx context.Context, u models.User) (models.User, error)
}
