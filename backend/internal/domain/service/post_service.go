
package service

import (
	"app/internal/domain/aggregate"
	"app/internal/domain/repository"
	"context"
	"github.com/google/uuid"
)

type PostDomainService struct {
	repo repository.PostRepository
}

func NewPostDomainService(repo repository.PostRepository) *PostDomainService {
	return &PostDomainService{repo: repo}
}

func (s *PostDomainService) CreatePost(ctx context.Context, title, content, authorID string) (*aggregate.Post, error) {
	post := aggregate.NewPost(uuid.New().String(), title, content, authorID)
	return post, s.repo.Save(ctx, post)
}

func (s *PostDomainService) ListPosts(ctx context.Context) ([]*aggregate.Post, error) {
	return s.repo.List(ctx)
}
