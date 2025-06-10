package db

import (
	"context"
    "github.com/google/uuid"
    "gorm.io/gorm"
)

type WordService struct {
	DB gorm.DB
}

func (s *WordService) CreateWord(ctx context.Context, wrd *Word) (*Word, error) {
	if wrd.UID == uuid.Nil {
        wrd.UID = uuid.New()
    }
    if err := s.DB.WithContext(ctx).Create(wrd).Error; err != nil {
        return nil, err
    }
    return wrd, nil
}
func (s *WordService) UpdateWord(ctx context.Context, wrd *Word) error {
	return s.DB.WithContext(ctx).Save(wrd).Error
}

func (s *WordService) DeleteWord(ctx context.Context, uid uuid.UUID) error {
	return s.DB.WithContext(ctx).Delete(&Word{}, "uid = ?", uid).Error
}

func (s *WordService) GetWord(ctx context.Context, uid uuid.UUID) (*Word, error) {
	var data Word
    if err := s.DB.WithContext(ctx).First(&data, "uid = ?", uid).Error; err != nil {
        return nil, err
    }
    return &data, nil
	
}

