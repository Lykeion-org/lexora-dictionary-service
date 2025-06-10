package db

import (
    "github.com/google/uuid"
)

type Referent struct {
    UID         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
    EnReference string    `gorm:"not null"`
    ImageSource *string
    Symbols     []Symbol `gorm:"foreignKey:ReferentUID"`
}

type Symbol struct {
    UID         uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
    ReferentUID uuid.UUID  `gorm:"type:uuid;not null"`
    Language    int
    SymbolType  int
    LemmaUID    *uuid.UUID
    Lemma       *Word      `gorm:"foreignKey:LemmaUID"`
    Words       []Word     `gorm:"many2many:symbol_words"`
}

type Word struct {
    UID         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
    Word        string    `gorm:"not null"`
    SoundSource *string
    IPA         *string
    WordType    int
}
