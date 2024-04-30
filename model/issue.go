package model

import (
	"time"

	"github.com/lib/pq"
)

// @Description Issue Structure
type Issue struct {
	ID               uint           `json:"issue_id" gorm:"primaryKey"`
	CreatedBy        string         `json:"created_by"`                                         // name of user who creates issue
	IssueDescription string         `json:"issue_description"`                                  // issue descriptiom
	Category         string         `json:"category"`                                           // issue category
	State            string         `json:"state"`                                              // issue state
	Priority         string         `json:"priority"`                                           // name of user who creates issue
	Tags             pq.StringArray `json:"tags" gorm:"type:text[]" swaggertype:"array,string"` // name of user who creates issue
	ETA              time.Time      `json:"eta"`                                                // name of user who creates issue
	AssignedTo       string         `json:"assigned_to"`                                        // name of user who creates issue
	CreatedOn        time.Time      `json:"created_on"`                                         // name of user who creates issue
	UpdatedOn        time.Time      `json:"updated_on"`                                         // name of user who creates issue
} // @name model.Issue
