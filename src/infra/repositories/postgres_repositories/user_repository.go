package postgres_repositories

import (
	"context"

	"github.com/Aliizi83/sample-golang-project/src/config"
	"github.com/Aliizi83/sample-golang-project/src/constants"
	"github.com/Aliizi83/sample-golang-project/src/domain/models"
	"github.com/Aliizi83/sample-golang-project/src/infra/presistence/db"
	"github.com/Aliizi83/sample-golang-project/src/pkg/logging"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const userFilterExp string = "username = ?"
const countFilterExp string = "count(*) > 0"
const emailFilterExp string = "email = ?"

type PostgresUserRepository struct {
	database *gorm.DB
	logger   logging.Logger
}

func NewPostgresUserRepository(cfg *config.Config) *PostgresUserRepository {
	logger := logging.NewLogger(cfg)
	db := db.GetDB()

	return &PostgresUserRepository{
		database: db,
		logger:   logger,
	}
}

func (r *PostgresUserRepository) ExistsMobileNumber(ctx context.Context, mobileNumber string) (bool, error) {
	var exists bool
	if err := r.database.WithContext(ctx).
		Model(models.User{}).
		Select(countFilterExp).
		Where("mobile_number = ?", mobileNumber).
		Find(&exists).Error; err != nil {
		r.logger.Error(err, logging.Postgres, logging.Select, err.Error(), nil)
		return false, err
	}
	return exists, nil
}
func (r *PostgresUserRepository) ExistsUsername(ctx context.Context, username string) (bool, error) {
	var exists bool
	if err := r.database.WithContext(ctx).
		Model(&models.User{}).
		Select(countFilterExp).
		Where(userFilterExp, username).
		Find(&exists).Error; err != nil {
		r.logger.Error(err, logging.Postgres, logging.Select, err.Error(), nil)
		return false, err
	}
	return exists, nil

}
func (r *PostgresUserRepository) ExistsEmail(ctx context.Context, email string) (bool, error) {
	var exists bool
	if err := r.database.WithContext(ctx).
		Model(&models.User{}).
		Select(countFilterExp).
		Where(emailFilterExp, email).
		Find(&exists).Error; err != nil {
		r.logger.Error(err, logging.Postgres, logging.Select, err.Error(), nil)
		return false, err
	}
	return exists, nil
}
func (r *PostgresUserRepository) FetchUserInfo(ctx context.Context, username string, password string) (models.User, error) {
	var user models.User
	if err := r.database.WithContext(ctx).
		Model(&models.User{}).
		Where(userFilterExp, username).
		Preload("UserRoles", func(tx *gorm.DB) *gorm.DB {
			return tx.Preload("Role")
		}).
		Find(&user).Error; err != nil {
		r.logger.Error(err, logging.Postgres, logging.Select, err.Error(), nil)
		return user, err
	}

	return user, bcrypt.CompareHashAndPassword([]byte(password), []byte(user.Password))

}
func (r *PostgresUserRepository) GetDefaultRole(ctx context.Context) (roleId int, err error) {
	if err = r.database.WithContext(ctx).Model(&models.Role{}).
		Select("id").
		Where("name = ?", constants.DefaultCustomerRoleTitle).
		First(&roleId).Error; err != nil {
		return 0, err
	}
	return roleId, nil
}
func (r *PostgresUserRepository) CreateUser(ctx context.Context, u models.User) (models.User, error) {
	var user models.User
	roleId, err := r.GetDefaultRole(ctx)
	if err != nil {
		r.logger.Error(err, logging.Postgres, logging.Rollback, err.Error(), nil)
		return user, err
	}

	tx := r.database.WithContext(ctx).Begin()
	err = tx.Create(&u).Error
	if err != nil {
		tx.Rollback()
		r.logger.Error(err, logging.Postgres, logging.Rollback, err.Error(), nil)
		return user, err
	}

	userRoleModel := &models.UserRole{
		RoleId: uint(roleId),
		UserId: u.Id,
	}
	err = tx.Create(&userRoleModel).Error
	if err != nil {
		tx.Rollback()
		r.logger.Error(err, logging.Postgres, logging.Rollback, err.Error(), nil)
		return user, err
	}
	tx.Commit()
	return user, nil

}
