package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// API - Languages//
func (s *APIServer) handleLanguages(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGetLanguages(w, r)
	}
	if r.Method == "POST" {
		return s.handleCreateLanguages(w, r)
	}

	return apiError{Err: "Method not implemented", Status: http.StatusNotImplemented}
}

func (s *APIServer) handleGetLanguages(w http.ResponseWriter, r *http.Request) error {
	// invoke db operation
	// encode db response to JSON + return
	eduArray, err := s.Store.GetLanguages()
	if err != nil {
		fmt.Println("store err")
		return err
	}

	return WriteJSON(w, http.StatusOK, eduArray)
}

func (s *APIServer) handleLanguagesByID(w http.ResponseWriter, r *http.Request) error {
	id, err := GetID(w, r)
	if err != nil {
		return err
	}

	if r.Method == "GET" {
		return s.handleGetLanguagesByID(w, id)
	}

	if r.Method == "DELETE" {
		return s.Store.DeleteLanguageByID(id)
	}

	if r.Method == "PUT" {
		return s.handleUpdateLanguagesByID(r, id)
	}
	return apiError{Err: "error", Status: http.StatusBadRequest}
}

func (s *APIServer) handleUpdateLanguagesByID(r *http.Request, id int) error {
	//1. Create buffer education object
	//2. Decode update Request
	//3. Assign updated values
	langUpdateRequest := new(Language)
	langUpdateRequest.ID = id

	if err := json.NewDecoder(r.Body).Decode(langUpdateRequest); err != nil {
		return err
	}

	updatedEdu := &Language{
		ID:          langUpdateRequest.ID,
		Language:    langUpdateRequest.Language,
		Level:       langUpdateRequest.Level,
		Description: langUpdateRequest.Description,
	}

	return s.Store.UpdateLanguage(updatedEdu)
}

func (s *APIServer) handleGetLanguagesByID(w http.ResponseWriter, id int) error {
	edu, err := s.Store.GetLanguageByID(id)
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, edu)
}
func (s *APIServer) handleCreateLanguages(w http.ResponseWriter, r *http.Request) error {
	// create empty object
	// fetch data from request (decode)
	// feed data to empty object
	// invoke database operation (s.Store.CreateLanguages())
	// return json with data
	createLanguagesRequest := new(Language)

	if err := json.NewDecoder(r.Body).Decode(createLanguagesRequest); err != nil {
		fmt.Println("formatting err")
		return err
	}

	newLanguages, err := NewLanguage(
		createLanguagesRequest.Language,
		createLanguagesRequest.Level,
		createLanguagesRequest.Description,
	)

	if err != nil {
		fmt.Println("constructor err")
		return err
	}

	if err := s.Store.CreateLanguage(newLanguages); err != nil {
		fmt.Println("store err")
		return err
	}

	return WriteJSON(w, http.StatusOK, newLanguages)
}
