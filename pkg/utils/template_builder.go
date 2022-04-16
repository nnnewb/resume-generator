package utils

import (
	"errors"
	htmltemplate "html/template"
)

type TemplateBuilder struct {
	main  string
	files []string
	funcs map[string]interface{}
}

func NewTemplateBuilder() *TemplateBuilder {
	return &TemplateBuilder{
		files: []string{},
		funcs: map[string]interface{}{},
	}
}

func (b *TemplateBuilder) SetMainTemplate(path string) *TemplateBuilder {
	b.main = path
	return b
}

func (b *TemplateBuilder) AddTemplateFile(path string) *TemplateBuilder {
	// first added template file as main template
	if b.main == "" {
		b.main = path
	}
	b.files = append(b.files, path)
	return b
}

func (b *TemplateBuilder) AddFunc(name string, f interface{}) *TemplateBuilder {
	b.funcs[name] = f
	return b
}

func (b *TemplateBuilder) ExtendFuncs(fm map[string]interface{}) *TemplateBuilder {
	for k, v := range fm {
		b.funcs[k] = v
	}
	return b
}

func (b *TemplateBuilder) BuildHTMLTemplate() (*htmltemplate.Template, error) {
	if b.main == "" {
		return nil, errors.New("main template not specified")
	}

	t := htmltemplate.New(b.main)
	t = t.Funcs(b.funcs)
	return t.ParseFiles(b.files...)
}
