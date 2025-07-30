package repository

import (
	"app/internal/domain/aggregate"
	"app/internal/domain/repository"
	"context"
	"errors"
	"sync"
)

type InMemoryUserRepo struct {
	mu    sync.RWMutex
	users map[string]*aggregate.User
}

func NewInMemoryUserRepo() repository.UserRepository {
	return &InMemoryUserRepo{users: make(map[string]*aggregate.User)}
}

func (r *InMemoryUserRepo) Save(ctx context.Context, user *aggregate.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.users[user.ID] = user
	return nil
}

func (r *InMemoryUserRepo) FindByID(ctx context.Context, id string) (*aggregate.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	user, ok := r.users[id]
	if !ok {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (r *InMemoryUserRepo) FindByEmail(ctx context.Context, email string) (*aggregate.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	for _, user := range r.users {
		if user.Email == email {
			return user, nil
		}
	}
	return nil, errors.New("user not found")
}
