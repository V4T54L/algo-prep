package usecase

import (
	"app/internal/domain/aggregate"
	"app/internal/domain/service"
	"context"
)

type ContestUseCase struct {
	svc *service.ContestDomainService
}

func NewContestUseCase(svc *service.ContestDomainService) *ContestUseCase {
	return  &ContestUseCase{svc: svc}
}

func (u *ContestUseCase) ListContests(ctx context.Context)([]*aggregate.ContestInfo, error){
	return  u.svc.ListContests(ctx)
}