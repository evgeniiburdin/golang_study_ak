package repository

import (
	"database/sql"
	"fmt"
)

type PostgresUserRepository struct {
	db *sql.DB
}

func NewPostgresUserRepository(db *sql.DB) (UserRepository, error) {
	createTableSQL := `CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        username TEXT NOT NULL UNIQUE,
        password TEXT NOT NULL
    );`
	_, err := db.Exec(createTableSQL)
	if err != nil {
		return nil, err
	}
	return &PostgresUserRepository{
		db: db,
	}, nil
}

func (ur *PostgresUserRepository) CreateUser(username, password string) error {
	tx, err := ur.db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	createUserSQL := `INSERT INTO users (username, password) VALUES ($1, $2)`

	_, err = tx.Exec(createUserSQL, username, password)
	if err != nil {
		return fmt.Errorf("failed to execute query: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func (ur *PostgresUserRepository) ReadUser(username string) (password string, err error) {
	query := `SELECT password FROM users WHERE username = $1`

	err = ur.db.QueryRow(query, username).Scan(&password)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("user not found: %w", err)
		}
		return "", fmt.Errorf("query execution failed: %w", err)
	}

	return password, nil
}

func (ur *PostgresUserRepository) UpdateUser(username, newPassword string) error {
	updateUserSQL := `UPDATE users SET password = $1 WHERE username = $2`

	result, err := ur.db.Exec(updateUserSQL, newPassword, username)
	if err != nil {
		return fmt.Errorf("failed to execute update query: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("user not found")
	}

	return nil
}

func (ur *PostgresUserRepository) DeleteUser(username string) error {
	deleteUserSQL := `DELETE FROM users WHERE username = $1`

	result, err := ur.db.Exec(deleteUserSQL, username)
	if err != nil {
		return fmt.Errorf("failed to execute delete query: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("user not found")
	}

	return nil
}
