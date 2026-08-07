package dto

type CreateProperty struct {
	Name        string
	CategoryId  int
	Icon        string
	Description string
	DataType    string
	Unit        string
}

type UpdateProperty struct {
	Name        string
	CategoryId  int
	Icon        string
	Description string
	DataType    string
	Unit        string
}

type Property struct {
	Id          int
	Name        string
	Icon        string
	Description string
	DataType    string
	Unit        string
	Category    PropertyCategory
}
