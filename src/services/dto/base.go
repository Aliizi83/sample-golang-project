package dto

import (
	"mime/multipart"
)

type IdName struct {
	Id   int
	Name string
}
type Name struct {
	Name string
}

type Country struct {
	IdName
	Cities    []City
	Companies []Company
}

type CreateCity struct {
	Name      string
	CountryId int
}

type UpdateCity struct {
	Name      string
	CountryId int
}
type City struct {
	IdName
	Country Country
}

type CreateFile struct {
	Name        string
	File        *multipart.FileHeader
	Directory   string
	Description string
	MimeType    string
}

type UpdateFile struct {
	Description string
}

type File struct {
	IdName
	Directory   string
	Description string
	MimeType    string
}

type PersianYearWithoutDate struct {
	Id           int
	PersianTitle string
	Year         int
}
