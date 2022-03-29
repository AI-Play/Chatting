package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/AI-Play/Chatting/brandnew/websocket"
)

func serveWs(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
    fmt.Println("WebSocket Endpoint Hit")
    conn, err := websocket.Upgrade(w, r)
    if err != nil {
        fmt.Fprintf(w, "%+v\n", err)
    }

    client := &websocket.Client{
        Conn: conn,
        Pool: pool,
    }

    pool.Register <- client
    client.Read()
}

func setupRoutes() {
    fmt.Println("Initializing Routes")
    pool := websocket.NewPool()
    go pool.Start()

    http.HandleFunc("/wss", func(w http.ResponseWriter, r *http.Request) {
        serveWs(pool, w, r)
    })
}

func main() {
    fmt.Println("AI-Play Public Chat App v0.1")
    setupRoutes()
    port := fmt.Sprintf(":%s", os.Getenv("PORT"))
    http.ListenAndServe(port, nil)
}