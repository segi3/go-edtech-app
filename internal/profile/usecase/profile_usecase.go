package profile

import (
	dto "edtech-app/internal/profile/dto"
	userUseCase "edtech-app/internal/user/usecase"
)

type ProfileUseCase interface {
	GetProfile(id int) (*dto.ProfileRespondBody, error)
}

type ProfileUseCaseImpl struct {
	userUseCase userUseCase.UserUseCase
}

// GetProfile implements ProfileUseCase
func (usecase *ProfileUseCaseImpl) GetProfile(id int) (*dto.ProfileRespondBody, error) {
	// Get profile
	user, err := usecase.userUseCase.FindById(id)

	if err != nil {
		return nil, err
	}

	userResponse := dto.CreateProfileResponse(*user)

	return &userResponse, nil
}

func NewProfileUseCase(userUseCase userUseCase.UserUseCase) ProfileUseCase {
	return &ProfileUseCaseImpl{userUseCase}
}
