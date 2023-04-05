package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
	"log"
	"os"
	"time"
)

type Config struct {
	Debug      string `envconfig:"DEBUG"`
	AppUrl     string `envconfig:"APP_URL"`
	AuthSecret string `envconfig:"AUTH_SECRET"`
	Env        string `envconfig:"ENV"`
	TimeZone   string `envconfig:"TZ"`
	DB         struct {
		Host     string `envconfig:"DB_HOST"`
		Port     string `envconfig:"DB_PORT"`
		Username string `envconfig:"DB_USERNAME"`
		Password string `envconfig:"DB_PASSWORD"`
		DBName   string `envconfig:"DB_DATABASE"`
		SSLMode  string `envconfig:"DB_SSL_MODE"`
	}

	HTTPConfig struct {
		Port               string        `yaml:"port"`
		ReadTimeout        time.Duration `yaml:"readTimeout"`
		WriteTimeout       time.Duration `yaml:"writeTimeout"`
		MaxHeaderMegabytes int           `yaml:"maxHeaderBytes"`
	} `yaml:"http"`

	SMTPConfig struct {
		Host     string `envconfig:"MAIL_HOST"`
		Port     int    `envconfig:"MAIL_PORT"`
		Username string `envconfig:"MAIL_USERNAME"`
		Password string `envconfig:"MAIL_PASSWORD"`
		From     string `envconfig:"MAIL_FROM_ADDRESS"`
	}

	RedisConfig struct {
		Url      string `envconfig:"REDIS_URL"`
		Host     string `envconfig:"REDIS_HOST"`
		Port     string `envconfig:"REDIS_PORT"`
		Password string `envconfig:"REDIS_PASSWORD"`
		Database int    `envconfig:"REDIS_CACHE_DB" default:"1"`
	}

	FirebaseConfig struct {
		Type                    string `envconfig:"FB_TYPE"`
		ProjectId               string `envconfig:"FB_PROJECT_ID"`
		PrivateKeyId            string `envconfig:"FB_PRIVATE_KEY_ID"`
		PrivateKey              string `envconfig:"FB_PRIVATE_KEY"`
		ClientEmail             string `envconfig:"FB_CLIENT_EMAIL"`
		ClientId                string `envconfig:"FB_CLIENT_ID"`
		AuthUri                 string `envconfig:"FB_AUTH_URI"`
		TokenUri                string `envconfig:"FB_TOKEN_URI"`
		AuthProviderX509CertUrl string `envconfig:"FB_AUTH_PROVIDER_X509_CERT_URL"`
		ClientX509CertUrl       string `envconfig:"FB_X509_CERT_URL"`
	}

	Limiter struct {
		RPS   int           `yaml:"rps"`
		Burst int           `yaml:"burst"`
		TTL   time.Duration `yaml:"ttl"`
	} `yaml:"limiter"`
}

func Init(configsDir string) (*Config, error) {
	var cfg Config

	readEnv(&cfg)
	readFile(&cfg, configsDir)

	return &cfg, nil
}

func processError(err error) {
	fmt.Println(err)
	os.Exit(2)
}

func readFile(cfg *Config, configsDir string) {
	f, err := os.Open(configsDir + "/main.yml")
	if err != nil {
		processError(err)
	}

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		processError(err)
	}
}

func readEnv(cfg *Config) {
	err := godotenv.Load(".env")
	if err != nil {
		//panic(err)
	}

	err = envconfig.Process("", cfg)
	if err != nil {
		processError(err)
	}
	log.Println("cfg")
	log.Println(cfg)
}
