package main

import (
	"main/routes"
	"net/http"
)

func main() {
	routes.LoadRoutes()
	http.ListenAndServe(":8000", nil)
}
