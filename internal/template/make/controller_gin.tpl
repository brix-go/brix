package controller

import (
	"github.com/gin-gonic/gin"
	"{{ .ProjectName }}/internal/domain/{{ .FileNameTitleLower }}"
	"{{ .ProjectName }}/shared"
	middleware "{{ .ProjectName }}/middleware/error"
)

func NewController(service {{ .FileNameTitleLower }}.{{ .FileName }}Service) {{ .FileNameTitleLower }}.{{ .FileName }}Controller {
	return &{{ .FileNameTitleLower }}Controller{
		{{ .FileNameTitleLower }}Service: service,
	}
}

type {{ .FileNameTitleLower }}Controller struct {
	{{ .FileNameTitleLower }}Service {{ .FileNameTitleLower }}.{{ .FileName }}Service
}

func (c *{{ .FileNameTitleLower }}Controller) Add(ctx *gin.Context) error {
	middleware.ResponseSuccess(ctx, shared.RespSuccess, res)
    return
}
