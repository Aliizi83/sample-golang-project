package dto

type SendOtpRequest struct {
	MobileNumber string `json:"mobileNumber" binding:"required,mobile,min=11,max=11"`
}

type TokenDetail struct {
	AccessToken            string `json:"accessToken"`
	RefreshToken           string `json:"refreshToken"`
	AccessTokenExpireTime  int    `json:"accessTokenExpireTime"`
	RefreshTokenExpireTime int    `json:"accessTokenExpireTime"`
}
