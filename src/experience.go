package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// API - Experience//
func (s *APIServer) handleExperience(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGetExperience(w, r)
	}
	if r.Method == "POST" {
		return s.handleCreateExperience(w, r)
	}

	return apiError{Err: "Method not implemented", Status: http.StatusNotImplemented}
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

func (s *APIServer) handleExperienceByID(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGetExperienceByID(w, r)
	}
	if r.Method == "DELETE" {
		return s.handleDeleteExperienceByID(w, r)
	}

	if r.Method == "PUT" {
		return s.handleUpdateExperienceByID(w, r)
	}

	return apiError{Err: "Invalid Method", Status: http.StatusBadRequest}
}

func (s *APIServer) handleGetExperienceByID(w http.ResponseWriter, r *http.Request) error {
	id, err := GetID(w, r)
	if err != nil {
		return err
	}

	exp, err := s.Store.GetExperienceByID(id)
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, exp)
}

func (s *APIServer) handleDeleteExperienceByID(w http.ResponseWriter, r *http.Request) error {
	id, err := GetID(w, r)
	if err != nil {
		return err
	}
	if err := s.Store.DeleteExperienceByID(id); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, nil)
}

func (s *APIServer) handleUpdateExperienceByID(w http.ResponseWriter, r *http.Request) error {
	id, err := GetID(w, r)
	if err != nil {
		return err
	}

	updateExpReq := new(Experience)
	updateExpReq.ID = id

	if err := json.NewDecoder(r.Body).Decode(updateExpReq); err != nil {
		return err
	}

	updatedExp := &Experience{
		ID:          updateExpReq.ID,
		Company:     updateExpReq.Company,
		Role:        updateExpReq.Role,
		DateStarted: updateExpReq.DateStarted,
		DateEnded:   updateExpReq.DateEnded,
	}

	return s.Store.UpdateExperience(updatedExp)
}
