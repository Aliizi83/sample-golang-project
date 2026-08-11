package dto

type CreateCarModelYear struct {
	CarModelId    int
	PersianYearId int
}

type UpdateCarModelYear struct {
	CarModelId    int
	PersianYearId int
}

type CarModelYear struct {
	Id          int
	PersianYear PersianYear
	CarModelId  int
}
