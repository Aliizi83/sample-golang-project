package services

import (
	"context"

	"github.com/Aliizi83/sample-golang-project/src/common"
	"github.com/Aliizi83/sample-golang-project/src/config"
	"github.com/Aliizi83/sample-golang-project/src/dependencies"
	"github.com/Aliizi83/sample-golang-project/src/domain/models"
	"github.com/Aliizi83/sample-golang-project/src/domain/repositories"
	"github.com/Aliizi83/sample-golang-project/src/pkg/logging"
	"github.com/Aliizi83/sample-golang-project/src/pkg/service_errors"
	"github.com/Aliizi83/sample-golang-project/src/services/dto"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	cfg          *config.Config
	logger       logging.Logger
	repository   repositories.UserRepository
	otpService   *OtpService
	tokenService *TokenService
}

func NewUserService(cfg *config.Config) *UserService {
	logger := logging.NewLogger(cfg)
	repository := dependencies.GetUserRepository(cfg)
	otpService := NewOtpService(cfg)
	tokenService := NewTokenService(cfg)

	return &UserService{
		cfg:          cfg,
		logger:       logger,
		repository:   repository,
		otpService:   otpService,
		tokenService: tokenService,
	}
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

func (s *UserService) RegisterAndLoginByMobile(ctx context.Context, mobileNumber string, otp string) (*dto.TokenDetail, error) {
	err := s.otpService.ValidateOtp(ctx, mobileNumber, otp)
	if err != nil {
		return nil, err
	}

	user := models.User{MobileNumber: mobileNumber, Username: mobileNumber}

	userExists, err := s.repository.ExistsMobileNumber(ctx, mobileNumber)
	if err != nil {
		return nil, err
	}

	if userExists {
		user, err = s.repository.GetUserByUsername(ctx, user.Username)
		if err != nil {
			return nil, err
		}

		return s.generateToken(user)
	}

	byteGeneratedPassword := []byte(common.GeneratePassword())
	hashedPassword, err := bcrypt.GenerateFromPassword(byteGeneratedPassword, bcrypt.DefaultCost)
	if err != nil {
		return nil, &service_errors.ServiceError{EndUserMessage: service_errors.HashFailed}
	}

	user.Password = string(hashedPassword)
	user, err = s.repository.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return s.generateToken(user)

}

func (s *UserService) LoginByUsername(ctx context.Context, username, password string) (*dto.TokenDetail, error) {
	user, err := s.repository.FetchUserInfo(ctx, username, password)
	if err != nil {
		return nil, &service_errors.ServiceError{EndUserMessage: service_errors.WrongPassword}
	}

	return s.generateToken(user)
}

func (s *UserService) generateToken(user models.User) (*dto.TokenDetail, error) {
	tokenDto := tokenDto{
		UserId:       int(user.Id),
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		Username:     user.Username,
		Email:        user.Email,
		MobileNumber: user.MobileNumber,
	}

	if len(*user.UserRoles) > 0 {
		for _, v := range *user.UserRoles {
			tokenDto.Roles = append(tokenDto.Roles, v.Role.Name)
		}
	}

	return s.tokenService.GenerateToken(&tokenDto)

}
