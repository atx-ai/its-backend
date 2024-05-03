package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/atx-ai/its-backend/model"
	"github.com/atx-ai/its-backend/service"
	chi "github.com/go-chi/chi/v5"
)

type CommnetController struct {
	CommnetService *service.CommnetService
}

func NewCommnetController(commnetService *service.CommnetService) *CommnetController {
	return &CommnetController{
		CommnetService: commnetService,
	}
}

func (c *CommnetController) Routes() chi.Router {
	r := chi.NewRouter()
	r.Get("/", c.ListCommnetsHandler)
	r.Post("/", c.CreateCommnetHandler)
	r.Route("/{id}", func(r chi.Router) {
		r.Get("/", c.GetCommnetByIDHandler)
		r.Put("/", c.UpdateCommnetHandler)
		r.Delete("/", c.DeleteCommnetHandler)
	})

	return r
}

// @summary List comments for a specific issue
// @description List all comments associated with a specific issue
// @tags comments
// @param issueID path uint true "Issue ID"
// @produce json
// @success 200 {array} model.Commnet
// @router /issues/{issueID}/comments [get]
func (c *CommnetController) ListCommnetsHandler(w http.ResponseWriter, r *http.Request) {
	issueID := chi.URLParam(r, "issueID")
	issueIDUint, err := strconv.ParseUint(issueID, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	comments, err := c.CommnetService.ListCommnets(r.Context(), uint(issueIDUint))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(comments)
}

// @summary Create a new comment
// @description Create a new comment for the specified issue
// @tags comments
// @param issueID path uint true "Issue ID"
// @accept json
// @produce json
// @param comment body model.Commnet true "Comment details"
// @success 201 {object} model.Commnet
// @router /issues/{issueID}/comments [post]
func (c *CommnetController) CreateCommnetHandler(w http.ResponseWriter, r *http.Request) {
	issueID := chi.URLParam(r, "issueID")
	issueIDUint, err := strconv.ParseUint(issueID, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var commnet model.Commnet
	err = json.NewDecoder(r.Body).Decode(&commnet)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	commnet.IssueID = uint(issueIDUint)

	err = c.CommnetService.CreateCommnet(r.Context(), &commnet)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(commnet)
}

// @summary Get a comment by ID
// @description Get a comment by its ID for the specified issue
// @tags comments
// @param issueID path uint true "Issue ID"
// @param id path uint true "Comment ID"
// @produce json
// @success 200 {object} model.Commnet
// @router /issues/{issueID}/comments/{id} [get]
func (c *CommnetController) GetCommnetByIDHandler(w http.ResponseWriter, r *http.Request) {
	issueID := chi.URLParam(r, "issueID")
	issueIDUint, err := strconv.ParseUint(issueID, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id := chi.URLParam(r, "id")
	commnetID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	commnet, err := c.CommnetService.GetCommnetByID(r.Context(), uint(commnetID))
	if err != nil || commnet.IssueID != uint(issueIDUint) {
		http.Error(w, "Comment not found for the provided issue", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(commnet)
}

// @summary Update a comment
// @description Update a comment by its ID for the specified issue
// @tags comments
// @param issueID path uint true "Issue ID"
// @param id path uint true "Comment ID"
// @accept json
// @produce json
// @param comment body model.Commnet true "Updated comment details"
// @success 200
// @router /issues/{issueID}/comments/{id} [put]
func (c *CommnetController) UpdateCommnetHandler(w http.ResponseWriter, r *http.Request) {
	issueID := chi.URLParam(r, "issueID")
	issueIDUint, err := strconv.ParseUint(issueID, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id := chi.URLParam(r, "id")
	commnetID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var commnet model.Commnet
	err = json.NewDecoder(r.Body).Decode(&commnet)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	commnet.ID = uint(commnetID)
	commnet.IssueID = uint(issueIDUint)

	err = c.CommnetService.UpdateCommnet(r.Context(), &commnet)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// @summary Delete a comment
// @description Delete a comment by its ID for the specified issue
// @tags comments
// @param issueID path uint true "Issue ID"
// @param id path uint true "Comment ID"
// @success 204
// @router /issues/{issueID}/comments/{id} [delete]
func (c *CommnetController) DeleteCommnetHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	commnetID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.CommnetService.DeleteCommnet(r.Context(), uint(commnetID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
