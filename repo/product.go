package repo

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type Product struct {
	ID          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
	Price       int    `json:"price" db:"price"`
}

type ProductRepo interface {
	Craete(p Product) (*Product, error)
	List() ([]*Product, error)
	Get(productId int) (*Product, error)
	Update(product Product) error
	Delete(productId int) error
}

type productRepo struct {
	db *sqlx.DB
}

// /constructor is a function
// Creating a constuctor for struct
func NewproductRepo(db *sqlx.DB) ProductRepo {
	//first of all constructor banaici
	return &productRepo{
		db: db,
	}
	//then Someproduct pass by this constuctor
}

func (r *productRepo) Craete(p Product) (*Product, error) {
	query := `
		INSERT INTO products(
			title,
			description,
			price
		)VALUES(
			$1,
			$2,
			$3
		)
		RETURNING id
	`
	row := r.db.QueryRow(query, p.Title, p.Description, p.Price)
	err := row.Scan(&p.ID)
	if err != nil {
		return nil, err
	}

	return &p, nil

}

func (r *productRepo) List() ([]*Product, error) {
	var prdList []*Product
	query := `
	SELECT *
	FROM products
	`
	err := r.db.Select(&prdList, query)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return prdList, nil

}

func (r *productRepo) Get(Id int) (*Product, error) {
	var prd Product
	query := `
		SELECT *
		FROM products
		WHERE id = $1
		
		`
	err := r.db.Get(&prd, query, Id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &prd, nil

}

func (r *productRepo) Update(p Product) error {
	query := `
		UPDATE users SET
			title = $1,
			description= $2,
			price = $3
		WHERE id = $5
	
	`
	row := r.db.QueryRow(query, p.Title, p.Description, p.Price)
	err := row.Err()
	if err != nil {
		return err
	}
	return nil

}

func (r *productRepo) Delete(id int) error {
	query := `
	DELETE FROM products WHERE id = $1
	
	`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil

}
