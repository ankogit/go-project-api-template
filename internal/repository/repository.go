package repository

import (
	"gorm.io/gorm"
	"myapiproject/internal/repository/postgresDB"
)

type Repositories struct {
	Users         UserRepository
	RefreshTokens RefreshTokenRepository
}

func NewRepositories(ormDB *gorm.DB) *Repositories {
	return &Repositories{
		RefreshTokens: postgresDB.NewRefreshTokenRepository(ormDB),
		Users:         postgresDB.NewUserRepository(ormDB),
	}
}
