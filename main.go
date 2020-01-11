package main

import (
	"encoding/json"
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

	send := make(chan []byte, 5)

	router.GET("/ws", gin.WrapF(func(w http.ResponseWriter, r *http.Request) {
		serveWs(send, w, r)
	}))

	server := &http.Server{
		Handler:      router,
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.SetFormatter(&log.JSONFormatter{})

	scheduler.Every(5).Seconds().Run(func() {
		post, err := json.Marshal(faker.NewPost())
		if err != nil {
			log.Fatalf("Error marshalling new post to JSON, error: %v", err)
		}

		go func() {
			defer log.Println(len(send))
			send <- post
		}()
	})

	log.Printf("Server started at port %v\n", server.Addr)
	log.Fatal(server.ListenAndServe())
}
