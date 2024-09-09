package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
}

func main() {
	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			http.Error(w, "Error upgrading connection", http.StatusInternalServerError)
			return
		}

		for {
			msgType, msg, err := conn.ReadMessage()
			if err != nil {
				http.Error(w, "Error reading message", http.StatusInternalServerError)
			}

			fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))

			time.Sleep(time.Second)

			if err = conn.WriteMessage(msgType, msg); err != nil {
				return
			}
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "websockets.html")
	})

	http.ListenAndServe(":80", nil)
}