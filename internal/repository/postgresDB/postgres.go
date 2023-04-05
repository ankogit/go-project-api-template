package postgresDB

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"log"
)

type Config struct {
	Host        string
	Port        string
	Username    string
	Password    string
	DBName      string
	SSLMode     string
	Environment string
}

func NewPostgresDB(config Config) (*gorm.DB, error) {
	dbLogger := gormLogger.Default.LogMode(gormLogger.Error)
	if config.Environment == "dev" {
		dbLogger = gormLogger.Default.LogMode(gormLogger.Info)
	}

	ormDB, err := gorm.Open(postgres.Open(fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", config.Host, config.Port, config.Username, config.Password, config.DBName, config.SSLMode)), &gorm.Config{
		Logger:                                   dbLogger,
		DisableForeignKeyConstraintWhenMigrating: true,
		SkipDefaultTransaction:                   true,
	})
	if err != nil {
		return nil, err
	}
	ormDB.Statement.RaiseErrorOnNotFound = true

	return ormDB, nil
}

func CloseDB(ormDB *gorm.DB) {
	db, err := ormDB.DB()
	if err != nil {
		log.Fatalf(fmt.Sprintf("cant close the connector: %s", err))
	}

	err = db.Close()
	if err != nil {
		log.Fatalf(fmt.Sprintf("cant close the connector: %s", err))
	}
}
