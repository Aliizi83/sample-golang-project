package dto

import "github.com/Aliizi83/sample-golang-project/src/services/dto"

type SendOtpRequest struct {
	MobileNumber string `json:"mobileNumber" binding:"required,mobile,min=11,max=11"`
}

type TokenDetail struct {
	AccessToken            string `json:"accessToken"`
	RefreshToken           string `json:"refreshToken"`
	AccessTokenExpireTime  int    `json:"accessTokenExpireTime"`
	RefreshTokenExpireTime int    `json:"accessTokenExpireTime"`
}

type RegisterUserByUsernameRequest struct {
	FirstName string `json:"firstName" binding:"required,min=3"`
	LastName  string `json:"lastName" binding:"required,min=3"`
	Username  string `json:"username" binding:"required,min=5"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,password,min=8"`
}

type RegisterLoginByMobileRequest struct {
	MobileNumber string `json:"mobileNumber" binding:"required,mobile"`
	Otp          string `json:"otp" binding:"required,min=6,max=6"`
}

type LoginByUsernameRequest struct {
	Username string `json:"username" binding:"required,min=5"`
	Password string `json:"password" binding:"required,min=6"`
}

func (from RegisterUserByUsernameRequest) ToRegisterUserByUsername() dto.RegisterUserByUsername {
	return dto.RegisterUserByUsername{
		FirstName: from.FirstName,
		LastName:  from.LastName,
		Username:  from.Username,
		Email:     from.Email,
		Password:  from.Password,
	}
}
