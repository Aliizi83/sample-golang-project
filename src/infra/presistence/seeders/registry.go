package seeders

import "gorm.io/gorm"

type SeederFunc func(database *gorm.DB) error

var Seeders []SeederFunc = []SeederFunc{
	seedRoles,
	seedUsers,
	seedCountries,
	seedProperties,
	seedCarRelatedModels,
}
