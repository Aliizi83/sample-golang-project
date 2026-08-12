package dto

import "github.com/Aliizi83/sample-golang-project/src/services/dto"

type CreateCarModelPropertyRequest struct {
	CarModelId int    `json:"carModelId" binding:"required"`
	PropertyId int    `json:"propertyId" binding:"required"`
	Value      string `json:"value" binding:"required,max=100"`
}

type UpdateCarModelPropertyRequest struct {
	Value string `json:"value" binding:"required,max=100"`
}

type CarModelPropertyResponse struct {
	Id         int              `json:"id"`
	CarModelId int              `json:"carModelId,omitempty"`
	Property   PropertyResponse `json:"property,omitempty"`
	Value      string           `json:"value"`
}

func ToCarModelPropertyResponse(from dto.CarModelProperty) CarModelPropertyResponse {
	return CarModelPropertyResponse{
		Id:         from.Id,
		CarModelId: from.CarModelId,
		Property:   ToPropertyResponse(from.Property),
		Value:      from.Value,
	}
}

func ToCreateCarModelProperty(from CreateCarModelPropertyRequest) dto.CreateCarModelProperty {
	return dto.CreateCarModelProperty{
		CarModelId: from.CarModelId,
		PropertyId: from.PropertyId,
		Value:      from.Value,
	}
}

func ToUpdateCarModelProperty(from UpdateCarModelPropertyRequest) dto.UpdateCarModelProperty {
	return dto.UpdateCarModelProperty{
		Value: from.Value,
	}
}
