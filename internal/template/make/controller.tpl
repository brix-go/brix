package controller

import (
	"github.com/gofiber/fiber/v2"
	"{{ .ProjectName }}/internal/domain/{{ .FileNameTitleLower }}"
	"testdd/shared"
	middleware "testdd/middleware/error"
)

func NewController(service {{ .FileNameTitleLower }}.{{ .FileName }}Service) {{ .FileNameTitleLower }}.{{ .FileName }}Controller {
	return &{{ .FileNameTitleLower }}Controller{
		{{ .FileNameTitleLower }}Service: service,
	}
}

type {{ .FileNameTitleLower }}Controller struct {
	{{ .FileNameTitleLower }}Service {{ .FileNameTitleLower }}.{{ .FileName }}Service
}

func (c *{{ .FileNameTitleLower }}Controller) Add(ctx *fiber.Ctx) error {
	return middleware.ResponseSuccess(ctx, shared.RespSuccess, nil)
}
