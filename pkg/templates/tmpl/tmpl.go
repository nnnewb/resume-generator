package tmpl

import (
	_ "embed"
)

var (
	//go:embed html/index.html.tpl
	HtmlTemplate string
	//go:embed markdown/index.md.tpl
	MarkdownTemplate string
)
