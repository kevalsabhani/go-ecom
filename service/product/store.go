package product

import (
	"database/sql"

	"github.com/kevalsabhani/go-ecom/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) GetProducts() ([]*types.Product, error) {
	rows, err := s.db.Query("SELECT * FROM products")
	if err != nil {
		return nil, err
	}

	products := make([]*types.Product, 0)
	for rows.Next() {
		product, err := scanRowIntoProduct(rows)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (s *Store) CreateProduct(product *types.Product) error {
	_, err := s.db.Exec(
		"INSERT INTO products (name, description, price) VALUES (?, ?, ?)",
		product.Name,
		product.Description,
		product.Price,
	)
	return err
}

func scanRowIntoProduct(rows *sql.Rows) (*types.Product, error) {
	product := new(types.Product)

	if err := rows.Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.Price,
		&product.CreatedAt,
	); err != nil {
		return nil, err
	}
	return product, nil
}
