package handlers

import (
	"net/http"

	"github.com/Aliizi83/sample-golang-project/src/api/dto"
	"github.com/Aliizi83/sample-golang-project/src/api/helpers"
	"github.com/Aliizi83/sample-golang-project/src/config"
	"github.com/Aliizi83/sample-golang-project/src/pkg/logging"
	"github.com/Aliizi83/sample-golang-project/src/services"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	cfg         *config.Config
	logger      logging.Logger
	otpService  services.OtpService
	userService services.UserService
}

func NewUserHandler(cfg *config.Config) *UserHandler {
	logger := logging.NewLogger(cfg)

	return &UserHandler{
		cfg:         cfg,
		logger:      logger,
		otpService:  *services.NewOtpService(cfg),
		userService: *services.NewUserService(cfg),
	}
}

// SendOtp godoc
// @Summary Send otp to user
// @Description Send otp to user
// @Tags Users
// @Accept  json
// @Produce  json
// @Param Request body dto.SendOtpRequest true "SendOtpRequest"
// @Success 201 {object} helpers.BaseHttpResponse "Success"
// @Failure 400 {object} helpers.BaseHttpResponse "Failed"
// @Failure 409 {object} helpers.BaseHttpResponse "Failed"
// @Router /v1/users/send-otp [post]
func (h *UserHandler) SendOtp(c *gin.Context) {
	req := new(dto.SendOtpRequest)
	err := c.ShouldBindJSON(req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helpers.GenerateBaseResponseWithValidationError(nil, false, int(helpers.ValidationError), err))
		return
	}

	err = h.otpService.SendOtp(c.Request.Context(), req.MobileNumber)
	if err != nil {
		c.AbortWithStatusJSON(helpers.TranslateErrorToStatusCode(err), helpers.GenerateBaseResponseWithValidationError(nil, false, -1, err))
		return
	}

	c.JSON(http.StatusCreated, helpers.GenerateBaseResponse(nil, true, int(helpers.Success)))
}

// RegisterUserByUsername godoc
// @Summary Register user by username and password
// @Description Register user by username and password
// @Tags Users
// @Accept  json
// @Produce  json
// @Param Request body dto.RegisterUserByUsernameRequest true "RegisterUserByUsernameRequest"
// @Success 201 {object} helpers.BaseHttpResponse "Success"
// @Failure 400 {object} helpers.BaseHttpResponse "Failed"
// @Failure 409 {object} helpers.BaseHttpResponse "Failed"
// @Router /v1/users/register-by-username [post]
func (h *UserHandler) RegisterUserByUsername(c *gin.Context) {
	req := new(dto.RegisterUserByUsernameRequest)
	err := c.ShouldBindJSON(req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helpers.GenerateBaseResponseWithValidationError(nil, false, int(helpers.ValidationError), err))
		return
	}

	err = h.userService.RegisterUserByUsername(c.Request.Context(), req.ToRegisterUserByUsername())
	if err != nil {
		c.AbortWithStatusJSON(helpers.TranslateErrorToStatusCode(err),
			helpers.GenerateBaseResponseWithError(err, false, int(helpers.InternalError), err))
		return
	}
	c.JSON(http.StatusCreated, helpers.GenerateBaseResponse(nil, true, int(helpers.Success)))
}
