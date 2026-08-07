package dto

type CreatePropertyCategory struct {
	Name string
	Icon string
}

type UpdatePropertyCategory struct {
	Name string
	Icon string
}

type PropertyCategory struct {
	Id         int
	Name       string
	Icon       string
	Properties []Property
}
