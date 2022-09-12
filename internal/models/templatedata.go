package models

import "github.com/jacstn/golang-btc-dep-rec/internal/forms"

type TemplateData struct {
	Data map[string]interface{}
	Form *forms.Form
}
