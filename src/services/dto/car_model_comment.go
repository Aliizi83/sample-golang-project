package dto

type CreateCarModelComment struct {
	CarModelId int
	UserId     int
	Message    string
}

type UpdateCarModelComment struct {
	Message string
}

type CarModelComment struct {
	Id         int
	CarModelId int
	User       User
	Message    string
}

type User struct {
	Id        int
	Username  string
	FirstName string
	LastName  string
	Email     string
}
