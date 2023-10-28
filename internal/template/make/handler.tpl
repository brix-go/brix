package {{ .FileNameTitleLower }}

import (
	"github.com/gofiber/fiber/v2"
	"{{ .ProjectName }}/internal/service"
)

func NewController(service {{ .FileNameTitleLower }}.{{ .FileName }}Service) {{ .FileNameTitleLower }}.{{ .FileName }}Controller {
	return &{{ .FileNameTitleLower }}Controller{
		service: service,
	}
}

type {{ .FileNameTitleLower }}Controller struct {
	*Handler
	{{ .FileNameTitleLower }}Service service.{{ .FileName }}Service
}

func (c *userController) Add(ctx *fiber.Ctx) error {
	return middleware.ResponseSuccess(ctx, shared.RespSuccess, nil)
}
