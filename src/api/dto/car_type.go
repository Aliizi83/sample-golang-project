package dto

import "github.com/Aliizi83/sample-golang-project/src/services/dto"

type CreateCarTypeRequest struct {
	Name string `json:"name" binding:"required,min=3,max=15"`
}

type UpdateCarTypeRequest struct {
	Name string `json:"name" binding:"required,min=3,max=15"`
}
type CarTypeResponse struct {
	Id   int    `json:"id"`
	Name string `json:"string"`
}

func ToCarTypeResponse(from dto.CarType) CarTypeResponse {
	return CarTypeResponse{
		Id:   from.Id,
		Name: from.Name,
	}
}

func ToCreateCarType(from CreateCarTypeRequest) dto.CreateCarType {
	return dto.CreateCarType{
		Name: from.Name,
	}
}

func ToUpdateCarType(from UpdateCarTypeRequest) dto.UpdateCarType {
	return dto.UpdateCarType{
		Name: from.Name,
	}
}
