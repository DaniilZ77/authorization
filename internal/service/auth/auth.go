package auth

import (
	"context"

	"github.com/DaniilZ77/authorization/internal/core"
	"github.com/DaniilZ77/authorization/internal/lib/jwt"
	"github.com/DaniilZ77/authorization/internal/lib/logger"
	"golang.org/x/crypto/bcrypt"
)

type service struct {
	userStorage core.UserStore
}

func New(userStorage core.UserStore) core.AuthService {
	return &service{
		userStorage: userStorage,
	}
}

func (s *service) Login(ctx context.Context, authConfig core.AuthConfig, user core.User) (*string, error) {
	userFromDB, err := s.userStorage.GetUserByUsername(ctx, user.Username)
	if err != nil {
		logger.Log().Error(ctx, err.Error())
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(userFromDB.PasswordHash), []byte(user.PasswordHash))
	if err != nil {
		logger.Log().Error(ctx, err.Error())
		return nil, core.ErrInvalidCredentials
	}

	token, err := jwt.GenerateToken(user.ID, authConfig)
	if err != nil {
		logger.Log().Error(ctx, err.Error())
		return nil, err
	}

	return token, nil
}

func (s *service) Signup(ctx context.Context, user core.User) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		logger.Log().Error(ctx, err.Error())
		return err
	}

	user.PasswordHash = string(passwordHash)

	_, err = s.userStorage.AddUser(ctx, user)
	if err != nil {
		return err
	}

	return nil
}
