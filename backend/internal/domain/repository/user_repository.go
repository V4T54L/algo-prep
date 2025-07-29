
package repository

import (
	"app/internal/domain/aggregate"
	"context"
)

type UserRepository interface {
	Save(ctx context.Context, user *aggregate.User) error
	FindByID(ctx context.Context, id string) (*aggregate.User, error)
	FindByEmail(ctx context.Context, email string) (*aggregate.User, error)
}
