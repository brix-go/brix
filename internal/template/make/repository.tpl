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

type {{ .FileNameTitleLower }}Repository struct {
	db *gorm.DB
}

func (r *{{ .FileNameTitleLower }}Repository) Add() (error) {
	return nil
}
