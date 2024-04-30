// service/issue_service.go

package service

import (
	"errors"

	"github.com/atx-ai/its-backend/model"
	"gorm.io/gorm"
)

type IssueService struct {
	DB *gorm.DB
}

func NewIssueService(db *gorm.DB) *IssueService {
	return &IssueService{DB: db}
}

// CreateIssue creates a new issue in the database
func (s *IssueService) CreateIssue(issue *model.Issue) error {
	return s.DB.Create(issue).Error
}

// GetIssue retrieves an issue by its ID from the database
func (s *IssueService) GetIssue(id uint) (*model.Issue, error) {
	issue := &model.Issue{}
	if err := s.DB.First(issue, id).Error; err != nil {
		return nil, err
	}
	return issue, nil
}

// UpdateIssue updates an existing issue in the database
func (s *IssueService) UpdateIssue(issue *model.Issue) error {
	return s.DB.Save(issue).Error
}

// DeleteIssue deletes an issue from the database
func (s *IssueService) DeleteIssue(id uint) error {
	return s.DB.Delete(&model.Issue{}, id).Error
}

// ListIssues returns a list of all issues from the database
func (s *IssueService) ListIssues() ([]*model.Issue, error) {
	var issues []*model.Issue
	if err := s.DB.Find(&issues).Error; err != nil {
		return nil, err
	}
	return issues, nil
}

func (s *IssueService) PatchIssue(id uint, updateFields map[string]interface{}) error {
	// Fetch the issue from the database
	var issue model.Issue
	if err := s.DB.First(&issue, id).Error; err != nil {
		return errors.New("issue not found")
	}

	// Update the specific fields
	if err := s.DB.Model(&issue).Updates(updateFields).Error; err != nil {
		return err
	}

	return nil
}
