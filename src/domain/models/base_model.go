package models

import (
	"database/sql"
	"time"

	"github.com/Aliizi83/sample-golang-project/src/domain/helpers"
	"gorm.io/gorm"
)

const UniqueCodeLength = 8

type BaseModel struct {
	Id         uint          `gorm:"primaryKey"`
	Code       string        `gorm:"type:varchar(32);not null;unique"`
	CreatedAt  time.Time     `gorm:"TIMESTAMP with time zone; not null"`
	ModifiedAt sql.NullTime  `gorm:"TIMESTAMP with time zone;null"`
	DeletedAt  sql.NullTime  `gorm:"TIMESTAMP with time zone;null"`
	CreatedBy  int           `gorm:"not null"`
	ModifiedBy sql.NullInt64 `gorm:"null"`
	DeletedBy  sql.NullInt64 `gorm:"null"`
}

func (m *BaseModel) BeforeCreate(tx *gorm.DB) error {
	value := tx.Statement.Context.Value("UserId")
	userId := -1

	if value != nil {
		userId = int(value.(float64))
	}
	m.Code = helpers.GenerateRandomString(UniqueCodeLength)
	m.CreatedAt = time.Now()
	m.CreatedBy = userId
	return nil
}

func (m *BaseModel) BeforeUpdate(tx *gorm.DB) error {
	value := tx.Statement.Context.Value("UserId")
	userId := &sql.NullInt64{Valid: false}

	if value != nil {
		userId = &sql.NullInt64{Valid: true, Int64: int64(value.(float64))}
	}

	m.ModifiedAt = sql.NullTime{Time: time.Now().UTC(), Valid: true}
	m.ModifiedBy = *userId
	return nil
}

func (m *BaseModel) BeforeDelete(tx *gorm.DB) error {
	value := tx.Statement.Context.Value("UserId")
	userId := &sql.NullInt64{Valid: false}

	if value != nil {
		userId = &sql.NullInt64{Valid: true, Int64: int64(value.(float64))}
	}

	m.DeletedAt = sql.NullTime{Time: time.Now().UTC(), Valid: true}
	m.DeletedBy = *userId
	return nil
}
