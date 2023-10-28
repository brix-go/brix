package template

import "embed"

//go:embed create/*.tpl
var CreateTemplateFS embed.FS
