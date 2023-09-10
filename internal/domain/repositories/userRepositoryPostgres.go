package repositories

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/luksrocha/house-system/internal/domain/entities"
)

type UserRepositoryPostgres struct {
	DB *sqlx.DB
}

func NewUserRepositoryPostgres(db *sqlx.DB) *UserRepositoryPostgres {
	return &UserRepositoryPostgres{DB: db}
}

func (u *UserRepositoryPostgres) Insert(user *entities.User) error {
	prepare, err := u.DB.Prepare("INSERT INTO users (id, first_name, last_name, email, hashed_password, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7)")

	if err != nil {
		return err
	}

	_, err = prepare.Exec(user.ID, user.FirstName, user.LastName, user.Email, user.HashedPassword, user.CreatedAt, user.UpdatedAt)

	if err != nil {
		return err
	}

	return nil
}

func (u *UserRepositoryPostgres) Delete(id string) error {
	prepare, err := u.DB.Prepare("DELETE FROM users WHERE id = $1")

	if err != nil {
		return err
	}

	_, err = prepare.Exec(id)

	if err != nil {
		return err
	}

	return nil
}

func (u *UserRepositoryPostgres) Find(id string) (*entities.User, error) {
	prepare, err := u.DB.Prepare("SELECT * FROM users WHERE id = $1")

	if err != nil {
		return nil, err
	}

	var user entities.User

	err = prepare.QueryRow(id).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.HashedPassword, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &user, nil

}

func (u *UserRepositoryPostgres) FindByEmail(email string) (*entities.User, error) {
	prepare, err := u.DB.Prepare("SELECT * FROM users WHERE email = $1")

	if err != nil {
		return nil, err
	}

	var user entities.User

	err = prepare.QueryRow(email).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.HashedPassword, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &user, nil

}

func (u *UserRepositoryPostgres) Update(user *entities.User) (*entities.User, error) {
	prepare, err := u.DB.Prepare("UPDATE users SET name = $1, address = $2, updated_at = $3 WHERE id = $4")

	if err != nil {
		return nil, err
	}

	userUpdated, err := prepare.Exec(user.FirstName, user.LastName, time.Now(), user.ID)

	fmt.Println(userUpdated)

	if err != nil {
		return nil, err
	}

	return user, nil
}
