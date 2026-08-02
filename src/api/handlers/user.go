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
	cfg        *config.Config
	logger     logging.Logger
	otpService services.OtpService
}

func NewUserHandler(cfg *config.Config) *UserHandler {
	logger := logging.NewLogger(cfg)

	return &UserHandler{
		cfg:        cfg,
		logger:     logger,
		otpService: *services.NewOtpService(cfg),
	}
}

func (h *UserHandler) SendOtp(c *gin.Context) {
	req := new(dto.SendOtpRequest)
	err := c.ShouldBindJSON(req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helpers.GenerateBaseResponseWithValidationError(nil, false, -1, err))
		return
	}

	err = h.otpService.SendOtp(c.Request.Context(), req.MobileNumber)
	if err != nil {
		c.AbortWithStatusJSON(helpers.TranslateErrorToStatusCode(err), helpers.GenerateBaseResponseWithValidationError(nil, false, -1, err))
		return
	}

	c.JSON(http.StatusCreated, helpers.GenerateBaseResponse(nil, true, int(helpers.Success)))
}
