package services

import (
	"BACKEND_GO/internal/domain/profile/models"
	"errors"
)

type ProfileUseCase struct {
	pr models.ProfileRepositoryImpl
}

func NewProfileUseCase(pr models.ProfileRepositoryImpl) *ProfileUseCase {
	return &ProfileUseCase{pr: pr}
}

func (c *ProfileUseCase) GetProfileService(id string) (*models.Profile, error) {
	profile, err := c.pr.GetProfile(id)
	if err != nil {
		return nil, errors.Unwrap(err)
	}

	return &profile, nil
}
