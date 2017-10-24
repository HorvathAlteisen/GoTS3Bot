package view

import "html/template"

// Views - majority of the HTML must only be in views (templates)
type view struct {
	templates []template.Template
}

func Configure() error {

	return nil
}

func LoadTemplates(templateDir string) error {

	return nil
}
