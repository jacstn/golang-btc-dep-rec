package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/jacstn/golang-btc-dep-rec/config"
	"github.com/jacstn/golang-btc-dep-rec/internal/models"
	"github.com/justinas/nosurf"
)

var app *config.AppConfig

func NewHandlers(c *config.AppConfig) {
	app = c
}

func Home(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	data["csrf_token"] = nosurf.Token(r)

	renderTemplate(w, "home", &models.TemplateData{
		Data: data,
	})
}

func Customers(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	data["csrf_token"] = nosurf.Token(r)
	customers, err := models.ListCustomers(app.DB)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	data["customers"] = customers

	renderTemplate(w, "customers", &models.TemplateData{
		Data: data,
	})
}

func renderTemplate(w http.ResponseWriter, templateName string, data *models.TemplateData) {
	funcMap := template.FuncMap{
		"incr": func(i int) int {
			return i + 1
		},
	}

	t, err := template.New(templateName+".tmpl").Funcs(funcMap).ParseFiles("./templates/"+templateName+".tmpl", "./templates/layout.main.tmpl")

	if err != nil {
		fmt.Fprint(w, "Error parsing template page!!", err)
		return
	}

	err = t.Execute(w, data)

	if err != nil {
		fmt.Fprint(w, "Error handling template page!!", err)
	}
}
