package main

///////////// TUTAJ API /////////////
import (
	"encoding/json"
	"fmt"
	"net/http"

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

	fmt.Printf("Server listening at port %s \n", s.ListenPort)
	http.ListenAndServe(s.ListenPort, nil)
}

// API //
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
	return nil
}
func (s *APIServer) handleCreateEducation(w http.ResponseWriter, r *http.Request) error {
	return nil
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
	return json.NewEncoder(w).Encode(v)                //zwracanie JSON'a (na poczatku tworzenie encodera z ResponseWriter potem kodowanie jasona do strumienia)
}
