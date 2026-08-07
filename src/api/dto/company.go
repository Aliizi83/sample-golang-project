package dto

import "github.com/Aliizi83/sample-golang-project/src/services/dto"

type CreateCompanyRequest struct {
	CountryId int    `json:"country" binding:"required"`
	Name      string `json:"name" binding:"required,min=3,max=15"`
}

type UpdateCompanyRequest struct {
	CountryId int    `json:"country" binding:"required"`
	Name      string `json:"name" binding:"required,min=3,max=15"`
}
type CompanyResponse struct {
	Id      int             `json:"id"`
	Name    string          `json:"name"`
	Country CountryResponse `json:"country"`
}

func ToCompanyResponse(from dto.Company) CompanyResponse {
	return CompanyResponse{
		Id:      from.Id,
		Name:    from.Name,
		Country: ToCountryResponse(from.Country),
	}
}

func ToCreateCompany(from CreateCompanyRequest) dto.CreateCompany {
	return dto.CreateCompany{
		CountryId: from.CountryId,
		Name:      from.Name,
	}
}

func ToUpdateCompany(from UpdateCompanyRequest) dto.UpdateCompany {
	return dto.UpdateCompany{
		CountryId: from.CountryId,
		Name:      from.Name,
	}
}
