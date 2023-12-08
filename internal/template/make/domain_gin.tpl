package {{ .FileNameTitleLower }}

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type {{ .FileName }} struct {
	gorm.Model
	ID         string `gorm:"primary_key"`
}

type {{ .FileName }}Repository interface {
	Add(db *gorm.DB) (error)
}

type {{ .FileName }}Service interface {
	Add() (err error)
}

type {{ .FileName }}Controller interface {
	Add(ctx *gin.Context) error
}

func (u *{{ .FileName }}) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New().String()
	return
}
