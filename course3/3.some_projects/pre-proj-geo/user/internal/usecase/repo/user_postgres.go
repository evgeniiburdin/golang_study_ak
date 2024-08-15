package repo

import (
	"context"
	"fmt"

	"user-service/internal/entity"
	"user-service/pkg/postgres"
)

const _defaultEntityCap = 64

// UserRepo -.
type UserRepo struct {
	*postgres.Postgres
}

// New -.
func New(pg *postgres.Postgres) *UserRepo {
	return &UserRepo{pg}
}

// Write -.
func (r *UserRepo) Write(ctx context.Context, u entity.User) error {
	sql, args, err := r.Builder.
		Insert("users").
		Columns("username", "email", "password", "user_role").
		Values(u.Username, u.Email, u.Password, u.Role).
		ToSql()
	if err != nil {
		return fmt.Errorf("UserRepo - Write - r.Builder: %w", err)
	}

	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("UserRepo - Write - r.Pool.Exec: %w", err)
	}

	return nil
}

// Get -.
func (r *UserRepo) Get(ctx context.Context, email string) (entity.User, error) {
	sql, args, err := r.Builder.
		Select("username, email, password, user_role").
		From("users").
		Where("email = ?", email).
		ToSql()
	if err != nil {
		return entity.User{}, fmt.Errorf("UserRepo - Get - r.Builder: %w", err)
	}

	var u entity.User
	err = r.Pool.QueryRow(ctx, sql, args...).Scan(&u.Username, &u.Email, &u.Password, &u.Role)
	if err != nil {
		return entity.User{}, fmt.Errorf("UserRepo - Get - r.Pool.QueryRow: %w", err)
	}

	return u, nil
}

// GetAll -.
func (r *UserRepo) GetAll(ctx context.Context) ([]entity.User, error) {
	sql, _, err := r.Builder.
		Select("username, email, user_role").
		From("users").
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("UserRepo - GetAll - r.Builder: %w", err)
	}

	rows, err := r.Pool.Query(ctx, sql)
	if err != nil {
		return nil, fmt.Errorf("UserRepo - GetAll - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	entities := make([]entity.User, 0, _defaultEntityCap)

	for rows.Next() {
		e := entity.User{}

		err = rows.Scan(&e.Username, &e.Email, &e.Role)
		if err != nil {
			return nil, fmt.Errorf("UserRepo - GetAll - rows.Scan: %w", err)
		}

		entities = append(entities, e)
	}

	return entities, nil
}

// Update -.
func (r *UserRepo) Update(ctx context.Context, u entity.User) error {
	sql, args, err := r.Builder.
		Update("users").
		Set("username", u.Username).
		Set("password", u.Password).
		Set("user_role", u.Role).
		Where("email = ?", u.Email).
		ToSql()
	if err != nil {
		return fmt.Errorf("UserRepo - Update - r.Builder: %w", err)
	}

	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("UserRepo - Update - r.Pool.QueryRow: %w", err)
	}

	return nil
}

// Delete -.
func (r *UserRepo) Delete(ctx context.Context, email string) error {
	sql, args, err := r.Builder.
		Delete("users").
		Where("email = ?", email).
		ToSql()
	if err != nil {
		return fmt.Errorf("UserRepo - Delete - r.Builder: %w", err)
	}

	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("UserRepo - Delete - r.Pool.Exec: %w", err)
	}

	return nil
}
