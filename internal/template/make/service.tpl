package service

import (
	"{{ .ProjectName }}/internal/domain/{{ .FileNameTitleLower }}"
)

func NewService(repo {{ .FileNameTitleLower }}.{{ .FileName }}Repository) {{ .FileNameTitleLower }}.{{ .FileName }}Service {
	return &{{ .FileNameTitleLower }}Service{
		repo:        repo,
	}
}

type {{ .FileNameTitleLower }}Service struct {
	repo {{ .FileNameTitleLower }}.{{ .FileName }}Repository
}

func (s *{{ .FileNameTitleLower }}Service) Add() (err error) {
	//TODO: Implement register
	return  nil
}
