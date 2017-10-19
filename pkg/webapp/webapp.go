package webapp

import (
	"html/template"
	"net/http"
	"path/filepath"
)

type webApp struct {
	title     string            // Title of the Website/app is stored in here
	templates template.Template // Templates are stored in here
}

type Page struct {
	titel string
	body  []byte
}

func NewWebApp(title string, pathTemplates string, serveMux *http.ServeMux) (*webApp, error) {

	app := new(webApp)
	app.title = title
	app.templates := template.Must(template.ParseGlob(filepath.Join(pathTemplates, "*.html")))

	return app, nil
}
