package repository

import (
	"database/sql"

	"github.com/xfpng345/ddd_learn/domain"
)

type UserRepository interface {
	Insert(DB *sql.DB, userID, name, email string) error
	GetByUserID(DB *sql.DB, userID string) (*domain.User, error)
}
