package main

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
)

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	SSLMode  string
}

func (cfg PostgresConfig) String() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Database, cfg.SSLMode)
}

func main() {
	// read from a file in the future
	cfg := PostgresConfig{
		Host:     "localhost",
		Port:     "5432",
		User:     "baloo",
		Password: "junglebook",
		Database: "pieratebay",
		SSLMode:  "disable",
	}

	db, err := sql.Open("pgx", cfg.String())
	if err != nil {
		panic(err)
	}

	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("connected")

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
		    id SERIAL PRIMARY KEY,
		    first_name TEXT,
		    email TEXT UNIQUE NOT NULL
		);

		CREATE TABLE IF NOT EXISTS orders (
		    id SERIAL PRIMARY KEY,
		    user_id INT NOT NULL,
		    amount INT,
		    description TEXT
		);
	`)
	if err != nil {
		panic(err)
	}
	fmt.Println("tables created")

	name := "bobo"
	email := "smarty@gmail.com"
	_, err = db.Exec(`
		INSERT INTO users (first_name, email)
		VALUES ($1, $2);`, name, email)
	if err != nil {
		panic(err)
	}
	fmt.Println("user created")
}
