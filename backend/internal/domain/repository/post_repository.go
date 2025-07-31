package repository

import (
	"app/internal/domain/aggregate"
	"context"
)

type PostRepository interface {
	Save(ctx context.Context, post *aggregate.PostDetail) error
	FindByID(ctx context.Context, id int) (*aggregate.PostDetail, error)
	List(ctx context.Context) ([]*aggregate.PostInfo, error)
}
