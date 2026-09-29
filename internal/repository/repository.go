package repository

import (
	"context"

	"github.com/NeoFSociety/cleanArch/internal/domain"
)

type UserRepository interface {
	Save(ctx context.Context, user *domain.User) error
	GetByUID(ctx context.Context, uid string) (*domain.User, error)
	Delete(ctx context.Context, uid string) error
}
