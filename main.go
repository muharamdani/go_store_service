package main

import (
	"github.com/muharamdani/store_service/routes"
)

func main() {
	// Echo instance
	e := routes.Init()
	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
