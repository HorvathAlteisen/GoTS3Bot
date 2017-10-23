package webapp

import (
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

	filepath.Walk("templates/", func(path string, info os.FileInfo, err error) error {

		// Check if the path is pointing to a file with the ending *.html
		if strings.HasSuffix(path, ".html") {

			fmt.Println(path)
			if app.templates == nil {
				app.templates = template.Must(template.ParseFiles(path))

				return nil
			}

			app.templates = template.Must(app.templates.ParseFiles(path))
			fmt.Println(app.templates.DefinedTemplates())

			return nil

		}

		return nil
	})

	//arr, err := filepath.Glob("templates/*/*.html")
	/*if err != nil {
		log.Println(err)
	}

	for _, k := range arr {
		fmt.Println(k)
	}*/

	return app, nil
}
