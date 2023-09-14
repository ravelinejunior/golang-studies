package main

import (
	"module/src/website/webproject/routes"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	routes.LoadRoutes()
	http.ListenAndServe(":8080", nil)
}
