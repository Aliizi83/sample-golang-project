package dto

import "time"

type CreateCarModelPriceHistory struct {
	CarModelYearId int
	PriceAt        time.Time
	Price          float64
}

type UpdateCarModelPriceHistory struct {
	PriceAt time.Time
	Price   float64
}

type CarModelPriceHistory struct {
	Id             int
	CarModelYearId int
	PriceAt        time.Time
	Price          float64
}
