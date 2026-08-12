package dto

type CreateCarModelImage struct {
	CarModelId  int
	ImageId     int
	IsMainImage bool
}

type UpdateCarModelImage struct {
	IsMainImage bool
}

type CarModelImage struct {
	Id          int
	CarModelId  int
	Image       File
	IsMainImage bool
}
