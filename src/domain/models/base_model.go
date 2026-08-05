package models

import (
	"context"
	"time"

	"github.com/Aliizi83/sample-golang-project/src/constants"
	"github.com/Aliizi83/sample-golang-project/src/domain/helpers"
	"gorm.io/gorm"
)

const UniqueCodeLength = 8

type BaseModel struct {
	Id         uint       `gorm:"primaryKey"`
	Code       string     `gorm:"type:varchar(32);not null;unique"`
	CreatedAt  time.Time  `gorm:"TIMESTAMP with time zone; not null"`
	ModifiedAt *time.Time `gorm:"TIMESTAMP with time zone;default:null"`
	DeletedAt  *time.Time `gorm:"TIMESTAMP with time zone;default:null"`
	CreatedBy  int        `gorm:"not null"`
	ModifiedBy *int       `gorm:"default:null"`
	DeletedBy  *int       `gorm:"default:null"`
}

func (m *BaseModel) BeforeCreate(tx *gorm.DB) error {
	userId := m.getUserIdFromClaims(tx.Statement.Context)
	m.Code = helpers.GenerateRandomString(UniqueCodeLength)
	m.CreatedAt = time.Now()
	m.CreatedBy = userId
	return nil
}

func (m *BaseModel) BeforeUpdate(tx *gorm.DB) error {
	now := time.Now().UTC()
	userId := m.getUserIdFromClaims(tx.Statement.Context)
	m.ModifiedAt = &now
	m.ModifiedBy = &userId
	return nil
}

func (m *BaseModel) BeforeDelete(tx *gorm.DB) error {
	now := time.Now().UTC()
	userId := m.getUserIdFromClaims(tx.Statement.Context)

	m.DeletedAt = &now
	m.DeletedBy = &userId
	return nil
}

func (m *BaseModel) getUserIdFromClaims(ctx context.Context) int {
	userId := -1
	claimsVal := ctx.Value(constants.ClaimsKey)
	if claimsVal != nil {
		claimsMap := claimsVal.(map[string]any)
		value, ok := claimsMap[constants.UserIdKey]
		if !ok {
			return userId
		}

		userId = int(value.(float64))
	}
	return userId
}
