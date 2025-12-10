package TemplateDir

import "embed"

//go:embed *.yml *.md docker/* .env Makefile
var TemplateFS embed.FS
