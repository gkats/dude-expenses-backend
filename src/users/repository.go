package users

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (repository *Repository) CreateUser(params UserParams) (User, error) {
	user, err := NewUser(params)
	if err != nil {
		return user, err
	}

	query := `
	INSERT INTO users (email, encrypted_password)
	VALUES($1, $2)
	RETURNING id, created_at, updated_at
	`
	err = repository.db.QueryRow(query, user.Email, user.EncryptedPassword).Scan(
		&user.Id, &user.CreatedAt, &user.UpdatedAt,
	)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (repository *Repository) FindUserByEmail(email string) (*User, error) {
	user := User{}
	err := repository.db.QueryRow("SELECT * FROM users WHERE email = $1", email).Scan(
		&user.Id, &user.Email, &user.EncryptedPassword, &user.CreatedAt, &user.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &user, nil
}
