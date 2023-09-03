package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

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
		return s.handleGetEducationByID(w, id)
	}

	if r.Method == "DELETE" {
		return s.Store.DeleteEducationByID(id)
	}

	if r.Method == "PUT" {
		return s.handleUpdateEducationByID(r, id)
	}
	return apiError{Err: "error", Status: http.StatusBadRequest}
}

func (s *APIServer) handleUpdateEducationByID(r *http.Request, id int) error {
	//1. Create buffer education object
	//2. Decode update Request
	//3. Assign updated values
	//4.
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

	return s.Store.UpdateEducation(updatedEdu)
}

func (s *APIServer) handleGetEducationByID(w http.ResponseWriter, id int) error {
	edu, err := s.Store.GetEducationByID(id)
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, edu)
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
