package storage

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sergssnorth27/life_manager/internal/models"
)

func GetUser(pool *pgxpool.Pool, telegramID int64) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user models.User
	err := pool.QueryRow(ctx,
		`SELECT id, telegram_id, username, first_name, last_name, created_at FROM users WHERE telegram_id = $1`,
		telegramID).Scan(&user.ID, &user.TelegramID, &user.Username, &user.FirstName, &user.LastName, &user.CreatedAt)

	if err == nil {
		return &user, nil
	}
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, err
	}
	if !errors.Is(err, pgx.ErrNoRows) {
		return nil, err
	}
	return &user, nil
}

func CreateUser(pool *pgxpool.Pool, telegramID int64, username string, firstName string, lastName string) (idUser int, createdAt time.Time, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var user models.User
	err = pool.QueryRow(ctx,
		`INSERT INTO users (telegram_id, username, first_name, last_name, created_at) VALUES ($1, $2, $3, $4, NOW()) RETURNING id, created_at`,
		telegramID, username, firstName, lastName).Scan(&user.ID, &user.CreatedAt)
	if err != nil {
		return 0, time.Time{}, err
	}
	return user.ID, user.CreatedAt, nil
}
