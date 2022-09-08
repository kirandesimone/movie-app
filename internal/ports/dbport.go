package ports

import "context"

type Db interface {
	Disconnect() error
	FindAll(ctx context.Context, results interface{}) error
}
