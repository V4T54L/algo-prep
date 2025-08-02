package service

import (
	"app/internal/domain/aggregate"
	"app/internal/domain/repository"
	"context"
)

type ContestDomainService struct {
	repo repository.ContestRepository
}

func NewContestService(repo repository.ContestRepository) *ContestDomainService {
	return &ContestDomainService{repo: repo}
}

func (s *ContestDomainService) ListContests(ctx context.Context) ([]*aggregate.ContestInfo, error) {
	return s.repo.List(ctx)
}
