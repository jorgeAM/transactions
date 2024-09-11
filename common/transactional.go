package common

import "context"

type TxKey string

type Transactional interface {
	WithinTransaction(ctx context.Context, fn func(ctx context.Context) error) error
}
