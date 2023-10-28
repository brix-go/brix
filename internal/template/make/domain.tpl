package {{ .FileNameTitleLower }}

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type {{ .FileName }} struct {
	gorm.Model
	ID         string `gorm:"primary_key"`
}

type {{ .FileName }}Repository interface {
	Add() (error)

}

type {{ .FileName }}Service interface {
	Add() (err error)
}

type {{ .FileName }}Controller interface {
	Add(ctx *fiber.Ctx) error
}

func (u *{{ .FileName }}) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New().String()
	return
}
