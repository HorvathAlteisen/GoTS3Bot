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
	title     string            // Title of the Website/app is stored in here
	templates template.Template // Templates are stored in here
}

type Page struct {
	titel string
	body  []byte
}

// Initializes a new webApp struct, initialiezs path-handlers and sets path from the
func NewWebApp(title string, pathTemplates string, serveMux *http.ServeMux) (*webApp, error) {

	app := new(webApp)
	app.title = title
	//app.templates := template.Must(template.ParseGlob(filepath.Join(pathTemplates, "*.html")))
	filepath.Walk("templates/", func(path string, info os.FileInfo, err error) error {

		// Check if the path is pointing to a file with the ending *.html
		if strings.HasSuffix(path, ".html") {

			fmt.Println(path)

		}
		//fmt.Println(info)

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
