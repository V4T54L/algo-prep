package repository

import (
	"app/internal/domain/aggregate"
	"context"
)

type ProblemRepository interface {
	List(ctx context.Context) ([]*aggregate.ProblemInfo, error)
}
