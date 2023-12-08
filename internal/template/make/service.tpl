package service

import (
	"{{ .ProjectName }}/internal/domain/{{ .FileNameTitleLower }}"
	"gorm.io/gorm"
)

func NewService(db *gorm.DB,repo {{ .FileNameTitleLower }}.{{ .FileName }}Repository) {{ .FileNameTitleLower }}.{{ .FileName }}Service {
	return &{{ .FileNameTitleLower }}Service{
		repo:        repo,
		db: db,
	}
}

type {{ .FileNameTitleLower }}Service struct {
	repo {{ .FileNameTitleLower }}.{{ .FileName }}Repository
	db        *gorm.DB
}

func (s *{{ .FileNameTitleLower }}Service) Add() (err error) {
    tx := s.db.Begin()
	//TODO: Implement register
	err = s.repo.Add(tx)
	if err != nil {
	    tx.Rollback()
	    return err
	}
	tx = s.db.Commit()
	return  nil
}
