package models

var AllModels []interface{}

func RegisterModel(model interface{}) {
	AllModels = append(AllModels, model)
}
