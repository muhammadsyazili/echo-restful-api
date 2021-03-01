package main

import (
	"github.com/muhammadsyazili/echo-rest/db"
	"github.com/muhammadsyazili/echo-rest/routes"
)

func main() {
	db.Init()
	e := routes.Init()

	e.Logger.Fatal(e.Start(":1010"))
}