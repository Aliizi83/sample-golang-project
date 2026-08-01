package services

import (
	"context"
	"fmt"

	"github.com/Aliizi83/sample-golang-project/src/common"
	"github.com/Aliizi83/sample-golang-project/src/config"
	"github.com/Aliizi83/sample-golang-project/src/constants"
	"github.com/Aliizi83/sample-golang-project/src/infra/cache"
	"github.com/Aliizi83/sample-golang-project/src/pkg/logging"
	"github.com/Aliizi83/sample-golang-project/src/pkg/service_errors"
	"github.com/redis/go-redis/v9"
)

type OtpService struct {
	logger logging.Logger
	cfg    *config.Config
	redis  *redis.Client
}

type otpDto struct {
	Value string
	Used  bool
}

func NewOtpService(cfg *config.Config) *OtpService {
	redis := cache.GetRedis()
	logger := logging.NewLogger(cfg)

	return &OtpService{
		logger: logger,
		cfg:    cfg,
		redis:  redis,
	}
}

func (s *OtpService) SendOtp(ctx context.Context, mobileNumber string) error {
	otp := common.GenerateOtp()
	err := s.SetOtp(ctx, mobileNumber, otp)
	if err != nil {
		return err
	}

	return nil
}

func (s *OtpService) SetOtp(ctx context.Context, mobileNumber string, otp string) error {
	key := fmt.Sprintf("%s:%s", constants.RedisOtpKey, mobileNumber)
	otpData := &otpDto{
		Value: otp,
		Used:  false,
	}

	value, err := cache.Get[otpDto](ctx, s.redis, key)

	if err == nil && !value.Used {
		return &service_errors.ServiceError{EndUserMessage: service_errors.OtpExists}
	} else if err == nil && value.Used {
		return &service_errors.ServiceError{EndUserMessage: service_errors.OtpUsed}
	}

	err = cache.Set(ctx, s.redis, key, otpData, s.cfg.Otp.ExpireTime)
	if err != nil {
		return err
	}
	return nil
}

func (s *OtpService) ValidateOtp(ctx context.Context, mobileNumber string, otp string) error {
	key := fmt.Sprintf("%s:%s", constants.RedisOtpKey, mobileNumber)

	res, err := cache.Get[otpDto](ctx, s.redis, key)
	if err != nil {
		return err
	}

	if res.Used {
		return &service_errors.ServiceError{EndUserMessage: service_errors.OtpUsed}
	}

	if res.Value != otp {
		return &service_errors.ServiceError{EndUserMessage: service_errors.OtpNotValid}
	} else {
		res.Used = true
		err = cache.Set(ctx, s.redis, key, res, s.cfg.Otp.ExpireTime)
	}
	return nil
}
