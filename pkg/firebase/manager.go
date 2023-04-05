package firebase

import (
	"context"
	"encoding/json"
	"firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
	"myapiproject/internal/config"
)

type FirebaseManager interface {
	GetUserInfo(uid string) (*auth.UserRecord, error)
}

type Manager struct {
	FirebaseApp *firebase.App
}

func NewManager(conf config.Config) (*Manager, error) {
	firebasebConfigJson, err := json.Marshal(&conf.FirebaseConfig)
	if err != nil {
		return nil, err
	}

	firebaseOption := option.WithCredentialsJSON(firebasebConfigJson)
	firebaseConfig := &firebase.Config{ProjectID: conf.FirebaseConfig.ProjectId}
	firebaseApp, err := firebase.NewApp(context.Background(), firebaseConfig, firebaseOption)
	if err != nil {
		return nil, err
	}
	return &Manager{
		FirebaseApp: firebaseApp,
	}, nil
}

func (manager Manager) GetUserInfo(uid string) (*auth.UserRecord, error) {
	client, err := manager.FirebaseApp.Auth(context.Background())
	if err != nil {
		return nil, err
	}
	return client.GetUser(context.Background(), uid)
}
