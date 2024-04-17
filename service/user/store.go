package user

import (
	"database/sql"
	"fmt"

	"github.com/0x-pankaj/ecom/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (h *Store) GetUserByEmail(email string) (*types.User, error) {
	rows, err := h.db.Query("SELECT * FROM users WHERE email = $1", email)
	if err != nil {
		return nil, err
	}

	u := new(types.User)

	for rows.Next() {
		u, err = ScanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}

		if u.ID == 0 {
			return nil, fmt.Errorf("user not found")
		}

	}

	return u, nil

}

func (h *Store) CreateUserWithEmail(user types.User) error {
	_, err := h.db.Exec("INSERT INTO users (firstname, lastname, email, password) VALUES ($1, $2, $3, $4)", user.FirstName, user.LastName, user.Email, user.Password)
	if err != nil {
		return fmt.Errorf("error while creating user")
	}
	return nil
}

func ScanRowIntoUser(rows *sql.Rows) (*types.User, error) {
	user := new(types.User)

	err := rows.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.CreateAt,
	)
	if err != nil {
		return nil, err
	}

	return user, nil

}
