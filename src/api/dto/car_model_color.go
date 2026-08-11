package dto

import "github.com/Aliizi83/sample-golang-project/src/services/dto"

type CreateCarModelColorRequest struct {
	CarModelId int `json:"carModelId" binding:"required"`
	ColorId    int `json:"colorId" binding:"required"`
}

type UpdateCarModelColorRequest struct {
	CarModelId int `json:"carModelId,omitempty"`
	ColorId    int `json:"colorId,omitempty"`
}

type CarModelColorResponse struct {
	Id    int           `json:"id"`
	Color ColorResponse `json:"color,omitempty"`
}

func ToCarModelColorResponse(from dto.CarModelColor) CarModelColorResponse {
	return CarModelColorResponse{
		Id:    from.Id,
		Color: ToColorResponse(from.Color),
	}
}

func ToCreateCarModelColor(from CreateCarModelColorRequest) dto.CreateCarModelColor {
	return dto.CreateCarModelColor{
		CarModelId: from.CarModelId,
		ColorId:    from.ColorId,
	}
}

func ToUpdateCarModelColor(from UpdateCarModelColorRequest) dto.UpdateCarModelColor {
	return dto.UpdateCarModelColor{
		CarModelId: from.CarModelId,
		ColorId:    from.ColorId,
	}
}
