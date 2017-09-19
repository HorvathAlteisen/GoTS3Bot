package main

import "github.com/HorvathAlteisen/GoTS3Bot/pkg/ts3/query"
import "github.com/HorvathAlteisen/GoTS3Bot/pkg/ts3/bot"

func main() {

	query, _ := ts3.NewQuery("127.0.0.1:10011")

	query

	defer query.Close()

	return
}
