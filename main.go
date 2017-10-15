package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		// Parsing Form to access it with r.Form
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
			log.Output(1, "logfile")

			return
		}

		// Routing Here
		//fmt.Fprintf(w, r.URL.Path[0:])
		if len(r.URL.Path) == 1 {
			http.Redirect(w, r, "/login", http.StatusFound)

			return
		}

		// Template are parsed here
		t, err := template.ParseFiles("templates/index/index.html")
		if err != nil {
			log.Println("executing template:", err)
		}

		// Template get executed here and send as a response to the client
		t.ExecuteTemplate(w, "index.html", nil)

		return
	})
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("templates/css"))))
	http.ListenAndServe(":8080" /*http.FileServer(http.Dir("/templates/"))*/, nil)

	/*query, _ := ts3.NewQuery("127.0.0.1:10011")
	defer query.Close()

	fmt.Println(query.WelcomeMsg)
	fmt.Println(query.SendCommand("use 1"))
	fmt.Println(query.SendCommand("login serveradmin ISkRQGtl"))
	fmt.Println(query.SendCommand("servernotifyregister event=server"))
	fmt.Println(query.SendCommand("servernotifyregister event=channel"))
	fmt.Println(query.SendCommand("servernotifyregister event=textserver"))
	fmt.Println(query.SendCommand("servernotifyregister event=textchannel"))
	fmt.Println(query.SendCommand("servernotifyregister event=textprivate"))

	for {

	}*/

	return
}
