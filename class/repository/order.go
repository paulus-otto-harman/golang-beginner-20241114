package repository

import (
	"20241114/class/model"
	"database/sql"
)

type Order struct {
	Db *sql.DB
}

func InitOrderRepo(db *sql.DB) *Order {
	return &Order{db}
}

func (repo *Order) Create(order *model.Order) error {
	return nil
}
