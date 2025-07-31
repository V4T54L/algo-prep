package service

import (
	"app/internal/domain/aggregate"
	"app/internal/domain/repository"
	"context"
)

type PostDomainService struct {
	repo repository.PostRepository
}

func NewPostDomainService(repo repository.PostRepository) *PostDomainService {
	return &PostDomainService{repo: repo}
}

func (s *PostDomainService) CreatePost(ctx context.Context, title, content, authorID string) (*aggregate.PostDetail, error) {
	post := aggregate.NewPost("https://picsum.photos/600/300?random=1", title, content, []string{"tag 1", "tag 2", "tag 3"})
	return post, s.repo.Save(ctx, post)
}

func (s *PostDomainService) ListPosts(ctx context.Context) ([]*aggregate.PostInfo, error) {
	return s.repo.List(ctx)
}
