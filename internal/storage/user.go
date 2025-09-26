package storage

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sergssnorth27/life_manager/internal/models"
)

func GetOrCreateUser(pool *pgxpool.Pool, telegramID int64) (*models.User, error)
