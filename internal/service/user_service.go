package service

import (
	"context"

	"github.com/google/uuid"

	"github.com/NeoFSociety/cleanArch/internal/domain"
	"github.com/NeoFSociety/cleanArch/internal/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Register(ctx context.Context, username string) (*domain.User, error) {

	user := &domain.User{
		UUID:     uuid.New().String(),
		Username: username,
	}

	if err := s.repo.Save(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) GetByUID(ctx context.Context, uid string) (*domain.User, error) {
	return s.repo.GetByUID(ctx, uid)
}

func (s *UserService) Delete(ctx context.Context, uid string) error {
	return s.repo.Delete(ctx, uid)
}
