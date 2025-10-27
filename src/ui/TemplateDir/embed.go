package TemplateDir

import "embed"

//go:embed *.yml *.md docker/* .env  
var TemplateFS embed.FS
