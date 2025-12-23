package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount(*Account) error
	DeletAccount(int) error
	UpdateAccount(*Account) error
	GetAccounts()([]*Account, error)
	GetAccountByID(int) (*Account, error)
}

type PostgresStore struct {
	db *sql.DB

}

func NewPostgresStore()(*PostgresStore, error){
	connStr := "user=postgres dbname=postgres password=gobank sslmode=disable"

	db, err := sql.Open("postgres", connStr)

	if err != nil{
		return nil, err
	}
	if err := db.Ping(); err != nil{
		return nil, err
	}

	return &PostgresStore{db: db},nil
}


func (s *PostgresStore) Init() error {
	return s.CreateAccountTable()
}

func (s *PostgresStore) CreateAccountTable() error {
	query := `CREATE TABLE IF NOT EXISTS account (
		id SERIAL PRIMARY KEY,
		first_name VARCHAR(50),
		last_name VARCHAR(50), 
		number SERIAL,
		balance SERIAL,
		created_at TIMESTAMP
	)`

	_, err := s.db.Exec(query)
	return err

}

func (s *PostgresStore) CreateAccount(acc *Account)error{
	query:=(`
	INSERT INTO account (first_name, last_name,number,balance,
	created_at) VALUES ($1, $2, $3, $4, $5)
	`)
	_, err := s.db.Exec(query, acc.FirstName, acc.LastName, acc.Number, acc.Balance,acc.CreatedAt)

	return err
}


func (s *PostgresStore) UpdateAccount(*Account) error {
	return nil
}

func (s *PostgresStore) DeletAccount(id int) error {
	return nil
}


func (s * PostgresStore)GetAccountByID(id int)(*Account, error) {
	return nil,nil
}

func (s *PostgresStore) GetAccounts()([]*Account, error){
	rows, err := s.db.Query("select * from account")
	if err != nil {
		return nil, err
	}

	accounts := []*Account{}
	for rows.Next(){
		account := new(Account)
		err := rows.Scan(&account.ID, 
			&account.FirstName, 
			&account.LastName,
			&account.Balance,
			&account.Number,
			&account.CreatedAt,
		)

		if err != nil{
			return nil, err
		}

		accounts = append(accounts, account)

		
	}

	return accounts, nil
}


// work -> better-> faster