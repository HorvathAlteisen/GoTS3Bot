package webapp

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/HorvathAlteisen/GoTS3Bot/pkg/webapp/router"
)

type webApp struct {
	title     string             // Title of the Website/app is stored in here
	templates *template.Template // Templates are stored in here
	router    *router.Router
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
func NewWebApp(title string, pathTemplates string) (*webApp, error) {

	app := new(webApp)
	app.title = title

	// Checks if the delivered path is existent
	if _, err := os.Stat(pathTemplates); os.IsNotExist(err) {

		return nil, errors.New("webapp: The template path is empty or non-existent")

	}

	app.router = router.New(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/login", 302)
	})

	app.router.Use(router.NewStatic("static/css", "static/js"))

	app.router.Handle("GET", "/login", func(w http.ResponseWriter, r *http.Request) {
		app.templates.ExecuteTemplate(w, "login", struct{ Title string }{Title: app.title})
		fmt.Println("GET-Method")
		r.ParseForm()
		fmt.Println(r.Form)
		for k, v := range r.Form {
			fmt.Println("key:", k)
			fmt.Println("val:", strings.Join(v, ""))
		}

	})

	app.router.Handle("POST", "/login", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		fmt.Println("POST-Method")
		fmt.Println(r.Form)
		for k, v := range r.Form {
			fmt.Println("key:", k)
			fmt.Println("val:", strings.Join(v, ""))
		}

		http.Redirect(w, r, "/index", 302)

	})

	app.router.GET("/index", func(w http.ResponseWriter, r *http.Request) {
		app.templates.ExecuteTemplate(w, "index", struct{ Title string }{Title: app.title})
	})

	filepath.Walk(pathTemplates, func(path string, info os.FileInfo, err error) error {

		//fmt.Println(path)

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

func (app *webApp) Run() error {
	http.ListenAndServe(":8080", app.router)
	return nil
}
