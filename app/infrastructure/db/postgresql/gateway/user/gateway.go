package user

import (
	"database/sql"
)

type Gateway interface {
	IncreaseBalance(user_id string, balanc_id string, amount int) error
	DecreaseBalance(user_id string, balanc_id string, amount int) error
}

type gateway struct {
	db *sql.DB
}

func NewGateway(db *sql.DB) Gateway {
	return &gateway{db: db}
}