package repositories

import (
	"context"

	"github.com/Aliizi83/sample-golang-project/src/domain/filters"
	"github.com/Aliizi83/sample-golang-project/src/domain/models"
)

type BaseRepository[TModel models.Identifiable] interface {
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

// Auto generated repository
type GearboxRepository interface {
	BaseRepository[models.Gearbox]
} //Auto generated repository
type CarTypeRepository interface {
	BaseRepository[models.CarType]
} //Auto generated repository
type CompanyRepository interface {
	BaseRepository[models.Company]
} //Auto generated repository
type ColorRepository interface {
	BaseRepository[models.Color]
} //Auto generated repository
type PersianYearRepository interface {
	BaseRepository[models.PersianYear]
} //Auto generated repository
type CarModelRepository interface {
	BaseRepository[models.CarModel]
}

// Auto generated repository
type CarModelColorRepository interface {
	BaseRepository[models.CarModelColor]
}
//Auto generated repository
type CarModelYearRepository interface {
	BaseRepository[models.CarModelYear]
}//Auto generated repository
type CarModelPriceHistoryRepository interface {
	BaseRepository[models.CarModelPriceHistory]
}//Auto generated repository
type CarModelImageRepository interface {
	BaseRepository[models.CarModelImage]
}//Auto generated repository
type CarModelPropertyRepository interface {
	BaseRepository[models.CarModelProperty]
}