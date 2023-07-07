package main

import "time"

// 1. Education
// 2. Experience
// 3. Skills
// 4. Languages
// 5. Projects
// 6. Contacts
type Person struct {
	ID        int
	FirstName string        `json:"firstName"`
	LastName  string        `json:"lastName"`
	EduArr    []*Education  `json:"eduArr"`
	ExpArr    []*Experience `json:"expArr"`
}

func NewPerson(fn, ln string) (*Person, error) {
	return &Person{
		FirstName: fn,
		LastName:  ln,
		EduArr:    []*Education{},
		ExpArr:    []*Experience{},
	}, nil
}

type Education struct {
	ID          int
	School      string    `json:"school"`
	Degree      string    `json:"degree"`
	Field       string    `json:"field"`
	DateStarted time.Time `json:"dateStarted"`
	DateEnded   time.Time `json:"dateEnded"`
}

func NewEducation(school, degree, field string, dateStarted, dateEnded time.Time) (*Education, error) {
	return &Education{
		School:      string(school),
		Degree:      string(degree),
		Field:       string(field),
		DateStarted: time.Time(dateStarted),
		DateEnded:   time.Time(dateEnded),
	}, nil
}

type Experience struct {
	ID          int
	Company     string    `json:"company"`
	Role        string    `json:"role"`
	DateStarted time.Time `json:"dateStarted"`
	DateEnded   time.Time `json:"dateEnded"`
}

func NewExperience(company, role string, dateStarted, dateEnded time.Time) (*Experience, error) {
	return &Experience{
		Company:     string(company),
		Role:        string(role),
		DateStarted: time.Time(dateStarted),
		DateEnded:   time.Time(dateEnded),
	}, nil
}

type Skills struct {
	Technology  string `json:"technology"`
	Description string `json:"description"`
}

type Languages struct {
	Language    string `json:"language"`
	Level       string `json:"level"`
	Description string `json:"description"`
}

type Projects struct {
	ProjectName    string `json:"projectName"`
	TechnologyUsed string `json:"technologyUsed"`
	Description    string `json:"description"`
}

type Contact struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Age       int    `json:"age"`
	Email     string `json:"email"`
}
