package main

import (
	"fmt"
	"sync"

	"github.com/AI-Play/Chatting/server/server"
)

var (
	wg sync.WaitGroup
	// mutex sync.Mutex
)

func main() {
	fmt.Println("[SERVER] 서버 생성")
	mainserver := server.NewServer()

	wg.Add(1)
	fmt.Println("[SERVER] 서버 시작 중")
	go mainserver.ServerStart(&wg)

	wg.Wait()
}
