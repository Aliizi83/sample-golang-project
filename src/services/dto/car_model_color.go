package dto

type CreateCarModelColor struct {
	CarModelId int
	ColorId    int
}

type UpdateCarModelColor struct {
	CarModelId int
	ColorId    int
}

type CarModelColor struct {
	Id    int
	Color Color
}
