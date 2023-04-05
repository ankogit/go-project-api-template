package app

import (
	"context"
	"errors"
	"fmt"
	"github.com/robfig/cron/v3"
	"log"
	"math/rand"
	"myapiproject/internal/config"
	"myapiproject/internal/migrations"
	"myapiproject/internal/repository"
	"myapiproject/internal/repository/postgresDB"
	"myapiproject/internal/service"
	"myapiproject/internal/transport/rest"
	"myapiproject/pkg/auth"
	"myapiproject/pkg/email/smtp"
	"myapiproject/pkg/logger"
	nHttp "net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run(configPath string) {
	cfg, err := config.Init(configPath)
	if err != nil {
		log.Panicln(err)

		return
	}

	ormDB, err := postgresDB.NewPostgresDB(postgresDB.Config{
		Host:        cfg.DB.Host,
		Port:        cfg.DB.Port,
		Username:    cfg.DB.Username,
		Password:    cfg.DB.Password,
		DBName:      cfg.DB.DBName,
		SSLMode:     cfg.DB.SSLMode,
		Environment: cfg.Env,
	})
	if err != nil {
		log.Fatalf(fmt.Sprintf("cant to connect postgress: %s", err))
		return
	}

	defer postgresDB.CloseDB(ormDB)

	migrations.RunMigrations(ormDB)

	repositories := repository.NewRepositories(ormDB)

	tokenManager, err := auth.NewManager(cfg.AuthSecret)
	if err != nil {
		log.Fatal(err)
		return
	}

	emailSender, err := smtp.NewSMTPSender(cfg.SMTPConfig.From, cfg.SMTPConfig.Password, cfg.SMTPConfig.Host, cfg.SMTPConfig.Port)
	if err != nil {
		logger.Error(err)

		return
	}

	timeZone, _ := time.LoadLocation("Europe/Moscow")

	scheduler := cron.New(cron.WithLocation(timeZone))
	defer scheduler.Stop()

	services := service.NewServices(service.Deps{
		Repositories: repositories,
		TokenManager: tokenManager,
		EmailSender:  emailSender,
		RedisConfig: service.ConfigRedis{
			Host:     cfg.RedisConfig.Host,
			Port:     cfg.RedisConfig.Port,
			Db:       cfg.RedisConfig.Database,
			Password: cfg.RedisConfig.Password,
		},
		CronScheduler: scheduler,
	})

	rand.Seed(time.Now().UnixNano())

	serverInstance := new(rest.Server)
	handlers := rest.NewHandler(services)

	services.CronService.Init()
	services.CronService.Start()
	go func() {
		if err := serverInstance.RunHttp(cfg, handlers.InitRoutes(cfg)); !errors.Is(err, nHttp.ErrServerClosed) {
			log.Fatalf("error http server: %s", err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	log.Println("Started.")
	<-quit

	if err := serverInstance.Shutdown(context.Background()); err != nil {
		log.Fatalf("error occurated on shuting down server: %s", err.Error())
	}

	postgresDB.CloseDB(ormDB)

	log.Println("Shutdown...")
}
