package service

import (
	"myapiproject/internal/models"
	"myapiproject/internal/repository"
	"myapiproject/pkg/auth"
	"strconv"
	"time"
)

type Users interface {
	Login(input LoginInput) (models.Session, error)
	RefreshToken(input RefreshInput) (models.Session, error)
	GetList(request user_dto.UserListDTO) ([]models.User, error)
}

type UserService struct {
	UserRepository         repository.UserRepository
	RefreshTokenRepository repository.RefreshTokenRepository
	TokenManager           auth.TokenManager
	RedisService           *RedisService
}

func NewUsersService(userRep repository.UserRepository, refreshRep repository.RefreshTokenRepository, m auth.TokenManager, redisService *RedisService) *UserService {
	return &UserService{
		UserRepository:         userRep,
		RefreshTokenRepository: refreshRep,
		TokenManager:           m,
		RedisService:           redisService,
	}
}

func (s *UserService) Login(input LoginInput) (models.Session, error) {
	user, err := s.UserRepository.Find(input.UserId)
	if err != nil {
		return models.Session{}, err
	}

	tokens, err := s.createTokens(user)
	if err != nil {
		return models.Session{}, err
	}

	return tokens, nil
}

func (s *UserService) RefreshToken(input RefreshInput) (models.Session, error) {
	refreshToken, err := s.RefreshTokenRepository.Find(input.Token)
	if err != nil {
		return models.Session{}, err
	}

	user, err := s.UserRepository.Find(refreshToken.UserID)
	if err != nil {
		return models.Session{}, err
	}

	tokens, err := s.createTokens(user)
	if err != nil {
		return models.Session{}, err
	}

	return tokens, nil
}

func (s *UserService) createTokens(user models.User) (token models.Session, err error) {
	// TODO: вынести ttl в сервис
	accessToken, _ := s.TokenManager.NewAccessToken(strconv.Itoa(int(user.ID)), 15*time.Minute)
	refreshToken, _ := s.TokenManager.NewRefreshToken()

	_, err = s.RefreshTokenRepository.Create(user.ID, refreshToken, time.Now().AddDate(0, 1, 0))

	if err != nil {
		return models.Session{}, err
	}
	return models.Session{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *UserService) GetList(request user_dto.UserListDTO) ([]models.User, error) {
	users, err := s.UserRepository.All()
	if err != nil {
		return nil, err
	}

	return users, nil
}
