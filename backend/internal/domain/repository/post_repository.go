
package repository

import (
	"app/internal/domain/aggregate"
	"context"
)

type PostRepository interface {
	Save(ctx context.Context, post *aggregate.Post) error
	FindByID(ctx context.Context, id string) (*aggregate.Post, error)
	List(ctx context.Context) ([]*aggregate.Post, error)
}
