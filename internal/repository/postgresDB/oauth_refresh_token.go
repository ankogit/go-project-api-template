package postgresDB

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"myapiproject/internal/models"
	"time"
)

type RefreshTokenRepository struct {
	DB *gorm.DB
}

func NewRefreshTokenRepository(db *gorm.DB) *RefreshTokenRepository {
	return &RefreshTokenRepository{DB: db}
}

func (r RefreshTokenRepository) Create(userId uint, refreshToken string, expiresIn time.Time) (models.RefreshToken, error) {
	rt := models.RefreshToken{
		ID:           0,
		UserID:       userId,
		RefreshToken: refreshToken,
		ExpiresIn:    expiresIn,
	}

	err := r.DB.Clauses(clause.Returning{}).Create(&rt).Error
	return rt, err
}

func (r RefreshTokenRepository) Find(refreshToken string) (models.RefreshToken, error) {

	var rt models.RefreshToken

	err := r.DB.Where("expires_in >= ?", time.Now()).First(&rt, "refresh_token = ?", refreshToken).Error
	return rt, err
}
