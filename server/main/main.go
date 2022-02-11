package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/AI-Play/Chatting/server/server"
	"github.com/AI-Play/Chatting/server/user_manage"
)

var (
	// wg sync.WaitGroup
	// mutex sync.Mutex
	users *user_manage.Users = &user_manage.Users{}
	// commandCh                    = make(chan string)
)

const (
	network string = "tcp"
	address string = "0.0.0.0:5000"
)

func main() {
	fmt.Println("[SERVER] 서버 생성")
	mainserver := server.NewServer(users)

	fmt.Println("[SERVER] 서버 시작 중, Network:", network)
	go mainserver.ServerStart(network, address)

	/////////////////////////////////
	// Mux를 통한 명령어로 서버 관리 //
	/////////////////////////////////

	// bufio.Scanner를 통해 서버에 직접 명령어를 입력한다.
	bScan := bufio.NewScanner(os.Stdin)
	for {
		tf := bScan.Scan() // 서버에 명령어 입력: 입력 전까지 대기하는 코드
		if !tf {
			fmt.Println("데이터 입력 중 에러 발생! :")

			// os.Stdin에 남아 있는 버퍼 내용을 비운다.
			bufio.NewReader(os.Stdin).ReadString('\n')
		} else {
			// 정상 입력시 입력 받은 값을 string으로 main/manage.go 함수에 명령어 전달
			tf := serverCommandChecker(bScan.Text())
			if !tf {
				break
			}
		}
	}
}
