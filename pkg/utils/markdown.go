package utils

import (
	"bytes"
	"html/template"

	"github.com/pkg/errors"
	"github.com/yuin/goldmark"
)

func Markdown(source string) string {
	buffer := bytes.NewBufferString("")
	err := goldmark.Convert([]byte(source), buffer)
	if err != nil {
		panic(errors.Wrap(err, "error occurred when rendering markdown"))
	}
	return buffer.String()
}

func Unescape(source string) template.HTML {
	return template.HTML(source)
}
