package main

///////////// TUTAJ API /////////////
import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type APIServer struct {
	ListenPort string
	Store      Storage
}

func NewAPIServer(port string, s Storage) *APIServer {
	return &APIServer{
		ListenPort: port,
		Store:      s,
	}
}

func (s *APIServer) Run() {
	//http.HandleFunc("/user", makeHTTPHandler(handleGetUserByID))
	router := mux.NewRouter()
	router.HandleFunc("/education", makeHTTPHandler(s.handleEducation))
	router.HandleFunc("/education/{id}", makeHTTPHandler(s.handleEducationByID))
	router.HandleFunc("/experience", makeHTTPHandler(s.handleExperience))

	fmt.Printf("Server listening at port %s \n", s.ListenPort)
	http.ListenAndServe(s.ListenPort, router)
}

// API - Experience//
func (s *APIServer) handleExperience(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGetExperience(w, r)
	}
	if r.Method == "POST" {
		return s.handleCreateExperience(w, r)
	}

	return apiError{Err: "Bad request", Status: http.StatusBadRequest}
}

func (s *APIServer) handleGetExperience(w http.ResponseWriter, r *http.Request) error {

	expArray, err := s.Store.GetExperience()
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, expArray)
}

func (s *APIServer) handleCreateExperience(w http.ResponseWriter, r *http.Request) error {
	// create empty object
	// fetch data from request (decode)
	// feed data to empty object
	// invoke database operation (s.Store.CreateEducation())
	// return json with data
	createExpReq := new(Experience)
	if err := json.NewDecoder(r.Body).Decode(createExpReq); err != nil {
		return err
	}

	exp, err := NewExperience(
		createExpReq.Company,
		createExpReq.Role,
		createExpReq.DateStarted,
		createExpReq.DateEnded,
	)

	if err != nil {
		fmt.Println("formatting error in Exp")
		return err
	}

	if err := s.Store.CreateExperience(exp); err != nil {
		return err
	}
	return WriteJSON(w, http.StatusOK, exp)
}

// API - Education//
func (s *APIServer) handleEducation(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGetEducation(w, r)
	}
	if r.Method == "POST" {
		return s.handleCreateEducation(w, r)
	}

	return apiError{Err: "Method not implemented", Status: http.StatusNotImplemented}
}

func (s *APIServer) handleGetEducation(w http.ResponseWriter, r *http.Request) error {
	// invoke db operation
	// encode db response to JSON + return
	eduArray, err := s.Store.GetEducation()
	if err != nil {
		fmt.Println("store err")
		return err
	}

	return WriteJSON(w, http.StatusOK, eduArray)
}

func (s *APIServer) handleEducationByID(w http.ResponseWriter, r *http.Request) error {
	id, err := GetID(w, r)
	if err != nil {
		return err
	}

	if r.Method == "GET" {

		edu, err := s.Store.GetEducationByID(id)
		if err != nil {
			return err
		}

		return WriteJSON(w, http.StatusOK, edu)
	}

	if r.Method == "DELETE" {
		return s.Store.DeleteEducationByID(id)
	}
	if r.Method == "PUT" {

		id, err := GetID(w, r)
		if err != nil {
			return err
		}
		eduUpdateRequest := new(Education)
		eduUpdateRequest.ID = id

		if err := json.NewDecoder(r.Body).Decode(eduUpdateRequest); err != nil {
			return err
		}

		updatedEdu := &Education{
			ID:          eduUpdateRequest.ID,
			School:      eduUpdateRequest.School,
			Degree:      eduUpdateRequest.Degree,
			Field:       eduUpdateRequest.Field,
			DateStarted: eduUpdateRequest.DateStarted,
			DateEnded:   eduUpdateRequest.DateEnded}

		if err != nil {
			return err
		}

		return s.Store.UpdateEducation(updatedEdu)
	}
	return apiError{Err: "error", Status: http.StatusBadRequest}
}
func (s *APIServer) handleCreateEducation(w http.ResponseWriter, r *http.Request) error {
	// create empty object
	// fetch data from request (decode)
	// feed data to empty object
	// invoke database operation (s.Store.CreateEducation())
	// return json with data
	createEducationRequest := new(Education)

	if err := json.NewDecoder(r.Body).Decode(createEducationRequest); err != nil {
		fmt.Println("formatting err")
		return err
	}

	newEducation, err := NewEducation(createEducationRequest.School,
		createEducationRequest.Degree,
		createEducationRequest.Field,
		createEducationRequest.DateStarted,
		createEducationRequest.DateEnded)

	if err != nil {
		fmt.Println("constructor err")
		return err
	}

	if err := s.Store.CreateEducation(newEducation); err != nil {
		fmt.Println("store err")
		return err
	}

	return WriteJSON(w, http.StatusOK, newEducation)
}

// typowana funkcja, bierze responsewritera i request jako arg i zwraca typ error
type apiFunc func(http.ResponseWriter, *http.Request) error

// funkcja tworzaca implementacje dla typow funkcji apiFunc (error handling + funkcjonalnosc)
func makeHTTPHandler(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil { //jezeli wystapi error w funkcji o typie apiFunc, idz dalej (zawsze wystapi bo apiFunc musi zwracac err)
			if e, ok := err.(apiError); ok { //probuje konwertowac zmienna err do typu apiError, jesli sie uda - przejdz dalej
				WriteJSON(w, e.Status, e) // wypisuje JSON'a do responseWriter'a + status + any (err msg)
				return
			}
			WriteJSON(w, http.StatusInternalServerError, apiError{Err: "internal server error"})
		}
	}
}

// WriteJSON -> zwraca JSON z odpowiednim statusem + naglowkiem
func WriteJSON(w http.ResponseWriter, status int, v any) error { //response writer -> konstruuje odpowiedzi HTTP
	w.WriteHeader(status)                              //wpisuje status Request'a do response writera
	w.Header().Add("Content-Type", "application/json") //dodawnanie naglowka http
	fmt.Printf("%v Request Status: %d \n", time.Now().UTC(), status)
	return json.NewEncoder(w).Encode(v) //zwracanie JSON'a (na poczatku tworzenie encodera z ResponseWriter potem kodowanie jasona do strumienia)
}

func GetID(w http.ResponseWriter, r *http.Request) (int, error) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return -1, err
	}

	return id, nil
}
