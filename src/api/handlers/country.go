package handlers

import (
	"net/http"
	"strconv"

	"github.com/Aliizi83/sample-golang-project/src/api/dto"
	"github.com/Aliizi83/sample-golang-project/src/api/helpers"
	"github.com/Aliizi83/sample-golang-project/src/config"
	"github.com/Aliizi83/sample-golang-project/src/pkg/logging"
	"github.com/Aliizi83/sample-golang-project/src/services"
	"github.com/gin-gonic/gin"
)

type CountryHandler struct {
	cfg            *config.Config
	logger         logging.Logger
	countryService services.CountryService
}

func NewCountryHandler(cfg *config.Config) *CountryHandler {
	logger := logging.NewLogger(cfg)
	countryService := services.NewCountryService(cfg)

	return &CountryHandler{
		cfg:            cfg,
		logger:         logger,
		countryService: *countryService,
	}
}

// CreateCountry godoc
// @Summary Create a country
// @Description Create a country
// @Tags Countries
// @Accept json
// @produces json
// @Param Request body dto.CreateAndUpdateCountryRequest true "Create a country"
// @Success 201 {object} helpers.BaseHttpResponse{result=dto.CountryResponse} "Country response"
// @Failure 400 {object} helpers.BaseHttpResponse "Bad request"
// @Router /v1/countries/ [post]
// @Security AuthBearer
func (h *CountryHandler) Create(c *gin.Context) {
	req := dto.CreateAndUpdateCountryRequest{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helpers.GenerateBaseResponseWithValidationError(nil, false, int(helpers.ValidationError), err))
		return
	}

	result, err := h.countryService.Create(c.Request.Context(), req)
	if err != nil {
		c.AbortWithStatusJSON(helpers.TranslateErrorToStatusCode(err),
			helpers.GenerateBaseResponseWithError(nil, false, int(helpers.InternalError), err))
		return
	}

	c.JSON(http.StatusCreated, helpers.GenerateBaseResponse(result, true, int(helpers.Success)))
}

// UpdateCountry godoc
// @Summary Update a country
// @Description Update a country
// @Tags Countries
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Param Request body dto.CreateAndUpdateCountryRequest true "Update a country"
// @Success 201 {object} helpers.BaseHttpResponse{result=dto.CountryResponse} "Country response"
// @Failure 400 {object} helpers.BaseHttpResponse "Bad request"
// @Router /v1/countries/{id} [put]
// @Security AuthBearer
func (h *CountryHandler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	req := dto.CreateAndUpdateCountryRequest{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helpers.GenerateBaseResponseWithValidationError(nil, false, int(helpers.ValidationError), err))
		return
	}

	result, err := h.countryService.Update(c.Request.Context(), id, req)
	if err != nil {
		c.AbortWithStatusJSON(helpers.TranslateErrorToStatusCode(err),
			helpers.GenerateBaseResponseWithError(nil, false, int(helpers.InternalError), err))
		return
	}

	c.JSON(http.StatusCreated, helpers.GenerateBaseResponse(result, true, int(helpers.Success)))
}

// DeleteCountry godoc
// @Summary Delete a country
// @Description Delete a country
// @Tags Countries
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helpers.BaseHttpResponse "response"
// @Failure 400 {object} helpers.BaseHttpResponse "Bad request"
// @Router /v1/countries/{id} [delete]
// @Security AuthBearer
func (h *CountryHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	if id == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound,
			helpers.GenerateBaseResponse(nil, false, 121))
		return
	}
	err := h.countryService.Delete(c.Request.Context(), id)
	if err != nil {
		c.AbortWithStatusJSON(helpers.TranslateErrorToStatusCode(err),
			helpers.GenerateBaseResponseWithError(nil, false, int(helpers.InternalError), err))
		return
	}

	c.JSON(http.StatusOK, helpers.GenerateBaseResponse(nil, true, 0))
}

// GetCountry godoc
// @Summary Get a country
// @Description Get a country
// @Tags Countries
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helpers.BaseHttpResponse{result=dto.CountryResponse} "Country response"
// @Failure 400 {object} helpers.BaseHttpResponse "Bad request"
// @Router /v1/countries/{id} [get]
// @Security AuthBearer
func (h *CountryHandler) GetById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	if id == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound,
			helpers.GenerateBaseResponse(nil, false, 121))
		return
	}

	result, err := h.countryService.GetById(c.Request.Context(), id)
	if err != nil {
		c.AbortWithStatusJSON(helpers.TranslateErrorToStatusCode(err),
			helpers.GenerateBaseResponseWithError(nil, false, 121, err))
		return
	}

	c.JSON(http.StatusOK, helpers.GenerateBaseResponse(result, true, 0))
}
