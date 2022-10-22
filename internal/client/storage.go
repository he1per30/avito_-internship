package client

import "context"

type Storage interface {
	moneyTransfer(ctx context.Context, client Client) (string, error)
}
