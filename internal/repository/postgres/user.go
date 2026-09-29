package postgres

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/NeoFSociety/cleanArch/internal/domain"
	"github.com/NeoFSociety/cleanArch/internal/repository"
)

type UserRepo struct {
	pool *pgxpool.Pool
}

var _ repository.UserRepository = (*UserRepo)(nil)

func NewUserRepo(pool *pgxpool.Pool) *UserRepo {
	return &UserRepo{pool: pool}
}

func (r *UserRepo) Save(ctx context.Context, user *domain.User) error {
	_, err := r.pool.Exec(
		ctx,
		`INSERT INTO users (user_uuid, user_name) VALUES ($1, $2)`,
		user.UUID, user.Username,
	)
	return err
}

func (r *UserRepo) GetByUID(ctx context.Context, uid string) (*domain.User, error) {
	var u domain.User
	err := r.pool.QueryRow(
		ctx,
		`SELECT user_uuid, user_name FROM users WHERE user_uuid = $1`,
		uid,
	).Scan(&u.UUID, &u.Username)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.ErrNotFound
		}
		return nil, err
	}
	return &u, nil
}

func (r *UserRepo) Delete(ctx context.Context, uid string) error {
	tag, err := r.pool.Exec(
		ctx,
		`DELETE FROM users WHERE user_uuid = $1`,
		uid,
	)
	if err != nil {
		return err
	}

	if tag.RowsAffected() == 0 {
		return domain.ErrNotFound
	}
	return nil
}
