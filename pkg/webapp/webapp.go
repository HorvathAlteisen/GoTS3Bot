package webapp

import (
	"net/http"
)

type webApp struct {
	title     string   // Title of the Website/app is stored in here
	templates Template // Templates are stored in here
}

type Page struct {
	titel string
	body  []byte
}

func newWebApp(title string, pathTemplates string, serveMux *http.ServeMux) (*webApp, error) {

	return
}
