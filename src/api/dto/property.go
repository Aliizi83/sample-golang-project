package dto

import "github.com/Aliizi83/sample-golang-project/src/services/dto"

type CreatePropertyRequest struct {
	Name        string `json:"name" binding:"required,alpha,min=3,max=50"`
	CategoryId  int    `json:"categoryId" binding:"required"`
	Icon        string `json:"icon" binding:"max=1000"`
	Description string `json:"description" binding:"max=1000"`
	DataType    string `json:"dataType" binding:"max=15"`
	Unit        string `json:"unit" binding:"max=15"`
}

type UpdatePropertyRequest struct {
	Name        string `json:"name,omitempty"`
	CategoryId  int    `json:"categoryId,omitempty"`
	Icon        string `json:"icon,omitempty" binding:"max=1000"`
	Description string `json:"description,omitempty" binding:"max=1000"`
	DataType    string `json:"dataType,omitempty" binding:"max=15"`
	Unit        string `json:"unit,omitempty" binding:"max=15"`
}
type PropertyResponse struct {
	Id          int                      `json:"id"`
	Name        string                   `json:"name"`
	Icon        string                   `json:"icon"`
	Description string                   `json:"description"`
	DataType    string                   `json:"dataType"`
	Unit        string                   `json:"unit"`
	Category    PropertyCategoryResponse `json:"category,omitempty"`
}

func ToPropertyResponse(from dto.Property) PropertyResponse {
	return PropertyResponse{
		Id:          from.Id,
		Name:        from.Name,
		Icon:        from.Icon,
		Description: from.Description,
		DataType:    from.DataType,
		Unit:        from.Unit,
		Category:    ToPropertyCategoryResponse(from.Category),
	}
}

func ToCreateProperty(from CreatePropertyRequest) dto.CreateProperty {
	return dto.CreateProperty{
		CategoryId:  from.CategoryId,
		Name:        from.Name,
		Icon:        from.Icon,
		Description: from.Description,
		DataType:    from.DataType,
		Unit:        from.Unit,
	}
}

func ToUpdateProperty(from UpdatePropertyRequest) dto.UpdateProperty {
	return dto.UpdateProperty{
		CategoryId:  from.CategoryId,
		Name:        from.Name,
		Icon:        from.Icon,
		Description: from.Description,
		DataType:    from.DataType,
		Unit:        from.Unit,
	}
}
