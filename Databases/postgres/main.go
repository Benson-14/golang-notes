package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var db *sql.DB

type Product struct {
	ID    int
	Name  string
	Price float64
}

func GetProductByID(db *sql.DB, id int) (*Product, error) {
	var p Product
	err := db.QueryRow("SELECT id, name, price FROM products WHERE id = $1", id).Scan(&p.ID, &p.Name, &p.Price)
	if err == sql.ErrNoRows {
		return nil, sql.ErrNoRows
	}
	if err != nil {
		return nil, err
	}

	return &p, nil
}

func QueryProducts(db *sql.DB) ([]Product, error) {
	rows, err := db.Query("SELECT name, price FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var p Product
		if err = rows.Scan(&p.Name, &p.Price); err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, rows.Err()
}

func SetupProducts(db *sql.DB) error {
	_, err := db.Exec(`DROP TABLE IF EXISTS products`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS products(
		id BIGSERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		price NUMERIC(10,2) NOT NULL
	)`)
	if err != nil {
		log.Fatal(err)
	}

	products := []struct {
		name  string
		price float64
	}{
		{"Laptop", 999.99},
		{"Mouse", 29.99},
	}

	var rowAffected int64
	for _, p := range products {
		res, err := db.Exec("INSERT INTO products (name, price) VALUES ($1, $2)", p.name, p.price)
		if err != nil {
			return fmt.Errorf("failed to insert %s: %w", p.name, err)
		}
		affected, _ := res.RowsAffected()
		rowAffected += affected
	}

	fmt.Println("Total rows affected: ", rowAffected)

	return nil
}

func main() {
	var err error
	connStr := "host=localhost port=5433 user=postgres password=postgres dbname=postgres sslmode=disable"
	db, err = sql.Open("pgx", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("Cannot connect:", err)
	} else {
		fmt.Println("Pinged successfully")
	}

	fmt.Println("Setting up products table")
	SetupProducts(db)

	fmt.Println("Querying entire table...")
	products, _ := QueryProducts(db)
	for _, p := range products {
		fmt.Println(p.Name, p.Price)
	}

	fmt.Println("Querying id: 1 data")

	p, _ := GetProductByID(db, 1)
	if p != nil {
		println(p.Name)
	}

}
