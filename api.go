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
	router.HandleFunc("/experience/{id}", makeHTTPHandler(s.handleExperienceByID))
	router.HandleFunc("/languages", makeHTTPHandler(s.handleLanguages))
	router.HandleFunc("/languages/{id}", makeHTTPHandler(s.handleLanguagesByID))
	router.HandleFunc("/projects", makeHTTPHandler(s.handleProjects))
	router.HandleFunc("/projects/{id}", makeHTTPHandler(s.handleProjectsByID))

	fmt.Printf("Server listening at port %s \n", s.ListenPort)
	if err := http.ListenAndServe(s.ListenPort, router); err != nil {
		fmt.Println(err)
	}
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
