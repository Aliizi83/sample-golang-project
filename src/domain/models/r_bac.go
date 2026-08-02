package models

type Role struct {
	BaseModel
	Title string `gorm:"type:string;size:20;not null;unique"`
	Name  string `gorm:"type:string;size:20;not null;unique"`
}

type UserRole struct {
	BaseModel
	Role   Role `gorm:"foreignKey:RoleId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	User   User `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserId uint
	RoleId uint
}
