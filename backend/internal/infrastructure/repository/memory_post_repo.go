
package repository

import (
	"app/internal/domain/aggregate"
	"context"
	"errors"
	"sync"
)

type InMemoryPostRepo struct {
	mu    sync.RWMutex
	posts map[string]*aggregate.Post
}

func NewInMemoryPostRepo() *InMemoryPostRepo {
	return &InMemoryPostRepo{posts: make(map[string]*aggregate.Post)}
}

func (r *InMemoryPostRepo) Save(ctx context.Context, post *aggregate.Post) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.posts[post.ID] = post
	return nil
}

func (r *InMemoryPostRepo) FindByID(ctx context.Context, id string) (*aggregate.Post, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	post, ok := r.posts[id]
	if !ok {
		return nil, errors.New("post not found")
	}
	return post, nil
}

func (r *InMemoryPostRepo) List(ctx context.Context) ([]*aggregate.Post, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	posts := make([]*aggregate.Post, 0, len(r.posts))
	for _, post := range r.posts {
		posts = append(posts, post)
	}
	return posts, nil
}
