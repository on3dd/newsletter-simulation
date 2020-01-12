package main

import (
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

const (
	writeWait = 10 * time.Second
	pongWait = 60 * time.Second
	pingPeriod = (pongWait * 9) / 10
	maxMessageSize = 512
)

var upgrader = websocket.Upgrader {
	ReadBufferSize:    1024,
	WriteBufferSize:   1024,
}

type Client struct {
	send <-chan []byte
	conn *websocket.Conn
}

func serveWs(send <-chan []byte, w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	var err error
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatalf("Error upgrading ws connection, error: %v", err)
	}

	client := &Client{send:send, conn:conn}

	log.Println("Client Connected")
	//if err := conn.WriteMessage(1, []byte("Hi Client!")); err != nil {
	//	log.Printf("Cannot write message to ws connection, error: %v", err)
	//}

	go client.writePump(conn)
}

func (c Client) writePump(conn *websocket.Conn) {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		conn.Close()
	}()
	for {
		select {
		case post, ok := <-c.send:
			conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				log.Println("Closing channel")
				if err := conn.WriteMessage(websocket.CloseMessage, []byte{}); err != nil {
					log.Printf("Error writing close message, error: %v", err)
				}
				return
			}

			w, err := conn.NextWriter(websocket.TextMessage)
			if err != nil {
				log.Printf("Cannot select next writer, error: %v\n", err)
			}

			if _, err := w.Write(post); err != nil {
				log.Printf("Cannot write post, error: %v\n", err)
			}

			n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write(<-c.send)
			}

			if err := w.Close(); err != nil {
				log.Printf("Cannot close writer, error: %v\n", err)
			}
		case <-ticker.C:
			conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				log.Printf("Ticker time is up, error: %v", err)
			}
		}
	}
}