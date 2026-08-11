package dto

type CreateCarModel struct {
	Name      string
	CompanyId int
	CarTypeId int
	GearboxId int
}

type UpdateCarModel struct {
	Name      string
	CompanyId int
	CarTypeId int
	GearboxId int
}

type CarModel struct {
	Id             int
	Name           string
	CarType        CarType
	Company        Company
	Gearbox        Gearbox
	CarModelColors []CarModelColor
}
