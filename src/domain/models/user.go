package models

type User struct {
	BaseModel
	Username     string `gorm:"type:string;size:20;not null;unique"`
	FirstName    string `gorm:"type:string;size:20;null"`
	LastName     string `gorm:"type:string;size:20;null"`
	MobileNumber string `gorm:"type:string;size:15;null"`
	Email        string `gorm:"type:string;size:100;null"`
	Password     string `gorm:"type:string;size:255;not null"`
	Enabled      bool   `gorm:"type:bool;default:true"`
}
