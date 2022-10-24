package userDb

import (
	"avito/internal/user"
	"avito/pkg/client/postgresql"
	"avito/pkg/logging"
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
)

type repository struct {
	client postgresql.Client
	logger *logging.Logger
	pool   *pgxpool.Pool
}

func (r *repository) RevenueRecognition(userId int, sum float64, serviceId int, orderId int) error {
	exists := r.checkClient(userId)
	if !exists {
		return errors.New("user not found")
	}
	var err error
	//var balance float64
	tx, err := r.pool.Begin(context.Background())
	defer func() {
		if err != nil {
			_ = tx.Rollback(context.Background())
		}
	}()

	_, err = r.client.Exec(context.Background(), `INSERT INTO accounting 
SELECT * FROM reserve WHERE client_id = $1 AND reserve_sum = $2 AND service_id = $3
AND  order_id = $4`, userId, sum, serviceId, orderId)
	if err != nil {
		return err
	}
	_, err = r.client.Exec(context.Background(), `DELETE FROM reserve WHERE client_id = $1
                AND reserve_sum = $2 AND service_id = $3 AND  order_id = $4`, userId, sum, serviceId, orderId)

	if err != nil {
		return err
	}
	return nil
}

func (r *repository) GetBalance(userId int) (float64, error) {
	exists := r.checkClient(userId)
	if !exists {
		return 0, errors.New("user not found")
	}
	var balance float64
	row := r.client.QueryRow(context.Background(), `SELECT balance FROM client WHERE id = $1`, userId)
	_ = row.Scan(&balance)

	return balance, nil
}

func (r *repository) ReserveAmount(userId int, sum float64, serviceId int, orderId int) error {
	exists := r.checkClient(userId)
	if !exists {
		return errors.New("user not found")
	}
	var err error
	var balance float64
	tx, err := r.pool.Begin(context.Background())
	defer func() {
		if err != nil {
			_ = tx.Rollback(context.Background())
		}
	}()

	row := r.client.QueryRow(context.Background(), `SELECT balance FROM client WHERE id = $1`, userId)
	_ = row.Scan(&balance)
	if balance < sum {
		return errors.New("not enough money")
	}

	_, err = r.client.Exec(context.Background(), "UPDATE client SET balance = balance - $1 WHERE id = $2", sum, userId)
	if err != nil {
		fmt.Println(err)
	}

	_, err = r.client.Exec(context.Background(), `INSERT INTO reserve (client_id, reserve_sum, service_id, order_id) 
VALUES ($1, $2,$3,$4)`, userId, sum, serviceId, orderId)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (r *repository) IncrementBalance(userId int, sum float64) error {
	exists := r.checkClient(userId)
	var err error
	if !exists {
		_, err = r.client.Exec(context.Background(), "INSERT INTO client (id,balance) VALUES ($1,$2)", userId, sum)
	} else {
		_, err = r.client.Exec(context.Background(), "UPDATE client SET balance = balance + $1 WHERE id = $2", sum, userId)
	}

	if err != nil {
		return err
	}

	return nil
}

func (r *repository) checkClient(userId int) bool {
	row := r.client.QueryRow(context.Background(), `SELECT TRUE FROM client WHERE id = $1`, userId)
	var exists bool
	_ = row.Scan(&exists)
	return exists
}

func NewRepository(client postgresql.Client, logger *logging.Logger, pool *pgxpool.Pool) user.Repository {
	return &repository{
		client: client,
		logger: logger,
		pool:   pool,
	}
}
