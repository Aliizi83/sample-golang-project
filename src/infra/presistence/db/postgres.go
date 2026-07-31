package db

import (
	"fmt"
	"log"
	"time"

	"github.com/Aliizi83/sample-golang-project/src/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbClient *gorm.DB

func InitDB(cfg *config.Config) error {
	var err error
	cnn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=Asia/Tehran",
		cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.User, cfg.Postgres.Password, cfg.Postgres.DbName, cfg.Postgres.SSLMode)

	dbClient, err = gorm.Open(postgres.Open(cnn), &gorm.Config{})

	if err != nil {
		return err
	}

	sqlDb, err := dbClient.DB()
	if err != nil {
		return err
	}

	err = sqlDb.Ping()
	if err != nil {
		return err
	}

	sqlDb.SetMaxIdleConns(cfg.Postgres.MaxIdleConns)
	sqlDb.SetMaxOpenConns(cfg.Postgres.MaxOpenConns)
	sqlDb.SetConnMaxLifetime(cfg.Postgres.ConnMaxLifetime * time.Minute)

	log.Println("Db connection stablished")

	return nil
}

func GetDB() *gorm.DB {
	return dbClient
}

func CloseDB() {
	db, _ := dbClient.DB()
	db.Close()
}
