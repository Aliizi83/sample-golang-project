package dto

import (
	"time"

	"github.com/Aliizi83/sample-golang-project/src/services/dto"
)

type CreateCarModelPriceHistoryRequest struct {
	CarModelYearId int       `json:"carModelYearId" binding:"required"`
	PriceAt        time.Time `json:"priceAt" binding:"required"`
	Price          float64   `json:"price" binding:"required"`
}

type UpdateCarModelPriceHistoryRequest struct {
	PriceAt time.Time `json:"priceAt,omitempty"`
	Price   float64   `json:"price,omitempty"`
}

type CarModelPriceHistoryResponse struct {
	Id             int       `json:"id"`
	CarModelYearId int       `json:"carModelYearId"`
	PriceAt        time.Time `json:"priceAt,omitempty"`
	Price          float64   `json:"price,omitempty"`
}

func ToCarModelPriceHistoryResponse(from dto.CarModelPriceHistory) CarModelPriceHistoryResponse {
	return CarModelPriceHistoryResponse{
		Id:             from.Id,
		CarModelYearId: from.CarModelYearId,
		PriceAt:        from.PriceAt,
		Price:          from.Price,
	}
}

func ToCreateCarModelPriceHistory(from CreateCarModelPriceHistoryRequest) dto.CreateCarModelPriceHistory {
	return dto.CreateCarModelPriceHistory{
		CarModelYearId: from.CarModelYearId,
		PriceAt:        from.PriceAt,
		Price:          from.Price,
	}
}

func ToUpdateCarModelPriceHistory(from UpdateCarModelPriceHistoryRequest) dto.UpdateCarModelPriceHistory {
	return dto.UpdateCarModelPriceHistory{
		PriceAt: from.PriceAt,
		Price:   from.Price,
	}
}
