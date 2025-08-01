package service

import (
	"app/internal/domain/aggregate"
	"app/internal/domain/repository"
	"context"
)

type ProblemDomainService struct {
	repo repository.ProblemRepository
}

func NewProblemDomainService(repo repository.ProblemRepository) *ProblemDomainService {
	return &ProblemDomainService{repo: repo}
}

func (s *ProblemDomainService) ListProblems(ctx context.Context) ([]*aggregate.ProblemInfo, error) {
	return s.repo.List(ctx)
}
