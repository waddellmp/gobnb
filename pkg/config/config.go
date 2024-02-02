package config

import "html/template"

// AppConfig is struct holding die application config
type AppConfig struct {
	TemplateCache map[string]*template.Template
}
