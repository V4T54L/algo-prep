package repository

import (
	"app/internal/domain/aggregate"
	"context"
)

type ContestRepository interface {
	List(ctx context.Context) ([]*aggregate.ContestInfo, error)
}
