package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateEducation(*Education) error
	DeleteEducation(int) error
	UpdateEducation(*Education) error
	GetEducationByID(int) (*Education, error)
	GetEducation() ([]*Education, error)
}

type PostgresStore struct {
	db *sql.DB
}

/// CRUD FUNCTIONS - EDUCATION DB///
func (s *PostgresStore) CreateEducation(edu *Education) error {
	// Connect with db
	// Insert a row
	// return
	query := `insert into education (school,degree,field,date_started,date_ended) values ($1, $2, $3, $4, $5)`

	response, err := s.db.Query(
		query,
		edu.School,
		edu.Degree,
		edu.Field,
		edu.DateStarted, edu.DateEnded)

	if err != nil {
		fmt.Printf("%v\n", response.Err())
		return err
	}

	return err
}

func (s *PostgresStore) GetEducation() ([]*Education, error) {
	query := `select * from education`

	eduArr := []*Education{}

	response, err := s.db.Query(query)
	if err != nil {
		fmt.Println("query err")
		return nil, err
	}

	for response.Next() {
		edu := new(Education)

		err := response.Scan(&edu.ID,
			&edu.School,
			&edu.Degree,
			&edu.Field,
			&edu.DateStarted,
			&edu.DateEnded,
		)

		if err != nil {
			fmt.Printf("scan err")
			return nil, err
		}

		eduArr = append(eduArr, edu)
	}

	return eduArr, nil
}

func (s *PostgresStore) DeleteEducation(int) error                { return nil }
func (s *PostgresStore) UpdateEducation(*Education) error         { return nil }
func (s *PostgresStore) GetEducationByID(int) (*Education, error) { return &Education{}, nil }

// DATABASE INIT //
func NewPostgresStore() (*PostgresStore, error) {
	// 1. connect with DB using connection string
	// 2. ping db
	// 3. return DB object
	connStr := "user=postgres dbname=postgres password=cvapi sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStore{
		db: db,
	}, nil
}

func (s *PostgresStore) Init() error {
	return s.CreateEducationTable()
}

func (s *PostgresStore) CreateEducationTable() error {
	query := `create table if not exists education(
		id serial primary key,
		school varchar(50),
		degree varchar(50),
		field varchar(50),
		date_started timestamp,
		date_ended timestamp
	)`

	_, err := s.db.Exec(query)

	return err
}
