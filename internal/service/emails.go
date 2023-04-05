package service

import (
	emailProvider "myapiproject/pkg/email"
)

type EmailService struct {
	sender emailProvider.Sender
}

func NewEmailsService(sender emailProvider.Sender) *EmailService {
	return &EmailService{
		sender: sender,
	}
}
