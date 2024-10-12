package repositories

import (
	"BACKEND_GO/internal/domain/profile/models"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type mysqlRepository struct {
	db *gorm.DB
}

func NewMysqlRepository(db *gorm.DB) models.ProfileRepositoryImpl {
	return &mysqlRepository{db: db}
}

func (r *mysqlRepository) GetProfile(id string) (models.Profile, error) {
	var profile models.Profile

	result := r.db.First(&profile, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return profile, nil // Return empty profile if not found
	}
	if result.Error != nil {
		return profile, errors.Wrap(result.Error, "failed to get profile")
	}

	return profile, nil
}
