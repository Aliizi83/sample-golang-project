package dto

import "github.com/Aliizi83/sample-golang-project/src/services/dto"

type CreatePropertyCategoryRequest struct {
	Name string `json:"name" binding:"required,alpha,min=3,max=50"`
	Icon string `json:"icon" biding:"required,max=1000"`
}

type UpdatePropertyCategoryRequest struct {
	Name string `json:"name,omitempty"`
	Icon string `json:"icon,omitempty"`
}
type PropertyCategoryResponse struct {
	Id         int                `json:"id"`
	Name       string             `json:"name"`
	Icon       string             `json:"icon"`
	Properties []PropertyResponse `json:"properties,omitempty"`
}

func ToPropertyCategoryResponse(from dto.PropertyCategory) PropertyCategoryResponse {

	properties := []PropertyResponse{}
	for _, property := range from.Properties {
		properties = append(properties, ToPropertyResponse(property))
	}

	return PropertyCategoryResponse{
		Id:         from.Id,
		Name:       from.Name,
		Icon:       from.Icon,
		Properties: properties,
	}
}

func ToCreatePropertyCategory(from CreatePropertyCategoryRequest) dto.CreatePropertyCategory {
	return dto.CreatePropertyCategory{
		Name: from.Icon,
		Icon: from.Icon,
	}
}

func ToUpdatePropertyCategory(from UpdatePropertyCategoryRequest) dto.UpdatePropertyCategory {
	return dto.UpdatePropertyCategory{
		Name: from.Icon,
		Icon: from.Icon,
	}
}
