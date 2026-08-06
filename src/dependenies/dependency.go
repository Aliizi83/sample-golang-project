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
