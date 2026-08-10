package dependencies

import (
	"github.com/Aliizi83/sample-golang-project/src/config"
	"github.com/Aliizi83/sample-golang-project/src/domain/models"
	"github.com/Aliizi83/sample-golang-project/src/domain/repositories"
	"github.com/Aliizi83/sample-golang-project/src/infra/presistence/db"
	"github.com/Aliizi83/sample-golang-project/src/infra/repositories/postgres_repositories"
)

func GetUserRepository(cfg *config.Config) repositories.UserRepository {
	return postgres_repositories.NewPostgresUserRepository(cfg)
}

func GetCountryRepository(cfg *config.Config) repositories.CountryRepository {
	return postgres_repositories.NewBaseRepository[models.Country](cfg, []db.PreloadEntity{{Entity: "Cities"}, {Entity: "Companies"}})
}

func GetCityRepository(cfg *config.Config) repositories.CityRepository {
	return postgres_repositories.NewBaseRepository[models.City](cfg, []db.PreloadEntity{{Entity: "Country"}})
}

func GetFileRepository(cfg *config.Config) repositories.FileRepository {
	return postgres_repositories.NewBaseRepository[models.File](cfg, nil)
}

func GetPropertyRepository(cfg *config.Config) repositories.PropertyRepository {
	return postgres_repositories.NewBaseRepository[models.Property](cfg, []db.PreloadEntity{{Entity: "Category"}})
}

func GetPropertyCategoryRepository(cfg *config.Config) repositories.PropertyCategoryRepository {
	return postgres_repositories.NewBaseRepository[models.PropertyCategory](cfg, []db.PreloadEntity{{Entity: "Properties"}})
}

func GetGearboxRepository(cfg *config.Config) repositories.GearboxRepository {
	return postgres_repositories.NewBaseRepository[models.Gearbox](cfg, []db.PreloadEntity{})
}

//Generated dependency function

func GetCarTypeRepository(cfg *config.Config) repositories.CarTypeRepository {
	return postgres_repositories.NewBaseRepository[models.CarType](cfg, []db.PreloadEntity{})
}

//Generated dependency function

func GetCompanyRepository(cfg *config.Config) repositories.CompanyRepository {
	return postgres_repositories.NewBaseRepository[models.Company](cfg, []db.PreloadEntity{{Entity: "Country"}})
}

//Generated dependency function

func GetColorRepository(cfg *config.Config) repositories.ColorRepository {
	return postgres_repositories.NewBaseRepository[models.Color](cfg, []db.PreloadEntity{})
}