package dto

import (
	"sync"

	"github.com/Aliizi83/sample-golang-project/src/services/dto"
)

type CreateCarModelRequest struct {
	Name      string `json:"name" binding:"required,min=3,max=15"`
	CompanyId int    `json:"companyId" binding:"required"`
	CarTypeId int    `json:"carTypeId" binding:"required"`
	GearboxId int    `json:"gearBoxId" binding:"required"`
}

type UpdateCarModelRequest struct {
	Name      string `json:"name,omitempty"`
	CompanyId int    `json:"companyId,omitempty"`
	CarTypeId int    `json:"carTypeId,omitempty"`
	GearboxId int    `json:"gearboxId,omitempty"`
}
type CarModelResponse struct {
	Id      int                     `json:"id"`
	Name    string                  `json:"name"`
	Company CompanyResponse         `json:"company"`
	CarType CarTypeResponse         `json:"carType"`
	Gearbox GearboxResponse         `json:"gearbox"`
	Colors  []CarModelColorResponse `json:"colors"`
}

func ToCarModelResponse(from dto.CarModel) CarModelResponse {
	colors := []CarModelColorResponse{}

	var wg sync.WaitGroup

	go func() {
		for _, v := range from.CarModelColors {
			colors = append(colors, ToCarModelColorResponse(v))
		}
		wg.Done()
	}()

	wg.Wait()

	return CarModelResponse{
		Id:      from.Id,
		Name:    from.Name,
		Company: ToCompanyResponse(from.Company),
		CarType: ToCarTypeResponse(from.CarType),
		Gearbox: ToGearboxResponse(from.Gearbox),
	}
}

func ToCreateCarModel(from CreateCarModelRequest) dto.CreateCarModel {
	return dto.CreateCarModel{
		Name:      from.Name,
		CompanyId: from.CompanyId,
		CarTypeId: from.CarTypeId,
		GearboxId: from.GearboxId,
	}
}

func ToUpdateCarModel(from UpdateCarModelRequest) dto.UpdateCarModel {
	return dto.UpdateCarModel{
		Name:      from.Name,
		CompanyId: from.CompanyId,
		CarTypeId: from.CarTypeId,
		GearboxId: from.GearboxId,
	}
}
