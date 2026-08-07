package seeders

import (
	"github.com/Aliizi83/sample-golang-project/src/domain/models"
	"gorm.io/gorm"
)

func CreateCountries(database *gorm.DB) error {
	count := 0
	database.
		Model(&models.Country{}).
		Select("count(*)").
		Find(&count)
	if count == 0 {
		if err := database.Create(&models.Country{Name: "Iran", Cities: []models.City{
			{Name: "Tehran"},
			{Name: "Isfahan"},
			{Name: "Shiraz"},
			{Name: "Chalus"},
			{Name: "Ahwaz"},
		}}).Error; err != nil {
			return err
		}
		if err := database.Create(&models.Country{Name: "USA", Cities: []models.City{
			{Name: "New York"},
			{Name: "Washington"},
		}}).Error; err != nil {
			return err
		}
		if err := database.Create(&models.Country{Name: "Germany", Cities: []models.City{
			{Name: "Berlin"},
			{Name: "Munich"},
		}}).Error; err != nil {
			return err
		}
		if err := database.Create(&models.Country{Name: "China", Cities: []models.City{
			{Name: "Beijing"},
			{Name: "Shanghai"},
		}}).Error; err != nil {
			return err
		}
		if err := database.Create(&models.Country{Name: "Italy", Cities: []models.City{
			{Name: "Roma"},
			{Name: "Turin"},
		}}).Error; err != nil {
			return err
		}
		if err := database.Create(&models.Country{Name: "France", Cities: []models.City{
			{Name: "Paris"},
			{Name: "Lyon"},
		}}).Error; err != nil {
			return err
		}
		if err := database.Create(&models.Country{Name: "Japan", Cities: []models.City{
			{Name: "Tokyo"},
			{Name: "Kyoto"},
		}}).Error; err != nil {
			return err
		}
		if err := database.Create(&models.Country{Name: "South Korea", Cities: []models.City{
			{Name: "Seoul"},
			{Name: "Ulsan"},
		}}).Error; err != nil {
			return err
		}
	}
	return nil
}
