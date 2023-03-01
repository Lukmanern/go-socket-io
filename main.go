package main

import (
	"log"
	"net/http"

	socketio "github.com/googollee/go-socket.io"
)

func main() {
	server := socketio.NewServer(nil)

	server.OnConnect("/", func(so socketio.Conn) error {
		log.Println("Client connected")

		// Send a message to the client when it connects
		so.Emit("message", "Hello from server")

		return nil
	})

	server.OnDisconnect("/", func(so socketio.Conn, reason string) {
		log.Println("Client disconnected:", reason)
	})

	http.Handle("/socket.io/", server)
	http.Handle("/", http.FileServer(http.Dir("./public")))
	log.Println("Server listening on :3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
