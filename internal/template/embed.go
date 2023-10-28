package template

import "embed"

//go:embed make/*.tpl
var CreateTemplateFS embed.FS
