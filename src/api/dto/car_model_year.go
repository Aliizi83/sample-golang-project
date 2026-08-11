package dto

import "github.com/Aliizi83/sample-golang-project/src/services/dto"

type CreateCarModelYearRequest struct {
	CarModelId    int `json:"carModelId" binding:"required"`
	PersianYearId int `json:"persianYearId" binding:"required"`
}

type UpdateCarModelYearRequest struct {
	CarModelId    int `json:"carModelId,omitempty"`
	PersianYearId int `json:"persianYearId,omitempty"`
}
type CarModelYearResponse struct {
	Id          int                 `json:"id"`
	PersianYear PersianYearResponse `json:"persianYear,omitempty"`
	CarModelId  int                 `json:"carModelId,omitempty"`
}

func ToCarModelYearResponse(from dto.CarModelYear) CarModelYearResponse {
	return CarModelYearResponse{
		Id:          from.Id,
		CarModelId:  from.CarModelId,
		PersianYear: ToPersianYearResponse(from.PersianYear),
	}
}

func ToCreateCarModelYear(from CreateCarModelYearRequest) dto.CreateCarModelYear {
	return dto.CreateCarModelYear{
		CarModelId:    from.CarModelId,
		PersianYearId: from.PersianYearId,
	}
}

func ToUpdateCarModelYear(from UpdateCarModelYearRequest) dto.UpdateCarModelYear {
	return dto.UpdateCarModelYear{
		CarModelId:    from.CarModelId,
		PersianYearId: from.PersianYearId,
	}
}
