package repositories

import (
	"context"

	"github.com/Aliizi83/sample-golang-project/src/domain/filters"
	"github.com/Aliizi83/sample-golang-project/src/domain/models"
)

type BaseRepository[TModel any] interface {
	Create(ctx context.Context, model TModel) (TModel, error)
	Update(ctx context.Context, id int, updateFields map[string]any) (TModel, error)
	Delete(ctx context.Context, id int) error
	GetById(ctx context.Context, id int) (TModel, error)
	GetByFilter(ctx context.Context, req filters.PaginationInputWithFilter) (int64, *[]TModel, error)
}

type CountryRepository interface {
	BaseRepository[models.Country]
}

type CityRepository interface {
	BaseRepository[models.City]
}

type FileRepository interface {
	BaseRepository[models.File]
}

type PropertyRepository interface {
	BaseRepository[models.Property]
}

type PropertyCategoryRepository interface {
	BaseRepository[models.PropertyCategory]
}
//Auto generated repository
type GearboxRepository interface {
	BaseRepository[models.Gearbox]
}