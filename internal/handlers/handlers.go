package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/jacstn/golang-btc-deposit-reconciliation/config"
	"github.com/jacstn/golang-btc-deposit-reconciliation/internal/forms"
	"github.com/jacstn/golang-btc-deposit-reconciliation/internal/models"
	"github.com/justinas/nosurf"
)

var app *config.AppConfig

func NewHandlers(c *config.AppConfig) {
	app = c
}

func Home(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	data["omise_key"] = app.OmisePublicKey
	data["csrf_token"] = nosurf.Token(r)

	renderTemplate(w, "home", &models.TemplateData{
		Form: forms.New(nil),
		Data: data,
	})
}

func renderTemplate(w http.ResponseWriter, templateName string, data *models.TemplateData) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+templateName+".tmpl", "./templates/base.layout.tmpl")

	err := parsedTemplate.Execute(w, data)
	if err != nil {
		fmt.Fprintf(w, "Error handling template page!!", err)
	}
}
