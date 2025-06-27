package repository

import (
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"github.com/Lykeion-org/lexora-dictionary-service/internal/db"
)

type SuggestionRepository interface {
	CreateSuggestion(ctx context.Context, suggestion *db.WordSuggestion) (*db.WordSuggestion, error)
	GetSuggestion(ctx context.Context, uid uuid.UUID) (*db.WordSuggestion, error)
	UpdateSuggestion(ctx context.Context, suggestion *db.WordSuggestion) error
	DeleteSuggestion(ctx context.Context, uid uuid.UUID) error
	ListPendingSuggestions(ctx context.Context) ([]db.WordSuggestion, error)
	ApproveSuggestion(ctx context.Context, uid uuid.UUID, approvedBy string) error
}

type suggestionRepository struct {
	db *gorm.DB
}

func NewSuggestionRepository(db *gorm.DB) *suggestionRepository {
	return &suggestionRepository{
		db: db,
	}
}

func (r *suggestionRepository) CreateSuggestion(ctx context.Context, suggestion *db.WordSuggestion) (*db.WordSuggestion, error) {
	if suggestion.UID == uuid.Nil {
		suggestion.UID = uuid.New()
	}
	if err := r.db.WithContext(ctx).Create(suggestion).Error; err != nil {
		return nil, err
	}
	return suggestion, nil
}

func (r *suggestionRepository) UpdateSuggestion(ctx context.Context, suggestion *db.WordSuggestion) error {
	return r.db.WithContext(ctx).Save(suggestion).Error
}

func (r *suggestionRepository) DeleteSuggestion(ctx context.Context, uid uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&db.WordSuggestion{}, "uid = ?", uid).Error
}

func (r *suggestionRepository) GetSuggestion(ctx context.Context, uid uuid.UUID) (*db.WordSuggestion, error) {
	var suggestion db.WordSuggestion
	if err := r.db.WithContext(ctx).Preload("WordAttributes").First(&suggestion, "uid = ?", uid).Error; err != nil {
		return nil, err
	}
	return &suggestion, nil
}

func (r *suggestionRepository) ListPendingSuggestions(ctx context.Context) ([]db.WordSuggestion, error) {
	var suggestions []db.WordSuggestion
	err := r.db.WithContext(ctx).
		Where("approved = ?", false).
		Find(&suggestions).Error
	return suggestions, err
}

func (r *suggestionRepository) ApproveSuggestion(ctx context.Context, uid uuid.UUID, approvedBy string) error {
	return r.db.WithContext(ctx).
		Model(&db.WordSuggestion{}).
		Where("uid = ?", uid).
		Updates(map[string]interface{}{
			"approved": true,
			"approved_by": approvedBy,
		}).Error
}
