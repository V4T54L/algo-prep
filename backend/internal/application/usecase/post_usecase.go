package usecase

import (
	"app/internal/domain/aggregate"
	"app/internal/domain/service"
	"context"
)

type PostUseCase struct {
	svc *service.PostDomainService
}

func NewPostUseCase(svc *service.PostDomainService) *PostUseCase {
	return &PostUseCase{svc: svc}
}

func (u *PostUseCase) CreatePost(ctx context.Context, title, content, authorID string) (*aggregate.PostDetail, error) {
	return u.svc.CreatePost(ctx, title, content, authorID)
}

func (u *PostUseCase) ListPosts(ctx context.Context) ([]*aggregate.PostInfo, error) {
	return u.svc.ListPosts(ctx)
}
