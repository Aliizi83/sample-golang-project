package migrations

import (
	"github.com/Aliizi83/sample-golang-project/src/config"
	"github.com/Aliizi83/sample-golang-project/src/domain/models"
	"github.com/Aliizi83/sample-golang-project/src/infra/presistence/db"
	"github.com/Aliizi83/sample-golang-project/src/infra/presistence/seeders"
	"github.com/Aliizi83/sample-golang-project/src/pkg/logging"
	"gorm.io/gorm"
)

var logger = logging.NewLogger(config.GetConfig())

func UpP_1() {
	database := db.GetDB()
	createTables(database)
}

func createTables(database *gorm.DB) {
	modelsToMigrate := []interface{}{
		models.Country{},
		models.City{},
		models.File{},
		models.PersianYear{},
		models.PropertyCategory{},
		models.Property{},
		models.User{},
		models.Role{},
		models.UserRole{},
		models.Company{},
		models.Gearbox{},
		models.Color{},
		models.CarType{},
		models.CarModel{},
		models.CarModelColor{},
		models.CarModelYear{},
		models.CarModelImage{},
		models.CarModelPriceHistory{},
		models.CarModelProperty{},
		models.CarModelComment{},
	}

	err := database.AutoMigrate(modelsToMigrate...)
	if err != nil {
		logger.Error(err, logging.Postgres, logging.Migration, err.Error(), nil)
		return
	}
	logger.Info(logging.Postgres, logging.Migration, "tables migrated successfully", nil)

	if err := AddDefaultData(database); err != nil {
		logger.Error(err, logging.Postgres, logging.Migration, "failed to add default data: "+err.Error(), nil)
	} else {
		logger.Info(logging.Postgres, logging.Migration, "default data added", nil)
	}
}

func AddDefaultData(database *gorm.DB) error {
	if err := seeders.AddDefaultRoles(database); err != nil {
		return err
	}

	if err := seeders.AddDefaultUsers(database); err != nil {
		return err
	}

	seeders.CreateCountries(database)

	return nil
}
