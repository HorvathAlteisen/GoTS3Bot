package main

import (
	"github.com/HorvathAlteisen/GoTS3Bot/pkg/webapp"
)

func main() {

	app, err := webapp.NewWebApp("GoTS3Bot", "templates/")
	if err != nil {

		return
	}
	app.Run()

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
