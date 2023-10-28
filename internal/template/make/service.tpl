package service

import (
	"{{ .ProjectName }}/internal/domain/{{ .FileNameTitleLower }}"
	"{{ .ProjectName }}/internal/domain/{{ .FileNameTitleLower }}/repository"
)

func NewService(repo {{ .FileNameTitleLower }}.{{ .FileName }}Repository) {{ .FileNameTitleLower }}.{{ .FileName }}Service {
	return &{{ .FileNameTitleLower }}Service{
		repo:        repo,
		redisClient: redis,
	}
}

type {{ .FileNameTitleLower }}Service struct {
	repo user.UserRepository
}

func (s *{{ .FileNameTitleLower }}Service) Add() (err error) {
	//TODO: Implement register
	return  nil
}
