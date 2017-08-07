package template

import (
	"time"
	"html/template"
)

var defaultFilters = template.FuncMap{
	"now":    time.Now,
}

