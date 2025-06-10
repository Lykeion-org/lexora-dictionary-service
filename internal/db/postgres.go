package db

import (
    "fmt"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

func ConnectAndMigrate(dsn string) (*gorm.DB, error) {
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        return nil, err
    }

    err = db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`).Error
    if err != nil {
        return nil, fmt.Errorf("enable uuid extension: %w", err)
    }

    err = db.AutoMigrate(&Referent{}, &Symbol{}, &Word{})
    if err != nil {
        return nil, fmt.Errorf("auto-migrate failed: %w", err)
    }

    return db, nil
}
