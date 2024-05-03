package service

import (
	"context"

	"github.com/atx-ai/its-backend/model"
	"gorm.io/gorm"
)

type CommnetService struct {
	DB *gorm.DB
}

func NewCommnetService(db *gorm.DB) *CommnetService {
	return &CommnetService{
		DB: db,
	}
}

func (s *CommnetService) CreateCommnet(ctx context.Context, commnet *model.Commnet) error {
	return s.DB.Create(commnet).Error
}

func (s *CommnetService) GetCommnetByID(ctx context.Context, id uint) (*model.Commnet, error) {
	var commnet model.Commnet
	if err := s.DB.First(&commnet, id).Error; err != nil {
		return nil, err
	}
	return &commnet, nil
}

func (s *CommnetService) UpdateCommnet(ctx context.Context, commnet *model.Commnet) error {
	return s.DB.Save(commnet).Error
}

func (s *CommnetService) DeleteCommnet(ctx context.Context, id uint) error {
	return s.DB.Delete(&model.Commnet{}, id).Error
}

func (s *CommnetService) ListCommnets(ctx context.Context, issueID uint) ([]model.Commnet, error) {
	var comments []model.Commnet
	if err := s.DB.Where("issue_id = ?", issueID).Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}
