package repository

import (
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"github.com/Lykeion-org/lexora-dictionary-service/internal/db"
)

type SymbolRepository interface {
	CreateSymbol(ctx context.Context, sym *db.Symbol) (*db.Symbol, error)
	GetSymbol(ctx context.Context, uid uuid.UUID) (*db.Symbol, error)
	UpdateSymbol(ctx context.Context, sym *db.Symbol) error
	DeleteSymbol(ctx context.Context, uid uuid.UUID) error
	ListSymbolsByWord(ctx context.Context, word string) ([]db.Symbol, error)
}

type symbolRepository struct {
	db *gorm.DB
}

func NewSymbolRepository(db *gorm.DB) *symbolRepository {
	return &symbolRepository{
		db: db,
	}
}

func (s *symbolRepository) CreateSymbol(ctx context.Context, sym *db.Symbol) (*db.Symbol, error) {
	if sym.UID == uuid.Nil {
		sym.UID = uuid.New()
	}
	if err := s.db.WithContext(ctx).Create(sym).Error; err != nil {
		return nil, err
	}
	return sym, nil
}
func (s *symbolRepository) UpdateSymbol(ctx context.Context, sym *db.Symbol) error {
	return s.db.WithContext(ctx).Save(sym).Error
}

func (s *symbolRepository) DeleteSymbol(ctx context.Context, uid uuid.UUID) error {
	return s.db.WithContext(ctx).Delete(&db.Symbol{}, "uid = ?", uid).Error
}

func (s *symbolRepository) GetSymbol(ctx context.Context, uid uuid.UUID) (*db.Symbol, error) {
	var data db.Symbol
	if err := s.db.WithContext(ctx).
		Preload("Words").
		Preload("Lemma").
		First(&data, "uid = ?", uid).
		Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (s *symbolRepository) ListSymbolsByWord(ctx context.Context, word string) ([]db.Symbol, error) {
	var symbols []db.Symbol

	err := s.db.
		WithContext(ctx).
		Model(&db.Symbol{}).
		Joins("JOIN symbol_words ON symbol_words.symbol_id = symbols.uid").
		Joins("JOIN words ON words.uid = symbol_words.word_id").
		Where("words.word = ?", word).
		Preload("Words"). // optional: also pull related Words for each Symbol
		Find(&symbols).Error

	return symbols, err
}
