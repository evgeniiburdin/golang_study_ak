package repo

import (
	"context"
	"fmt"

	"auth-service/pkg/postgres"
)

const _defaultEntityCap = 64

// AuthRepo -.
type AuthRepo struct {
	*postgres.Postgres
}

// New -.
func New(pg *postgres.Postgres) *AuthRepo {
	return &AuthRepo{pg}
}

// WriteRefreshToken -.
func (r *AuthRepo) WriteRefreshToken(ctx context.Context, email, refreshToken string) error {
	sql, args, err := r.Builder.
		Insert("auth").
		Columns("email", "refreshToken").
		Values(email, refreshToken).
		ToSql()
	if err != nil {
		return fmt.Errorf("AuthRepo - WriteRefreshToken - r.Builder: %w", err)
	}

	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("AuthRepo - WriteRefreshToken - r.Pool.Exec: %w", err)
	}

	return nil
}

// ValidateRefreshToken -.
func (r *AuthRepo) ValidateRefreshToken(ctx context.Context, refreshToken string) error {
	sql, args, err := r.Builder.
		Select("email").
		From("auth").
		Where("refreshToken = ?", refreshToken).
		ToSql()
	if err != nil {
		return fmt.Errorf("AuthRepo - ValidateRefreshToken - r.Builder: %w", err)
	}

	_, err = r.Pool.Query(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("AuthRepo - ValidateRefreshToken - r.Pool.QueryRow: %w", err)
	}

	return nil
}

// UpdateRefreshToken -.
func (r *AuthRepo) UpdateRefreshToken(ctx context.Context, email, refreshToken string) error {
	sql, args, err := r.Builder.
		Update("auth").
		Set("refreshToken", refreshToken).
		Where("email = ?", email).
		ToSql()
	if err != nil {
		return fmt.Errorf("AuthRepo - UpdateRefreshToken - r.Builder: %w", err)
	}

	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("AuthRepo - UpdateRefreshToken - r.Pool.QueryRow: %w", err)
	}

	return nil
}

// RemoveRefreshToken -.
func (r *AuthRepo) RemoveRefreshToken(ctx context.Context, email string) error {
	sql, args, err := r.Builder.
		Update("auth").
		Set("refreshToken", "").
		Where("email = ?", email).
		ToSql()
	if err != nil {
		return fmt.Errorf("AuthRepo - RemoveRefreshToken - r.Builder: %w", err)
	}

	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("AuthRepo - RemoveRefreshToken - r.Pool.Exec: %w", err)
	}

	return nil
}
