package repository

import (
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"github.com/Lykeion-org/lexora-dictionary-service/internal/db"
)

type RelationshipRepository interface {
	AddSymbolToReferent(ctx context.Context, referentUID, symbolUID uuid.UUID) error
	RemoveSymbolFromReferent(ctx context.Context, referentUID, symbolUID uuid.UUID) error
	AddWordToSymbol(ctx context.Context, symbolUID, wordUID uuid.UUID) error
	RemoveWordFromSymbol(ctx context.Context, symbolUID, wordUID uuid.UUID) error
	SetSymbolLemma(ctx context.Context, symbolUID, wordUID uuid.UUID) error
	ClearSymbolLemma(ctx context.Context, symbolUID uuid.UUID) error
}

type relationshipRepository struct {
	db *gorm.DB
}

func NewRelationshipRepository(db *gorm.DB) *relationshipRepository {
	return &relationshipRepository{
		db: db,
	}
}

func (r *relationshipRepository) AddSymbolToReferent(ctx context.Context, referentUID, symbolUID uuid.UUID) error {
	return r.db.
		WithContext(ctx).
		Model(&db.Referent{UID: referentUID}).
		Association("Symbols").
		Append(&db.Symbol{UID: symbolUID})
}

func (r *relationshipRepository) RemoveSymbolFromReferent(ctx context.Context, referentUID, symbolUID uuid.UUID) error {
	return r.db.WithContext(ctx).
		Model(&db.Referent{UID: referentUID}).
		Association("Symbols").
		Delete(&db.Symbol{UID: symbolUID})
}

func (r *relationshipRepository) AddWordToSymbol(ctx context.Context, symbolUID uuid.UUID, wordUID uuid.UUID) error {
	return r.db.
		WithContext(ctx).
		Model(&db.Symbol{UID: symbolUID}).
		Association("Words").
		Append(&db.Word{UID: wordUID})
}

func (r *relationshipRepository) RemoveWordFromSymbol(ctx context.Context, symbolUID, wordUID uuid.UUID) error {
	return r.db.WithContext(ctx).
		Model(&db.Symbol{UID: symbolUID}).
		Association("Words").
		Delete(&db.Word{UID: wordUID})
}

func (r *relationshipRepository) SetSymbolLemma(ctx context.Context, symbolUID, wordUID uuid.UUID) error {
	return r.db.WithContext(ctx).
		Model(&db.Symbol{UID: symbolUID}).
		Update("lemma_uid", wordUID).Error
}

func (r *relationshipRepository) ClearSymbolLemma(ctx context.Context, symbolUID uuid.UUID) error {
	return r.db.WithContext(ctx).
		Model(&db.Symbol{UID: symbolUID}).
		Update("lemma_uid", nil).Error
}
