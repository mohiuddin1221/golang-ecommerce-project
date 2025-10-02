package repo

import "github.com/jmoiron/sqlx"

type User struct {
	ID          int    `json:"id" db:"id"`
	FirstName   string `json:"first_name" db:"first_name"`
	LastName    string `json:"last_name" db:"last_name"`
	Email       string `json:"email" db:"email"`
	Password    string `json:"password" db:"password"`
	IsShopOwner bool   `json:"is_shop_owner" db:"is_shop_owner"`
}
type UserRepo interface {
	Create(u User) (*User, error)
	FindUser(email, password string) (*User, error)
}
type userRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) UserRepo {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) Create(u User) (*User, error) {
	query := `
			INSERT INTO users (
			first_name, 
			last_name, 
			email, 
			password, 
			is_shop_owner
			)
			VALUES (
			:first_name, 
			:last_name, 
			:email, 
			:password, 
			:is_shop_owner
			)
			RETURNING id
		`

	stmt, err := r.db.PrepareNamed(query)
	if err != nil {
		return nil, err
	}

	var id int
	err = stmt.Get(&id, &u)
	if err != nil {
		return nil, err
	}

	u.ID = id
	return &u, nil

}

func (r *userRepo) FindUser(email, password string) (*User, error) {
	query := `
        SELECT id, first_name, last_name, email, password, is_shop_owner
        FROM users
        WHERE email = $1 AND password = $2
        LIMIT 1
    `

	var user User
	err := r.db.Get(&user, query, email, password)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
