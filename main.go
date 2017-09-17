package main

import "github.com/HorvathAlteisen/GoTS3Bot/ts3"

func main() {

	query, _ := ts3.NewQuery("127.0.0.1:10011")

	bot, _ := ts3.NewBot(query)

	defer bot.Close()

	return
}
