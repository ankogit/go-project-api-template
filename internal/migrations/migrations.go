package migrations

import (
	"fmt"
	"gorm.io/gorm"
	"log"
	"myapiproject/internal/models"
)

type Migration interface {
	Migrate(*gorm.DB) error
}

func RunMigrations(ormDB *gorm.DB) {
	if err := ormDB.AutoMigrate(
		models.User{},
		models.Profile{},
		models.RefreshToken{},
	); err != nil {
		log.Fatalf(fmt.Sprintf("migraion process was failed: %s", err))
	}
}
