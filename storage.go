package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateEducation(*Education) error
	DeleteEducationByID(int) error
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

func (s *PostgresStore) DeleteEducationByID(id int) error {
	if _, err := s.db.Query(`delete from education where id = $1`, id); err != nil {
		return err
	}

	return nil
}
func (s *PostgresStore) UpdateEducation(*Education) error { return nil }
func (s *PostgresStore) GetEducationByID(id int) (*Education, error) {

	response, err := s.db.Query(`select * from education where id = $1`, id)

	if err != nil {
		fmt.Println("Query err")
		return nil, err
	}

	education := new(Education)

	for response.Next() {
		err := response.Scan(
			&education.ID,
			&education.School,
			&education.Degree,
			&education.Field,
			&education.DateStarted,
			&education.DateEnded)

		if err != nil {
			fmt.Println("scan err")
			return nil, err
		}
	}

	return education, nil
}

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
