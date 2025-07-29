package usecase

import (
	"app/internal/domain/aggregate"
	"app/internal/domain/service"
	"context"
)

type PostUseCase interface {
	CreatePost(ctx context.Context, title, content, authorID string) (*aggregate.Post, error)
	ListPosts(ctx context.Context) ([]*aggregate.Post, error)
}

type postUseCase struct {
	svc *service.PostDomainService
}

func NewPostUseCase(svc *service.PostDomainService) PostUseCase {
	return &postUseCase{svc: svc}
}

func (u *postUseCase) CreatePost(ctx context.Context, title, content, authorID string) (*aggregate.Post, error) {
	return u.svc.CreatePost(ctx, title, content, authorID)
}

func (u *postUseCase) ListPosts(ctx context.Context) ([]*aggregate.Post, error) {
	return u.svc.ListPosts(ctx)
}
