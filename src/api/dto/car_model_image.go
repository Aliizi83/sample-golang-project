package dto

import "github.com/Aliizi83/sample-golang-project/src/services/dto"

type CreateCarModelImageRequest struct {
	CarModelId  int  `json:"carModelId" binding:"required"`
	ImageId     int  `json:"imageId" binding:"required"`
	IsMainImage bool `json:"isMainImage"`
}

type UpdateCarModelImageRequest struct {
	IsMainImage bool `json:"isMainImage,omitempty"`
}

type CarModelImageResponse struct {
	Id          int          `json:"id"`
	CarModelId  int          `json:"carModelId,omitempty"`
	Image       FileResponse `json:"image,omitempty"`
	IsMainImage bool         `json:"isMainImage"`
}

func ToCarModelImageResponse(from dto.CarModelImage) CarModelImageResponse {
	return CarModelImageResponse{
		Id:          from.Id,
		CarModelId:  from.CarModelId,
		IsMainImage: from.IsMainImage,
		Image:       ToFileResponse(from.Image),
	}
}

func ToCreateCarModelImage(from CreateCarModelImageRequest) dto.CreateCarModelImage {
	return dto.CreateCarModelImage{
		CarModelId:  from.CarModelId,
		ImageId:     from.ImageId,
		IsMainImage: from.IsMainImage,
	}
}

func ToUpdateCarModelImage(from UpdateCarModelImageRequest) dto.UpdateCarModelImage {
	return dto.UpdateCarModelImage{
		IsMainImage: from.IsMainImage,
	}
}
