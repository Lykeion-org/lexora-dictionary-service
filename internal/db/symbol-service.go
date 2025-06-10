package db

import (
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SymbolService struct {
	DB gorm.DB
}

func (s *SymbolService) CreateSymbol(ctx context.Context, sym *Symbol) (*Symbol, error) {
	if sym.UID == uuid.Nil {
		sym.UID = uuid.New()
	}
	if err := s.DB.WithContext(ctx).Create(sym).Error; err != nil {
		return nil, err
	}
	return sym, nil
}
func (s *SymbolService) UpdateSymbol(ctx context.Context, sym *Symbol) error {
	return s.DB.WithContext(ctx).Save(sym).Error
}

func (s *SymbolService) DeleteSymbol(ctx context.Context, uid uuid.UUID) error {
	return s.DB.WithContext(ctx).Delete(&Symbol{}, "uid = ?", uid).Error
}

func (s *SymbolService) GetSymbol(ctx context.Context, uid uuid.UUID) (*Symbol, error) {
	var data Symbol
	if err := s.DB.WithContext(ctx).
		Preload("Words").
		Preload("Lemma").
		First(&data, "uid = ?", uid).
		Error; err != nil {
		return nil, err
	}
	return &data, nil

}
