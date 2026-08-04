package handlers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/Aliizi83/sample-golang-project/src/api/helpers"
	"github.com/Aliizi83/sample-golang-project/src/config"
	"github.com/Aliizi83/sample-golang-project/src/domain/filters"
	"github.com/Aliizi83/sample-golang-project/src/pkg/logging"
	"github.com/gin-gonic/gin"
)

var logger = logging.NewLogger(config.GetConfig())

func Create[TRequest, TServiceInput, TServiceOutput, TResponse any](
	c *gin.Context,
	requestMapper func(req TRequest) TServiceInput,
	responseMapper func(serviceOutput TServiceOutput) TResponse,
	serviceCreateFunction func(ctx context.Context, serviceInput TServiceInput) (TServiceOutput, error),

) {
	request := new(TRequest)
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helpers.GenerateBaseResponseWithValidationError(nil, false, int(helpers.ValidationError), err))
		return
	}

	serviceInput := requestMapper(*request)
	response, err := serviceCreateFunction(c.Request.Context(), serviceInput)
	if err != nil {
		c.AbortWithStatusJSON(helpers.TranslateErrorToStatusCode(err),
			helpers.GenerateBaseResponseWithError(nil, false, int(helpers.InternalError), err))
		return
	}

	c.JSON(http.StatusCreated, helpers.GenerateBaseResponse(responseMapper(response), true, int(helpers.Success)))
}

func Update[TRequest, TServiceInput, TServiceOutput, TResponse any](
	c *gin.Context,
	requestMapper func(req TRequest) TServiceInput,
	responseMapper func(serviceOutput TServiceOutput) TResponse,
	serviceUpdateFunction func(ctx context.Context, id int, serviceInput TServiceInput) (TServiceOutput, error),
) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	request := new(TRequest)
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helpers.GenerateBaseResponseWithValidationError(nil, false, int(helpers.ValidationError), err))
		return

	}

	serviceOutput, err := serviceUpdateFunction(c.Request.Context(), id, requestMapper(*request))
	if err != nil {
		c.AbortWithStatusJSON(helpers.TranslateErrorToStatusCode(err),
			helpers.GenerateBaseResponseWithError(nil, false, int(helpers.InternalError), err))
		return
	}

	c.JSON(http.StatusOK, helpers.GenerateBaseResponse(responseMapper(serviceOutput), true, int(helpers.Success)))

}

func Delete(c *gin.Context, serviceDeleteFunc func(ctx context.Context, id int) error) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	if id == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound,
			helpers.GenerateBaseResponse(nil, false, int(helpers.ValidationError)))
		return
	}

	err := serviceDeleteFunc(c.Request.Context(), id)
	if err != nil {
		c.AbortWithStatusJSON(helpers.TranslateErrorToStatusCode(err),
			helpers.GenerateBaseResponseWithError(nil, false, int(helpers.InternalError), err))
		return
	}

	c.JSON(http.StatusOK, helpers.GenerateBaseResponse(nil, true, int(helpers.Success)))
}

func GetById[TServiceOutput, TResponse any](
	c *gin.Context,
	responseMapper func(output TServiceOutput) TResponse,
	serviceGetById func(ctx context.Context, id int) (TServiceOutput, error),
) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	if id == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound,
			helpers.GenerateBaseResponse(nil, false, int(helpers.ValidationError)))
		return
	}

	serviceOutput, err := serviceGetById(c.Request.Context(), id)
	if err != nil {
		c.AbortWithStatusJSON(helpers.TranslateErrorToStatusCode(err),
			helpers.GenerateBaseResponseWithError(nil, false, int(helpers.InternalError), err))
		return
	}

	c.JSON(http.StatusOK, helpers.GenerateBaseResponse(responseMapper(serviceOutput), true, int(helpers.Success)))
}

func GetByFilter[TServiceOutput, TResponse any](
	c *gin.Context,
	responseMapper func(output TServiceOutput) TResponse,
	serviceGetByFilterFunc func(ctx context.Context, req filters.PaginationInputWithFilter) (*filters.PagedList[TServiceOutput], error),
) {
	req := new(filters.PaginationInputWithFilter)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helpers.GenerateBaseResponseWithValidationError(nil, false, int(helpers.ValidationError), err))
		return
	}

	serviceOutput, err := serviceGetByFilterFunc(c.Request.Context(), *req)
	if err != nil {
		c.AbortWithStatusJSON(helpers.TranslateErrorToStatusCode(err),
			helpers.GenerateBaseResponseWithError(nil, false, int(helpers.InternalError), err))
		return
	}

	response := filters.PagedList[TResponse]{
		PageNumber:      serviceOutput.PageNumber,
		PageSize:        serviceOutput.PageSize,
		TotalRows:       serviceOutput.TotalRows,
		TotalPages:      serviceOutput.TotalPages,
		HasPreviousPage: serviceOutput.HasPreviousPage,
		HasNextPage:     serviceOutput.HasNextPage,
	}

	items := []TResponse{}
	for _, item := range *serviceOutput.Items {

		items = append(items, responseMapper(item))
	}
	response.Items = &items

	c.JSON(http.StatusOK, helpers.GenerateBaseResponse(response, true, 0))
}
