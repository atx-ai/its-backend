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

// @Summary Get an issue by ID
// @Description Get an issue by its ID
// @Accept json
// @Produce json
// @Param id path int true "Issue ID"
// @Success 200 {object} model.Issue
// @Router /issues/{id} [get]
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

// @Summary Create a new issue
// @Description Create a new issue
// @Accept json
// @Produce json
// @Param issue body model.Issue true "Issue object to be created"
// @Success 201 {object} model.Issue
// @Router /issues [post]
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

// @Summary Update an existing issue
// @Description Update an existing issue
// @Accept json
// @Produce json
// @Param id path int true "Issue ID"
// @Param issue body model.Issue true "Updated issue object"
// @Success 200 {object} model.Issue
// @Router /issues/{id} [put]
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

// @Summary Delete an issue by ID
// @Description Delete an issue by its ID
// @Accept json
// @Produce json
// @Param id path int true "Issue ID"
// @Success 200 "OK"
// @Router /issues/{id} [delete]
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

// @Summary List all issues
// @Description Get a list of all issues
// @Produce json
// @Success 200 {array} model.Issue
// @Router /issues [get]
func (c *IssueController) ListIssues(w http.ResponseWriter, r *http.Request) {
	issues, err := c.Service.ListIssues()
	if err != nil {
		http.Error(w, "Failed to fetch issues", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(issues)
}
