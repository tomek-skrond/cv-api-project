package main

import "time"

// 1. Education
// 2. Experience
// 3. Skills
// 4. Languages
// 5. Projects
// 6. Contacts
// type Entry interface {
// 	ToString()
// }

type Person struct {
	ID        int
	FirstName string        `json:"firstName"`
	LastName  string        `json:"lastName"`
	EduArr    []*Education  `json:"edu_arr"`
	ExpArr    []*Experience `json:"exp_arr"`
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
	DateStarted time.Time `json:"date_started"`
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

// func (e Education) ToString() {
// 	print(e)
// }

type Experience struct {
	ID          int
	Company     string    `json:"company"`
	Role        string    `json:"role"`
	DateStarted time.Time `json:"date_started"`
	DateEnded   time.Time `json:"date_ended"`
}

func NewExperience(company, role string, dateStarted, dateEnded time.Time) (*Experience, error) {
	return &Experience{
		Company:     string(company),
		Role:        string(role),
		DateStarted: time.Time(dateStarted),
		DateEnded:   time.Time(dateEnded),
	}, nil
}

// func (e Experience) ToString() {
// 	print(e)
// }

type Language struct {
	ID          int
	Language    string `json:"language"`
	Level       string `json:"level"`
	Description string `json:"description"`
}

func NewLanguage(lang string, level string, desc string) (*Language, error) {
	return &Language{
		Language:    lang,
		Level:       level,
		Description: desc,
	}, nil
}

// func (l Language) ToString() {
// 	print(l)
// }

type Project struct {
	ID             int
	ProjectName    string `json:"project_name"`
	TechnologyUsed string `json:"technology_used"`
	Description    string `json:"description"`
}

func NewProject(name string, tech string, desc string) (*Project, error) {
	return &Project{
		ProjectName:    name,
		TechnologyUsed: tech,
		Description:    desc,
	}, nil
}

// func (p Project) ToString() {
// 	print(p)
// }

type Contact struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Age       int    `json:"age"`
	Email     string `json:"email"`
}

type Skills struct {
	ID          int
	Technology  string `json:"technology"`
	Description string `json:"description"`
}

func NewSkills(tech, desc string) (*Skills, error) {
	return &Skills{
		Technology:  string(tech),
		Description: string(desc),
	}, nil
}
