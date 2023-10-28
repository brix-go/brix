package repository

import (
	"{{ .ProjectName }}/internal/domain/{{ .FileNameTitleLower }}"
	"gorm.io/gorm"
)

func NewRepository(db *gorm.DB) {{ .FileNameTitleLower }}.{{ .FileName }}Repository {
	return &{{ .FileNameTitleLower }}Repository{
		db: db,
	}
}

type userRepository struct {
	db *gorm.DB
}

func (r *userRepository) Add() (error) {
	return updatedUser, nil
}
