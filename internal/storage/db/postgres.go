package db

import (
	"fmt"

	"simple-leave-tracker/internal/storage"
	"simple-leave-tracker/internal/storage/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type mySQL struct {
	db *gorm.DB
}

func New(dsn string) (storage.Storage, error) {
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		return nil, fmt.Errorf("error when connecting to mysql, %w", err)
	}

	if err := db.AutoMigrate(model.GetModels()...); err != nil {
		return nil, fmt.Errorf("error when migrating modes, %w", err)
	}

	return &mySQL{
		db: db,
	}, nil
}
