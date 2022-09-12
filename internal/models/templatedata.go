package models

import "github.com/jacstn/golang-btc-deposit-reconciliation/internal/forms"

type TemplateData struct {
	Data map[string]interface{}
	Form *forms.Form
}
