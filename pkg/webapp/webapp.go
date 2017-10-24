package webapp

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type webApp struct {
	title     string             // Title of the Website/app is stored in here
	templates *template.Template // Templates are stored in here
	serveMux  *http.ServeMux
	pages     []page
}

type page struct {
	title   string
	content []byte
	actions []actions
}

type actions struct {
	name       string
	permission string
}

// NewWebApp Initializes a new webApp struct, initialiezs path-handlers and sets path from the
func NewWebApp(title string, pathTemplates string, serveMux *http.ServeMux) (*webApp, error) {

	app := new(webApp)
	app.title = title
	app.serveMux = serveMux

	// Checks if the delivered path is existent
	if _, err := os.Stat(pathTemplates); os.IsNotExist(err) {

		return nil, errors.New("webapp: The template path is empty or non-existent")

	} else if serveMux == nil {

		return nil, errors.New("webapp: The serveMux mustn't be nil")
	}

	filepath.Walk(pathTemplates, func(path string, info os.FileInfo, err error) error {

		fmt.Println(path)

		// Check if the path is pointing to a file with the ending *.html
		if strings.HasSuffix(path, ".html") {

			// First time initialization
			if app.templates == nil {
				app.templates = template.Must(template.ParseFiles(path))

				return err
			}

			app.templates = template.Must(app.templates.ParseFiles(path))
			fmt.Println(app.templates.DefinedTemplates())

			return err

		}

		return err
	})

	return app, nil
}

func (app *webApp) Start() error {
	return app.serveMux.ServeHTTP
}
