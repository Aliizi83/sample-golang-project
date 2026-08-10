package dto

type CreateColor struct {
	Name    string
	HexCode string
}

type UpdateColor struct {
	Name    string
	HexCode string
}

type Color struct {
	Id      int
	Name    string
	HexCode string
}
