package dto

type CreateCompany struct {
	Name      string
	CountryId int
}

type UpdateCompany struct {
	Name      string
	CountryId int
}
type Company struct {
	Id      int
	Name    string
	Country Country
}
