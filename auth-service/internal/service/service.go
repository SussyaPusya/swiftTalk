package service

import (
	"context"
	"errors"

	"github.com/SussyaPusya/swiftTalk/auth-service/internal/models"
	"github.com/SussyaPusya/swiftTalk/auth-service/internal/pkg"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

type repository interface {
	CreateUser(ctx context.Context, user *models.CreateUserDTO) error
	GetUserByID(ctx context.Context, userID string) (*models.GetUserDTO, error)
	GetUserByUsername(ctx context.Context, username string) (*models.User, error)
	ChangePassword(ctx context.Context, userID string, newPassword string) error
	ChangeName(ctx context.Context, userID string, newName string) error
}

type service struct {
	r      repository
	logger *zap.Logger
}

func NewService(r repository, logger *zap.Logger) *service {
	return &service{r: r, logger: logger}
}

func (s *service) CreateUser(ctx context.Context, user *models.CreateUserDTO) error {

	hashedPassword, err := pkg.Encode(user.Password)
	if err != nil {
		s.logger.Error("failed to hash password", zap.Error(err))
		return err
	}
	user.Password = hashedPassword
	err = s.r.CreateUser(ctx, user)
	if err != nil {
		if err.Error() == pgerrcode.UniqueViolation {

			s.logger.Error("username already taken", zap.String("username", user.Username))
			return errors.New("username already taken")
		}
		s.logger.Error("failed to create user", zap.Error(err))
		return err
	}
	return nil
}
func (s *service) Login(ctx context.Context, login *models.LoginDTO) (*models.GetUserDTO, error) {
	s.logger.Info("login attempt", zap.String("username", login.Username))

	user, err := s.r.GetUserByUsername(ctx, login.Username)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			s.logger.Error("invalid user", zap.String("username", login.Username), zap.Error(err))
			return nil, errors.New("invalid user")
		}
		return nil, err
	}

	if !pkg.Check(login.Password, user.Password) {
		s.logger.Error("invalid password", zap.String("username", login.Username), zap.Error(err))
		return nil, errors.New("invalid password")
	}

	userDTO := &models.GetUserDTO{
		ID:        user.ID,
		Username:  user.Username,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
	}
	return userDTO, nil
}

func (s *service) GetUserByID(ctx context.Context, userID string) (*models.GetUserDTO, error) {
	user, err := s.r.GetUserByID(ctx, userID)
	if err != nil {
		s.logger.Error("failed to get user by id", zap.String("userID", userID), zap.Error(err))
	}
	return user, err
}

func (s *service) ChangePassword(ctx context.Context, userID string, newPassword string) error {

	err := s.r.ChangePassword(ctx, userID, newPassword)
	if err != nil {
		s.logger.Error("failed to change password", zap.String("userID", userID), zap.Error(err))
	}
	return err
}

func (s *service) ChangeName(ctx context.Context, userID string, newName string) error {
	return s.r.ChangeName(ctx, userID, newName)
}
