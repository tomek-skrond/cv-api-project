package main

///////////// TUTAJ DEFINICJE DLA ERROR HANDLINGU /////////////

import "net/http"

var ErrUserInvalid = apiError{Err: "user not valid", Status: http.StatusForbidden}
var ErrInvalidMethod = apiError{Err: "Invalid method", Status: http.StatusMethodNotAllowed}

// "wlasna" definicja error'a
type apiError struct {
	Err    string //err message
	Status int    //status
}

//funkcja zwracajaca error jako string (f. zdefiniowana przez strukture apiError)
func (e apiError) Error() string {
	return e.Err
}
