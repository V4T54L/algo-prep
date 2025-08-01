package usecase

import (
	"app/internal/domain/aggregate"
	"app/internal/domain/service"
	"context"
)

type ProblemUseCase struct {
	svc *service.ProblemDomainService
}

func NewProblemUseCase(svc *service.ProblemDomainService) *ProblemUseCase {
	return &ProblemUseCase{svc: svc}
}

func (u *ProblemUseCase) ListProblems(ctx context.Context) ([]*aggregate.ProblemInfo, error) {
	return u.svc.ListProblems(ctx)
}
