package main

import (
	"net/http"
	"time"

	"newsletter-simulation/faker"

	"github.com/carlescere/scheduler"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func main() {
	router := gin.New()

	router.Use(gin.Logger())

	router.Use(static.Serve("/", static.LocalFile("client/dist", true)))
	router.LoadHTMLGlob("client/dist/*.html")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	server := &http.Server{
		Handler:      router,
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.SetFormatter(&log.JSONFormatter{})

	scheduler.Every(10).Seconds().Run(func() {
		p := faker.NewPost()
		log.WithFields(log.Fields{
			"Post Id":         p.Id,
			"Author name":     p.Author.Name,
			"Author username": p.Author.Username,
			"Posted at":       p.PostedAt,
			"Text":            p.Content.Text,
		}).Info()
	})

	log.Printf("Server started at port %v\n", server.Addr)
	log.Fatal(server.ListenAndServe())
}
