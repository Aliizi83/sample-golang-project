package dto

import "github.com/Aliizi83/sample-golang-project/src/services/dto"

type CreateGearboxRequest struct {
	Name string `json:"name" binding:"required,min=3,max=30"`
}

type UpdateGearboxRequest struct {
	Name string `json:"name" binding:"required,min=3,max=30"`
}
type GearboxResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func ToGearboxResponse(from dto.Gearbox) GearboxResponse {
	return GearboxResponse{
		Id:   from.Id,
		Name: from.Name,
	}
}

func ToCreateGearbox(from CreateGearboxRequest) dto.CreateGearbox {
	return dto.CreateGearbox{
		Name: from.Name,
	}
}

func ToUpdateGearbox(from UpdateGearboxRequest) dto.UpdateGearbox {
	return dto.UpdateGearbox{
		Name: from.Name,
	}
}
