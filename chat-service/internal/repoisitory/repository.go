package repository

import (
	"context"

	"github.com/SussyaPusya/swiftTalk/chat-service/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"

	sq "github.com/Masterminds/squirrel"
)

type repository struct {
	pg *pgxpool.Pool
}

func NewRepository(pg *pgxpool.Pool) *repository {
	return &repository{pg: pg}
}

func (r *repository) CreateChat(ctx context.Context, data *models.CreateChatDTO) error {
	query := sq.Insert("chats").
		Columns("name", "type").
		Values(data.Name, data.Type).
		PlaceholderFormat(sq.Dollar)

	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}

	_, err = r.pg.Exec(ctx, sql, args...)
	return err
}

func (r *repository) AddMemberToChat(ctx context.Context, chatID, userID int) error {
	query := sq.Insert("chat_members").
		Columns("chat_id", "user_id").
		Values(chatID, userID).
		PlaceholderFormat(sq.Dollar)

	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}

	_, err = r.pg.Exec(ctx, sql, args...)
	return err
}

func (r *repository) GetChatMembers(ctx context.Context, chatID string) ([]models.Member, error) {
	query := sq.Select("user_id", "joined_at").
		From("chat_members").
		Where(sq.Eq{"chat_id": chatID}).
		PlaceholderFormat(sq.Dollar)

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.pg.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var members []models.Member
	for rows.Next() {
		var m models.Member
		if err := rows.Scan(&m.UserID, &m.JoinedAt); err != nil {
			return nil, err
		}
		members = append(members, m)
	}
	return members, nil
}

func (r *repository) CreateMessage(ctx context.Context, chatID string, data *models.SendMessageDTO) error {
	query := sq.Insert("messages").
		Columns("chat_id", "sender_id", "content").
		Values(chatID, data.SenderID, data.Content).
		PlaceholderFormat(sq.Dollar)

	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}

	_, err = r.pg.Exec(ctx, sql, args...)
	return err
}

func (r *repository) GetMessages(ctx context.Context, chatID string) ([]models.Message, error) {
	query := sq.Select("message_id", "sender_id", "content", "timestamp").
		From("messages").
		Where(sq.Eq{"chat_id": chatID}).
		PlaceholderFormat(sq.Dollar)

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.pg.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []models.Message
	for rows.Next() {
		var m models.Message
		if err := rows.Scan(&m.MessageID, &m.SenderID, &m.Content, &m.Timestamp); err != nil {
			return nil, err
		}
		messages = append(messages, m)
	}
	return messages, nil
}

func (r *repository) DeleteMessage(ctx context.Context, messageID string) error {
	query := sq.Delete("messages").
		Where(sq.Eq{"message_id": messageID}).
		PlaceholderFormat(sq.Dollar)

	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}

	_, err = r.pg.Exec(ctx, sql, args...)
	return err
}

func (r *repository) DeleteMemberFromChat(ctx context.Context, chatID, userID string) error {
	query := sq.Delete("chat_members").
		Where(sq.Eq{"chat_id": chatID, "user_id": userID}).
		PlaceholderFormat(sq.Dollar)

	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}

	_, err = r.pg.Exec(ctx, sql, args...)
	return err
}
