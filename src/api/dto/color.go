package dto

import "github.com/Aliizi83/sample-golang-project/src/services/dto"

type CreateColorRequest struct {
	Name    string `json:"name" binding:"required,alpha,min=3,max=15"`
	HexCode string `json:"hexCode" binding:"min=7,max=7"`
}

type UpdateColorRequest struct {
	Name    string `json:"name" binding:"required,alpha,min=3,max=15"`
	HexCode string `json:"hexCode" binding:"min=7,max=7"`
}
type ColorResponse struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	HexCode string `json:"hexCode"`
}

func ToColorResponse(from dto.Color) ColorResponse {
	return ColorResponse{
		Id:      from.Id,
		Name:    from.Name,
		HexCode: from.HexCode,
	}
}

func ToCreateColor(from CreateColorRequest) dto.CreateColor {
	return dto.CreateColor{
		Name:    from.Name,
		HexCode: from.HexCode,
	}
}

func ToUpdateColor(from UpdateColorRequest) dto.UpdateColor {
	return dto.UpdateColor{
		Name:    from.Name,
		HexCode: from.HexCode,
	}
}
