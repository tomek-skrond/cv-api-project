package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateEducation(*Education) error
	DeleteEducationByID(int) error
	UpdateEducation(*Education) error
	GetEducationByID(int) (*Education, error)
	GetEducation() ([]*Education, error)
	CreateExperience(*Experience) error
	DeleteExperienceByID(int) error
	UpdateExperience(*Experience) error
	GetExperienceByID(int) (*Experience, error)
	GetExperience() ([]*Experience, error)
}

type PostgresStore struct {
	db *sql.DB
}

// CRUD FUNCTIONS - EDUCATION DB //
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

	check := `select * from education where id = $1`

	if !s.rowExists(check, id) {
		return apiError{Err: "not existing resource", Status: http.StatusNotFound}
	}

	if _, err := s.db.Query(`delete from education where id = $1`, id); err != nil {
		return err
	}

	return apiError{Err: "Deleted requested resource", Status: http.StatusOK}
}
func (s *PostgresStore) UpdateEducation(edu *Education) error {
	// TODO: CHECK IF RECORD EXISTS
	query := `update education set school = $1, degree = $2, field = $3, date_started = $4, date_ended = $5 where id = $6`

	if !s.rowExists(`select * from education where id = $1`, edu.ID) {
		return apiError{Err: "not existing resource", Status: http.StatusNotFound}
	}

	response, err := s.db.Query(query,
		edu.School,
		edu.Degree,
		edu.Field,
		edu.DateStarted,
		edu.DateEnded,
		edu.ID,
	)
	if err != nil {
		fmt.Printf("%v\n", response.Err())
		return err
	}

	return nil
}

func (s *PostgresStore) rowExists(query string, args ...interface{}) bool {
	var exists bool
	query = fmt.Sprintf("SELECT exists (%s)", query)
	err := s.db.QueryRow(query, args...).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		log.Fatalf("error")
	}
	return exists
}

func (s *PostgresStore) GetEducationByID(id int) (*Education, error) {

	query := fmt.Sprintf(`select * from education where id = $1`)

	if !s.rowExists(query, id) {
		return nil, apiError{Err: "not existing resource", Status: http.StatusNotFound}
	}

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

	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	database := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=disable", user, pass, host, database)

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
	if err := s.CreateEducationTable(); err != nil {
		return err
	}
	if err := s.CreateExperienceTable(); err != nil {
		return err
	}

	return nil

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

func (s *PostgresStore) CreateExperienceTable() error {
	query := `create table if not exists experience(
		id serial primary key,
		company varchar(50),
		role varchar(50),
		date_started timestamp,
		date_ended timestamp
	)`

	_, err := s.db.Exec(query)

	return err
}

func (s *PostgresStore) CreateExperience(exp *Experience) error {
	query := `insert into experience (company,role,date_started,date_ended) values ($1, $2, $3, $4)`

	response, err := s.db.Query(
		query,
		exp.Company,
		exp.Role,
		exp.DateStarted, exp.DateEnded)

	if err != nil {
		fmt.Printf("%v\n", response.Err())
		return err
	}

	return err
}

func (s *PostgresStore) DeleteExperienceByID(id int) error {

	if !s.rowExists(`select * from experience where id = $1`, id) {
		return fmt.Errorf("Permission denied!")
	}

	if _, err := s.db.Query(`delete from experience where id = $1`, id); err != nil {
		return err
	}
	return apiError{Err: "Resource deleted successfully", Status: http.StatusOK}
}

func (s *PostgresStore) UpdateExperience(exp *Experience) error {
	query := `update experience set company = $1, role = $2, date_started = $3, date_ended = $4 where id = $5`

	if !s.rowExists(`select * from experience where id = $1`, exp.ID) {
		return apiError{Err: "not existing resource", Status: http.StatusNotFound}
	}

	if _, err := s.db.Query(query,
		exp.ID,
		exp.Company,
		exp.Role,
		exp.DateStarted,
		exp.DateEnded,
	); err != nil {
		return apiError{Err: "query err", Status: http.StatusInternalServerError}
	}

	return nil
}
func (s *PostgresStore) GetExperienceByID(id int) (*Experience, error) {

	query := `select * from experience where id = $1`

	if !s.rowExists(`select * from experience where id = $1`, id) {
		return nil, apiError{Err: "not existing resource", Status: http.StatusNotFound}
	}

	response, err := s.db.Query(query)
	if err != nil {
		return nil, apiError{Err: "query error", Status: http.StatusInternalServerError}
	}

	exp := new(Experience)

	for response.Next() {

		err := response.Scan(
			&exp.ID,
			&exp.Company,
			&exp.Role,
			&exp.DateStarted,
			&exp.DateEnded,
		)

		if err != nil {
			fmt.Println("formating err GetExperienceByID")
			return nil, err
		}
	}

	return exp, nil
}
func (s *PostgresStore) GetExperience() ([]*Experience, error) {

	query := `select * from experience`

	expArr := []*Experience{}

	response, err := s.db.Query(query)

	if err != nil {
		fmt.Println("query err")
		return nil, err
	}

	for response.Next() {
		exp := new(Experience)

		err := response.Scan(&exp.ID,
			&exp.Company,
			&exp.Role,
			&exp.DateStarted,
			&exp.DateEnded,
		)

		if err != nil {
			fmt.Printf("scan err")
			return nil, err
		}

		expArr = append(expArr, exp)
	}

	return expArr, nil
}
