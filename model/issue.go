package model

import (
	"time"

	"github.com/lib/pq"
)

type Issue struct {
	ID               uint           `gorm:"primaryKey"`
	CreatedBy        string         `json:"created_by"`
	IssueID          string         `json:"issue_id"`
	IssueDescription string         `json:"issue_description"`
	Category         string         `json:"category"`
	State            string         `json:"state"`
	Priority         string         `json:"priority"`
	Tags             pq.StringArray `json:"tags" gorm:"type:text[]"`
	ETA              time.Time      `json:"eta"`
	AssignedTo       string         `json:"assigned_to"`
	CreatedOn        time.Time      `json:"created_on"`
	UpdatedOn        time.Time      `json:"updated_on"`
}
