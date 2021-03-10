package main

import (
	"code-be-docudigital/db"
	"code-be-docudigital/route"
)

func main() {
	db.Init()

	e := route.Init()

	e.Logger.Fatal(e.Start(":1323"))

}
