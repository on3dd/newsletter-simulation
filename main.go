package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()

	// Middleware
	router.Use(gin.Logger())

	// Frontend
	router.Use(static.Serve("/", static.LocalFile("client/dist", true)))
	router.LoadHTMLGlob("client/dist/*.html")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	server := &http.Server{
		Handler: router,
		Addr: ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout: 15 * time.Second,
	}

	log.Printf("Server started at port %v\n", server.Addr)
	log.Fatal(server.ListenAndServe())
}