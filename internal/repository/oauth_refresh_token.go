package repository

import (
	"myapiproject/internal/models"
	"time"
)

type RefreshTokenRepository interface {
	Create(userId uint, refreshToken string, expiresIn time.Time) (models.RefreshToken, error)
	Find(refreshToken string) (models.RefreshToken, error)
}
