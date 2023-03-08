package database

import (
	"aula1/internal/entity"
	"database/sql"
)

type OrderRepository struct {
	Db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{
		Db: db,
	}
}

func (r *OrderRepository) Save(order *entity.Order) error {
	_, err := r.Db.Exec("INSERT INTO ORDERS (id, price, tax, final_price) VALUES(?, ?, ?, ?)",
		order.ID, order.Price, order.Tax, order.FinalPrice)

	if err != nil {
		return err
	}
	return nil
}
