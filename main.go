package main

import (
	"github.com/HorvathAlteisen/GoTS3Bot/TS3ServerQuery"
)

func main() {

	query, _ := TS3ServerQuery.NewQuery("127.0.0.1")

	defer query.Close()

	return
}
