package dto

import "github.com/Aliizi83/sample-golang-project/src/services/dto"

type CreateCarModelCommentRequest struct {
	CarModelId int    `json:"carModelId" binding:"required"`
	Message    string `json:"message" binding:"required,max=100"`
}

type UpdateCarModelCommentRequest struct {
	Message string `json:"message" binding:"required,max=100"`
}

type CarModelCommentResponse struct {
	Id         int          `json:"id"`
	CarModelId int          `json:"carModelId"`
	User       UserResponse `json:"user"`
	Message    string       `json:"message"`
}

type UserResponse struct {
	Id        int    `json:"id"`
	Username  string `json:"username"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

func ToCarModelCommentResponse(from dto.CarModelComment) CarModelCommentResponse {
	return CarModelCommentResponse{
		Id:         from.Id,
		CarModelId: from.CarModelId,
		Message:    from.Message,
		User:       ToUserResponse(from.User),
	}
}

func ToCreateCarModelComment(from CreateCarModelCommentRequest) dto.CreateCarModelComment {
	return dto.CreateCarModelComment{
		CarModelId: from.CarModelId,
		Message:    from.Message,
	}
}

func ToUpdateCarModelComment(from UpdateCarModelCommentRequest) dto.UpdateCarModelComment {
	return dto.UpdateCarModelComment{
		Message: from.Message,
	}
}

func ToUserResponse(from dto.User) UserResponse {
	return UserResponse{
		Id:        from.Id,
		Username:  from.Username,
		FirstName: from.FirstName,
		LastName:  from.LastName,
		Email:     from.Email,
	}
}
