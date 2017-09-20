package main

import "github.com/HorvathAlteisen/GoTS3Bot/pkg/ts3"

func main() {

	query, _ := ts3.NewQuery("127.0.0.1:10011")

	defer query.Close()

	return
}
