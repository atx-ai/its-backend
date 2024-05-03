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
	Priority         string         `json:"priority"`                                           // issue prioroty
	Tags             pq.StringArray `json:"tags" gorm:"type:text[]" swaggertype:"array,string"` // tags to filter issue
	ETA              time.Time      `json:"eta"`                                                // time whne issue will be fixed
	AssignedTo       string         `json:"assigned_to"`                                        // name of assignee of the issue
	CreatedAt        time.Time      `json:"created_on"`                                         // issue created date
	UpdatedAt        time.Time      `json:"updated_on"`                                         // issue updated time
} // @name model.Issue
