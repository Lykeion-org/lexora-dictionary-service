package db

import (
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RelationshipService struct {
	DB gorm.DB
}

// Referent-Symbol relationships
func (s *RelationshipService) AddSymbolToReferent(ctx context.Context, referentUID, symbolUID uuid.UUID) error {
	return s.DB.WithContext(ctx).Model(&Referent{UID: referentUID}).
		Association("Symbols").
		Append(&Symbol{UID: symbolUID})
}

func (s *RelationshipService) RemoveSymbolFromReferent(ctx context.Context, referentUID, symbolUID uuid.UUID) error {
	return s.DB.WithContext(ctx).Model(&Referent{UID: referentUID}).
		Association("Symbols").
		Delete(&Symbol{UID: symbolUID})
}

// Symbol-Word relationships
func (s *RelationshipService) AddWordToSymbol(ctx context.Context, symbolUID, wordUID uuid.UUID) error {
	return s.DB.WithContext(ctx).Model(&Symbol{UID: symbolUID}).
		Association("Words").
		Append(&Word{UID: wordUID})
}

func (s *RelationshipService) RemoveWordFromSymbol(ctx context.Context, symbolUID, wordUID uuid.UUID) error {
	return s.DB.WithContext(ctx).Model(&Symbol{UID: symbolUID}).
		Association("Words").
		Delete(&Word{UID: wordUID})
}

// Lemma relationships
func (s *RelationshipService) SetSymbolLemma(ctx context.Context, symbolUID, wordUID uuid.UUID) error {
	return s.DB.WithContext(ctx).Model(&Symbol{UID: symbolUID}).
		Update("lemma_uid", wordUID).Error
}

func (s *RelationshipService) ClearSymbolLemma(ctx context.Context, symbolUID uuid.UUID) error {
	return s.DB.WithContext(ctx).Model(&Symbol{UID: symbolUID}).
		Update("lemma_uid", nil).Error
}
