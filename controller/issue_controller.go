package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/atx-ai/its-backend/model"
	"github.com/atx-ai/its-backend/service"
	"github.com/go-chi/chi"
)

type IssueController struct {
	Service *service.IssueService
}

func NewIssueController(service *service.IssueService) *IssueController {
	return &IssueController{Service: service}
}

func (c *IssueController) GetIssue(w http.ResponseWriter, r *http.Request) {
	issueID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid issue ID", http.StatusBadRequest)
		return
	}

	issue, err := c.Service.GetIssue(uint(issueID))
	if err != nil {
		http.Error(w, "Issue not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(issue)
}

func (c *IssueController) CreateIssue(w http.ResponseWriter, r *http.Request) {
	var issue model.Issue
	if err := json.NewDecoder(r.Body).Decode(&issue); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := c.Service.CreateIssue(&issue); err != nil {
		http.Error(w, "Failed to create issue", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (c *IssueController) UpdateIssue(w http.ResponseWriter, r *http.Request) {
	issueID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid issue ID", http.StatusBadRequest)
		return
	}

	var updatedIssue model.Issue
	if err := json.NewDecoder(r.Body).Decode(&updatedIssue); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	updatedIssue.ID = uint(issueID)
	if err := c.Service.UpdateIssue(&updatedIssue); err != nil {
		http.Error(w, "Failed to update issue", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (c *IssueController) DeleteIssue(w http.ResponseWriter, r *http.Request) {
	issueID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid issue ID", http.StatusBadRequest)
		return
	}

	if err := c.Service.DeleteIssue(uint(issueID)); err != nil {
		http.Error(w, "Failed to delete issue", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (c *IssueController) ListIssues(w http.ResponseWriter, r *http.Request) {
	issues, err := c.Service.ListIssues()
	if err != nil {
		http.Error(w, "Failed to fetch issues", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(issues)
}
