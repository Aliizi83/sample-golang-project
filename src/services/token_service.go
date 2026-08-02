package services

import (
	"time"

	"github.com/Aliizi83/sample-golang-project/src/api/dto"
	"github.com/Aliizi83/sample-golang-project/src/config"
	"github.com/Aliizi83/sample-golang-project/src/pkg/logging"
	"github.com/Aliizi83/sample-golang-project/src/pkg/service_errors"
	"github.com/golang-jwt/jwt/v5"
)

type TokenService struct {
	cfg    *config.Config
	logger *logging.Logger
}

type tokenDto struct {
	UserId       int
	FirstName    string
	LastName     string
	Username     string
	Email        string
	MobileNumber string
	Roles        []string
}

func NewTokenService(cfg *config.Config) *TokenService {
	logger := logging.NewLogger(cfg)

	return &TokenService{
		cfg:    cfg,
		logger: &logger,
	}
}

func (s *TokenService) GenerateToken(token *tokenDto) (*dto.TokenDetail, error) {
	tokenDetail := &dto.TokenDetail{}

	tokenDetail.AccessTokenExpireTime = int(time.Now().Add(s.cfg.JWT.AccessTokenExpireDuration * time.Minute).Unix())
	tokenDetail.RefreshTokenExpireTime = int(time.Now().Add(s.cfg.JWT.RefreshTokenExpireDuration * time.Minute).Unix())

	accessTokenClaims := jwt.MapClaims{}

	accessTokenClaims["user_id"] = token.UserId
	accessTokenClaims["first_name"] = token.FirstName
	accessTokenClaims["last_name"] = token.LastName
	accessTokenClaims["username"] = token.Username
	accessTokenClaims["email"] = token.Email
	accessTokenClaims["mobile_number"] = token.MobileNumber
	accessTokenClaims["roles"] = token.Roles
	accessTokenClaims["exp"] = tokenDetail.AccessTokenExpireTime

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)

	var err error
	tokenDetail.AccessToken, err = accessToken.SignedString([]byte(s.cfg.JWT.Secret))

	if err != nil {
		return nil, err
	}

	refreshTokenClaims := jwt.MapClaims{}
	refreshTokenClaims["user_id"] = token.UserId
	refreshTokenClaims["exp"] = tokenDetail.RefreshTokenExpireTime

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)

	tokenDetail.RefreshToken, err = refreshToken.SignedString([]byte(s.cfg.JWT.RefreshSecret))

	if err != nil {
		return nil, err
	}

	return tokenDetail, nil
}

func (s *TokenService) VerifyToken(token string) (*jwt.Token, error) {
	ac, err := jwt.Parse(token, func(t *jwt.Token) (any, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, &service_errors.ServiceError{EndUserMessage: service_errors.UnExpectedError}
		}

		return []byte(s.cfg.JWT.Secret), nil
	})

	if err != nil {
		return nil, err
	}
	return ac, nil
}

func (s *TokenService) GetClaims(token string) (claimMap map[string]any, err error) {
	claimMap = map[string]interface{}{}

	verifyToken, err := s.VerifyToken(token)
	if err != nil {
		return nil, err
	}

	claims, ok := verifyToken.Claims.(jwt.MapClaims)
	if ok && verifyToken.Valid {
		for k, v := range claims {
			claimMap[k] = v
		}

		return claimMap, nil
	}

	return nil, &service_errors.ServiceError{EndUserMessage: service_errors.ClaimsNotFound}
}
