package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jaquelineabreu/templ-htmx/internal"
)

func main() {
	router := gin.Default()

	//initialize config
	app := internal.Config{Router: router}

	//routes
	app.Routes()

	router.Run(":8080")

}
