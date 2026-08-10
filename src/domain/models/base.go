package models

import "time"

type Country struct {
	BaseModel
	Name      string `gorm:"size:15;type:string;not null;"`
	Cities    []City
	Companies []Company
}

//CRUD done

type City struct {
	BaseModel
	Name      string `gorm:"size:10;type:string;not null;"`
	CountryId int
	Country   Country `gorm:"foreignKey:CountryId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
}

//CRUD done

type PersianYear struct {
	BaseModel
	PersianTitle  string    `gorm:"size:10;type:string;not null;unique"`
	Year          int       `gorm:"type:int;uniqueIndex;not null"`
	StartAt       time.Time `gorm:"type:TIMESTAMP with time zone;not null;unique"`
	EndAt         time.Time `gorm:"type:TIMESTAMP with time zone;not null;unique"`
	CarModelYears []CarModelYear
}

//CRUD done

type Color struct {
	BaseModel
	Name           string `gorm:"size:15;type:string;not null,unique"`
	HexCode        string `gorm:"size:7;type:string;not null,unique"`
	CarModelColors []CarModelColor
}

//CRUD done

type File struct {
	BaseModel
	Name        string `gorm:"size:100;type:string;not null"`
	Directory   string `gorm:"size:100;type:string;not null"`
	Description string `gorm:"size:500;type:string;not null"`
	MimeType    string `gorm:"size:20;type:string;not null"`
}

//CRUD done

func init() {
	RegisterModel(&Country{})
	RegisterModel(&City{})
	RegisterModel(&PersianYear{})
	RegisterModel(&Color{})
	RegisterModel(&File{})
}
