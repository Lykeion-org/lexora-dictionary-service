package db

import (
	"github.com/google/uuid"
)

type Referent struct {
	UID         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	EnReference string    `gorm:"not null"`
	ImageSource *string
	Symbols     []Symbol `gorm:"many2many:referent_symbols"`
}

type Symbol struct {
	UID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Language   int       `gorm:"not null"`
	SymbolType int       `gorm:"not null"`
	LemmaUID   *uuid.UUID
	Lemma      *Word  `gorm:"foreignKey:LemmaUID"`
	Words      []Word `gorm:"many2many:symbol_words"`
}

type Word struct {
	UID         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Word        string    `gorm:"not null"`
	SoundSource *string
	IPA         *string
	WordType    int
}
