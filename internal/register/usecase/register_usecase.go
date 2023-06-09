package register

import (
	dto "edtech-app/internal/register/dto"
	userDto "edtech-app/internal/user/dto"
	userUseCase "edtech-app/internal/user/usecase"
	mail "edtech-app/pkg/mail/sendgrid"
)

type RegisterUseCase interface {
	Register(userDto userDto.UserRequestBody) error
}

type RegisterUseCaseImpl struct {
	userUseCase userUseCase.UserUseCase
	mail        mail.Mail
}

func NewRegisterUseCase(
	userUseCase userUseCase.UserUseCase,
	mail mail.Mail,
) RegisterUseCase {
	return &RegisterUseCaseImpl{userUseCase, mail}
}

func (ru *RegisterUseCaseImpl) Register(userDto userDto.UserRequestBody) error {
	// Check ke dalam module user
	user, err := ru.userUseCase.Create(userDto)

	if err != nil {
		return err
	}

	// Kirim email melalui sendgrid
	email := dto.CreateEmailVerification{
		SUBJECT:           "Kode verifikasi",
		EMAIL:             user.Email,
		VERIFICATION_CODE: user.CodeVerified,
	}

	go ru.mail.SendVerificationEmail(user.Email, email)

	return nil
}
