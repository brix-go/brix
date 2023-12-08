package repository

import (
	"{{ .ProjectName }}/internal/domain/{{ .FileNameTitleLower }}"
	"gorm.io/gorm"
)

func NewRepository(db *gorm.DB) {{ .FileNameTitleLower }}.{{ .FileName }}Repository {
	return &{{ .FileNameTitleLower }}Repository{
	}
}

type {{ .FileNameTitleLower }}Repository struct {
}

func (r *{{ .FileNameTitleLower }}Repository) Add(db *gorm.DB) (error) {
	return nil
}
