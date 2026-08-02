package dependencies

import (
	"github.com/Aliizi83/sample-golang-project/src/config"
	"github.com/Aliizi83/sample-golang-project/src/domain/repositories"
	"github.com/Aliizi83/sample-golang-project/src/infra/repositories/postgres_repositories"
)

func GetUserRepository(cfg *config.Config) repositories.UserRepository {
	return postgres_repositories.NewPostgresUserRepository(cfg)
}
