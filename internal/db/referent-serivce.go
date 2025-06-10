package db

import (
	"context"
    "github.com/google/uuid"
    "gorm.io/gorm"
)

type ReferentService struct {
	DB gorm.DB
}

func (s *ReferentService) CreateReferent(ctx context.Context, ref *Referent) (*Referent, error) {
	if ref.UID == uuid.Nil {
        ref.UID = uuid.New()
    }
    if err := s.DB.WithContext(ctx).Create(ref).Error; err != nil {
        return nil, err
    }
    return ref, nil
}
func (s *ReferentService) UpdateReferent(ctx context.Context, ref *Referent) error {
	return s.DB.WithContext(ctx).Save(ref).Error
}

func (s *ReferentService) DeleteReferent(ctx context.Context, uid uuid.UUID) error {
	return s.DB.WithContext(ctx).Delete(&Referent{}, "uid = ?", uid).Error
}

func (s *ReferentService) GetReferent(ctx context.Context, uid uuid.UUID) (*Referent, error) {
    var data Referent
    if err := s.DB.WithContext(ctx).Preload("Symbols.Words").First(&data, "uid = ?", uid).Error; err != nil {
        return nil, err
    }
    return &data, nil
}

