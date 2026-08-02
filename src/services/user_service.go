package services

import (
	"context"

	"github.com/Aliizi83/sample-golang-project/src/config"
	"github.com/Aliizi83/sample-golang-project/src/domain/repositories"
	"github.com/Aliizi83/sample-golang-project/src/pkg/logging"
	"github.com/Aliizi83/sample-golang-project/src/pkg/service_errors"
	"github.com/Aliizi83/sample-golang-project/src/services/dto"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	cfg        *config.Config
	logger     logging.Logger
	repository repositories.UserRepository
}

func (s *UserService) RegisterUserByUsername(ctx context.Context, req dto.RegisterUserByUsername) error {
	user := req.ToUserModel()

	exists, err := s.repository.ExistsEmail(ctx, user.Email)
	if err != nil {
		return err
	}
	if exists {
		return &service_errors.ServiceError{EndUserMessage: service_errors.EmailExists}
	}

	exists, err = s.repository.ExistsUsername(ctx, user.Username)
	if err != nil {
		return err
	}
	if exists {
		return &service_errors.ServiceError{EndUserMessage: service_errors.UsernameExists}
	}

	bcryptPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		s.logger.Error(err, logging.General, logging.HashPassword, err.Error(), nil)
		return err
	}
	user.Password = string(bcryptPassword)

	_, err = s.repository.CreateUser(ctx, user)
	if err != nil {
		s.logger.Error(err, logging.Postgres, logging.Insert, err.Error(), nil)
		return err
	}
	return nil
}

