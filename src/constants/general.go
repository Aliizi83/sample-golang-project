package constants

const (
	RedisOtpKey string = "otp"

	// Claims

	AuthorizationHeaderKey string = "Authorization"
	UserIdKey              string = "UserId"
	FirstNameKey           string = "FirstName"
	LastNameKey            string = "LastName"
	UsernameKey            string = "Username"
	EmailKey               string = "Email"
	MobileNumberKey        string = "MobileNumber"
	RolesKey               string = "Roles"
	ExpireTimeKey          string = "Exp"
	ClaimsKey              string = "UserClaims"

	// JWT
	RefreshTokenCookieName string = "refresh_token"
)
