package models

type Permission struct {
	BaseModel
	Title           string `gorm:"type:string;size:20;not null;unique"`
	Name            string `gorm:"type:string;size:20;not null;unique"`
	RolePermissions *[]RolePermission
}

type Role struct {
	BaseModel
	Title           string `gorm:"type:string;size:20;not null;unique"`
	Name            string `gorm:"type:string;size:20;not null;unique"`
	RolePermissions *[]RolePermission
}

type RolePermission struct {
	BaseModel
	RoleId       uint
	PermissionId uint
	Role         Role       `gorm:"foreignKey:RoleID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Permissions  Permission `gorm:"foreignKey:PermissionID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type UserRole struct {
	BaseModel
	Role   Role `gorm:"foreignKey:RoleId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	User   User `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserId uint
	RoleId uint
}
