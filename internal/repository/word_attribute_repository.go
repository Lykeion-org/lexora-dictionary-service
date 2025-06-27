package repository

import (
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"github.com/Lykeion-org/lexora-dictionary-service/internal/db"
)

type WordAttributesRepository interface {
	CreateWordAttributes(ctx context.Context, attr *db.WordAttributes) (*db.WordAttributes, error)
	GetWordAttributes(ctx context.Context, wordUID uuid.UUID) (*db.WordAttributes, error)
	UpdateWordAttributes(ctx context.Context, attr *db.WordAttributes) error
	DeleteWordAttributes(ctx context.Context, wordUID uuid.UUID) error
}

type wordAttributesRepository struct {
	db *gorm.DB
}

func NewWordAttributesRepository(db *gorm.DB) *wordAttributesRepository {
	return &wordAttributesRepository{
		db: db,
	}
}

func (r *wordAttributesRepository) CreateWordAttributes(ctx context.Context, attr *db.WordAttributes) (*db.WordAttributes, error) {
	if err := r.db.WithContext(ctx).Create(attr).Error; err != nil {
		return nil, err
	}
	return attr, nil
}

func (r *wordAttributesRepository) UpdateWordAttributes(ctx context.Context, attr *db.WordAttributes) error {
	return r.db.WithContext(ctx).Save(attr).Error
}

func (r *wordAttributesRepository) DeleteWordAttributes(ctx context.Context, wordUID uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&db.WordAttributes{}, "word_id = ?", wordUID).Error
}

func (r *wordAttributesRepository) GetWordAttributes(ctx context.Context, wordUID uuid.UUID) (*db.WordAttributes, error) {
	var attr db.WordAttributes
	if err := r.db.WithContext(ctx).First(&attr, "word_id = ?", wordUID).Error; err != nil {
		return nil, err
	}
	return &attr, nil
}
