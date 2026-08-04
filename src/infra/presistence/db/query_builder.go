package db

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/Aliizi83/sample-golang-project/src/common"
	"github.com/Aliizi83/sample-golang-project/src/domain/filters"
	"gorm.io/gorm"
)

const softDeleteEx string = "deleted_by IS NULL"

type PreloadEntity struct {
	Entity string
}

func GenerateDynamicQuery[T any](filter *filters.DynamicFilter) string {
	t := new(T)
	typeT := reflect.TypeOf(*t)
	query := make([]string, 0)
	query = append(query, softDeleteEx)
	if filter.Filters != nil {
		for name, filter := range filter.Filters {
			if fld, ok := typeT.FieldByName(name); ok {
				query = append(query, GenerateDynamicFilter(fld, filter))
			}
		}
	}

	return strings.Join(query, " AND ")
}

func GenerateDynamicFilter(fld reflect.StructField, filter filters.Filter) string {
	conditionQuery := ""
	fld.Name = common.ToSnakeCase(fld.Name)
	switch filter.Type {
	case "contains":
		conditionQuery = fmt.Sprintf("%s ILIKE '%%%s%%'", fld.Name, filter.From)
	case "notContains":
		conditionQuery = fmt.Sprintf("%s NOT ILIKE '%%%s%%'", fld.Name, filter.From)
	case "startsWith":
		conditionQuery = fmt.Sprintf("%s ILIKE '%s%%'", fld.Name, filter.From)
	case "endsWith":
		conditionQuery = fmt.Sprintf("%s ILIKE '%%%s'", fld.Name, filter.From)
	case "equals":
		conditionQuery = fmt.Sprintf("%s ILIKE '%s'", fld.Name, filter.From)
	case "notEquals":
		conditionQuery = fmt.Sprintf("%s NOT ILIKE '%s'", fld.Name, filter.From)
	case "lessThan":
		conditionQuery = fmt.Sprintf("%s < '%s'", fld.Name, filter.From)
	case "lessThanOrEqual":
		conditionQuery = fmt.Sprintf("%s <= '%s'", fld.Name, filter.From)
	case "greaterThan":
		conditionQuery = fmt.Sprintf("%s > '%s'", fld.Name, filter.From)
	case "greaterThanOrEqual":
		conditionQuery = fmt.Sprintf("%s >= '%s'", fld.Name, filter.From)
	case "inRange":
		if fld.Type.Kind() == reflect.String {
			conditionQuery = fmt.Sprintf("%s >= '%s%%' AND ", fld.Name, filter.From)
			conditionQuery += fmt.Sprintf("%s <= '%%%s'", fld.Name, filter.To)
		} else {
			conditionQuery = fmt.Sprintf("%s >= %s AND ", fld.Name, filter.From)
			conditionQuery += fmt.Sprintf("%s <= %s", fld.Name, filter.To)
		}

	}

	return conditionQuery
}

func GenerateDynamicSort[T any](filter *filters.DynamicFilter) string {
	t := new(T)
	typeT := reflect.TypeOf(*t)
	sort := make([]string, 0)
	if filter.Sorts != nil {
		for _, tp := range *filter.Sorts {
			fld, ok := typeT.FieldByName(tp.ColumnId)
			if ok && (tp.Sort == "asc" || tp.Sort == "desc") {
				fld.Name = common.ToSnakeCase(fld.Name)
				sort = append(sort, fmt.Sprintf("%s %s", fld.Name, tp.Sort))
			}
		}
	}
	return strings.Join(sort, ", ")
}

// Preload
func Preload(db *gorm.DB, preloads []PreloadEntity) *gorm.DB {
	for _, item := range preloads {
		db = db.Preload(item.Entity)
	}
	return db
}
