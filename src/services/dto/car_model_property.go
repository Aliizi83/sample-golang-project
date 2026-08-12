package dto

type CreateCarModelProperty struct {
	CarModelId int
	PropertyId int
	Value      string
}

type UpdateCarModelProperty struct {
	Value string
}

type CarModelProperty struct {
	Id         int
	CarModelId int
	Property   Property
	Value      string
}
