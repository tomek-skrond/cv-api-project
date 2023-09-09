package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (s *APIServer) handleProjects(w http.ResponseWriter, r *http.Request) error {

	if r.Method == "GET" {
		return s.handleGetProjects(w, r)
	}
	if r.Method == "POST" {
		return s.handleCreateProject(w, r)
	}

	return apiError{Err: "bad request", Status: http.StatusBadRequest}
}

func (s *APIServer) handleProjectsByID(w http.ResponseWriter, r *http.Request) error {

	if r.Method == "GET" {
		return s.handleGetProjectByID(w, r)
	}

	if r.Method == "PUT" {
		return s.handleUpdateProjectByID(w, r)
	}
	if r.Method == "DELETE" {
		return s.handleDeleteProjectByID(w, r)
	}

	return apiError{Err: "bad request", Status: http.StatusBadRequest}
}

func (s *APIServer) handleGetProjects(w http.ResponseWriter, r *http.Request) error {

	projArray, err := s.Store.GetProjects()
	if err != nil {
		fmt.Println("store err")
		return err
	}

	return WriteJSON(w, http.StatusOK, projArray)
}

func (s *APIServer) handleCreateProject(w http.ResponseWriter, r *http.Request) error {

	newProjRequest := new(Project)

	if err := json.NewDecoder(r.Body).Decode(newProjRequest); err != nil {
		return err
	}

	newProject, err := NewProject(
		newProjRequest.ProjectName,
		newProjRequest.TechnologyUsed,
		newProjRequest.Description,
	)
	if err != nil {
		fmt.Println("constructor err")
		return err
	}

	if err := s.Store.CreateProject(newProject); err != nil {
		fmt.Println("store err createproj")
		return err
	}

	return WriteJSON(w, http.StatusOK, newProject)

}

func (s *APIServer) handleGetProjectByID(w http.ResponseWriter, r *http.Request) error {
	id, err := GetID(w, r)
	if err != nil {
		return err
	}

	proj, err := s.Store.GetProjectsByID(id)
	if err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, proj)

}
func (s *APIServer) handleDeleteProjectByID(w http.ResponseWriter, r *http.Request) error {
	id, err := GetID(w, r)
	if err != nil {
		return err
	}

	if err := s.Store.DeleteProjectByID(id); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, nil)
}

func (s *APIServer) handleUpdateProjectByID(w http.ResponseWriter, r *http.Request) error {
	//1. Create buffer education object
	//2. Decode update Request
	//3. Assign updated values
	id, err := GetID(w, r)
	if err != nil {
		return err
	}
	projUpdateRequest := new(Project)
	projUpdateRequest.ID = id

	if err := json.NewDecoder(r.Body).Decode(projUpdateRequest); err != nil {
		return err
	}

	updatedProj := &Project{
		ID:             projUpdateRequest.ID,
		ProjectName:    projUpdateRequest.ProjectName,
		TechnologyUsed: projUpdateRequest.TechnologyUsed,
		Description:    projUpdateRequest.Description,
	}

	return s.Store.UpdateProject(updatedProj)
}
