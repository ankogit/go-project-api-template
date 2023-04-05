package service

import (
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
	"myapiproject/internal/repository"
	"myapiproject/pkg/auth"
	"myapiproject/pkg/email"
)

type LoginInput struct {
	UserId uint
}

type RefreshInput struct {
	Token string
}

type Services struct {
	Users        Users
	TokenManager auth.TokenManager
	Repositories *repository.Repositories
	EmailService *EmailService
	RedisService *RedisService
	CronService  *CronService
}
type Deps struct {
	Repositories  *repository.Repositories
	TokenManager  auth.TokenManager
	Logger        *logrus.Logger
	EmailSender   email.Sender
	RedisConfig   ConfigRedis
	CronScheduler *cron.Cron
}

func NewServices(deps Deps) *Services {
	emailsService := NewEmailsService(deps.EmailSender)
	redisService := NewRedisService(deps.RedisConfig)
	usersService := NewUsersService(deps.Repositories.Users, deps.Repositories.RefreshTokens, deps.TokenManager, redisService)
	cronService := NewCronService(deps.CronScheduler, usersService)

	return &Services{
		Users:        usersService,
		TokenManager: deps.TokenManager,
		Repositories: deps.Repositories,
		EmailService: emailsService,
		RedisService: redisService,
		CronService:  cronService,
	}
}
