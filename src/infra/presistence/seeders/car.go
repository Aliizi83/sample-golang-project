package seeders

import (
	"time"

	"github.com/Aliizi83/sample-golang-project/src/domain/models"
	"gorm.io/gorm"
)

var carTypes []models.CarType = []models.CarType{
	{Name: "Crossover"},
	{Name: "Sedan"},
	{Name: "Sports"},
	{Name: "Coupe"},
	{Name: "Hatchback"},
}

var gearboxes []models.Gearbox = []models.Gearbox{
	{Name: "Manual"},
	{Name: "Automatic"},
}

var colors []models.Color = []models.Color{
	{Name: "Black", HexCode: "#000000"},
	{Name: "White", HexCode: "#ffffff"},
	{Name: "Blue", HexCode: "#0000ff"},
}

var years []models.PersianYear = []models.PersianYear{
	{PersianTitle: "1402", Year: 1402, StartAt: time.Date(2023, time.Month(3), 21, 0, 0, 0, 0, time.UTC), EndAt: time.Date(2024, time.Month(3), 20, 0, 0, 0, 0, time.UTC)},
	{PersianTitle: "1401", Year: 1401, StartAt: time.Date(2022, time.Month(3), 21, 0, 0, 0, 0, time.UTC), EndAt: time.Date(2023, time.Month(3), 21, 0, 0, 0, 0, time.UTC)},
	{PersianTitle: "1400", Year: 1400, StartAt: time.Date(2021, time.Month(3), 21, 0, 0, 0, 0, time.UTC), EndAt: time.Date(2022, time.Month(3), 21, 0, 0, 0, 0, time.UTC)},
	{PersianTitle: "1399", Year: 1399, StartAt: time.Date(2020, time.Month(3), 20, 0, 0, 0, 0, time.UTC), EndAt: time.Date(2021, time.Month(3), 21, 0, 0, 0, 0, time.UTC)},
	{PersianTitle: "1398", Year: 1398, StartAt: time.Date(2019, time.Month(3), 21, 0, 0, 0, 0, time.UTC), EndAt: time.Date(2020, time.Month(3), 20, 0, 0, 0, 0, time.UTC)},
}

func seedCarRelatedModels(database *gorm.DB) error {
	count := 0

	database.Select("COUNT(*)").Model(&models.CarType{}).Find(&count)
	if count == 0 {
		for _, m := range carTypes {
			if err := database.Model(&m).Create(&m).Error; err != nil {
				return err
			}
		}
	}

	database.Select("COUNT(*)").Model(&models.Gearbox{}).Find(&count)
	if count == 0 {
		for _, m := range gearboxes {
			if err := database.Model(&m).Create(&m).Error; err != nil {
				return err
			}
		}
	}

	database.Select("COUNT(*)").Model(&models.Color{}).Find(&count)
	if count == 0 {
		for _, m := range colors {
			if err := database.Model(&m).Create(&m).Error; err != nil {
				return err
			}
		}
	}

	database.Select("COUNT(*)").Model(&models.PersianYear{}).Find(&count)
	if count == 0 {
		for _, m := range years {
			if err := database.Model(&m).Create(&m).Error; err != nil {
				return err
			}
		}
	}
	return nil
}
