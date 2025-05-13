package user

import (
	"context"
	"sync"
	
	mailDomain "github.com/onion0904/CarShareSystem/app/domain/mail"
)

type SendEmailUseCase struct {
	mailService mailDomain.MailService
	VerificationCodes map[string]string
	CodeMutex         sync.Mutex
}

func NewSendEmailUseCase(
	mailService mailDomain.MailService,
) *SendEmailUseCase {
	return &SendEmailUseCase{
		mailService:       mailService,
		VerificationCodes: make(map[string]string),
	}
}

type SendEmailUseCaseDto struct {
	Email string
	Code  string
}

func (uc *SendEmailUseCase) Run(ctx context.Context, dto SendEmailUseCaseDto) error {
	return uc.mailService.SendEmail(dto.Email, dto.Code)
}