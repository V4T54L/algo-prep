
package service

import (
	"app/internal/domain/aggregate"
	"app/internal/domain/repository"
	"context"
	"errors"
	"github.com/google/uuid"
)

type UserDomainService struct {
	repo repository.UserRepository
}

func NewUserDomainService(repo repository.UserRepository) *UserDomainService {
	return &UserDomainService{repo: repo}
}

func (s *UserDomainService) CreateUser(ctx context.Context, email, username, password string) (*aggregate.User, error) {
	if existing, _ := s.repo.FindByEmail(ctx, email); existing != nil {
		return nil, errors.New("email already exists")
	}
	user := aggregate.NewUser(uuid.New().String(), email, username, password)
	return user, s.repo.Save(ctx, user)
}

func (s *UserDomainService) GetUserByID(ctx context.Context, id string) (*aggregate.User, error) {
	return s.repo.FindByID(ctx, id)
}
