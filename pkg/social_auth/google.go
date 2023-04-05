package social_auth

import (
	"encoding/json"
	"io"
	"net/http"
)

const googleUserInfoURL = "https://www.googleapis.com/oauth2/v2/userinfo"

type GoogleUser struct {
	Email     string  `json:"email"`
	FirstName *string `json:"given_name"`
	LastName  *string `json:"family_name"`
	AvatarURL *string `json:"picture"`
	ID        string  `json:"id"`
}

type GoogleManager interface {
	GetUserInfo(token string) (*GoogleUser, error)
}

type Manager struct {
}

func NewManager() (*Manager, error) {
	return &Manager{}, nil
}

func (manager Manager) GetUserInfo(token string) (user *GoogleUser, err error) {
	response, err := http.Get(googleUserInfoURL + "?access_token=" + token)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bytes, &user)
	return user, err
}
