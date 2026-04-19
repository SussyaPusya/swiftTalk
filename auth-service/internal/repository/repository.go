package repository

import (
	"context"

	"github.com/SussyaPusya/swiftTalk/auth-service/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"

	sq "github.com/Masterminds/squirrel"
)

type repository struct {
	pg *pgxpool.Pool
}

func NewRepository(pg *pgxpool.Pool) *repository {
	return &repository{pg: pg}
}

func (r *repository) CreateUser(ctx context.Context, user *models.CreateUserDTO) error {
	query := sq.Insert("users").
		Columns("username", "name", "password_hash").
		Values(user.Username, user.Name, user.Password).
		PlaceholderFormat(sq.Dollar)
	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}

	_, err = r.pg.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) GetUserByID(ctx context.Context, userID string) (*models.GetUserDTO, error) {
	query := sq.Select("id", "username", "name", "created_at").
		From("users").
		Where(sq.Eq{"id": userID}).
		PlaceholderFormat(sq.Dollar)
	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	var user models.GetUserDTO
	err = r.pg.QueryRow(ctx, sql, args...).Scan(&user.ID, &user.Username, &user.Name, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *repository) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	query := sq.Select("id", "username", "name", "password_hash", "created_at").
		From("users").
		Where(sq.Eq{"username": username}).
		PlaceholderFormat(sq.Dollar)
	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	var user models.User
	err = r.pg.QueryRow(ctx, sql, args...).Scan(&user.ID, &user.Username, &user.Name, &user.Password, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *repository) ChangePassword(ctx context.Context, userID string, newPassword string) error {
	query := sq.Update("users").
		Set("password", newPassword).
		Where(sq.Eq{"id": userID}).
		PlaceholderFormat(sq.Dollar)
	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}

	_, err = r.pg.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) ChangeName(ctx context.Context, userID string, newName string) error {
	query := sq.Update("users").
		Set("name", newName).
		Where(sq.Eq{"id": userID}).
		PlaceholderFormat(sq.Dollar)
	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	_, err = r.pg.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}
	return nil
}
