package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/AI-Play/Chatting/server/server"
	"github.com/AI-Play/Chatting/server/user_manage"
)

var (
	wg sync.WaitGroup
	// mutex sync.Mutex
	users *user_manage.Users = &user_manage.Users{}
)

const (
	network string = "tcp"
	address string = "0.0.0.0:5000"
)

func main() {
	fmt.Println("[SERVER] 서버 생성")
	mainserver := server.NewServer(users)

	// wg.Add(1)
	fmt.Println("[SERVER] 서버 시작 중, Network: ", network)
	go mainserver.ServerStart(&wg, network, address)

	// wg.Wait()

	// 전체 메시지 보내기 테스트 함수
	bScan := bufio.NewScanner(os.Stdin)
	str := ""
	for {
		tf := bScan.Scan() // 입력 받는 함수
		if !tf {
			fmt.Println("데이터 입력 중 에러 발생! :")
			bufio.NewReader(os.Stdin).ReadString('\n')
		} else {
			// fmt.Println(str, n)
			str = bScan.Text() // 입력 받은 값을 string으로 저장
			str = fmt.Sprintln("[SERVER]", str)
			errs := users.SendAll(str) // 유저 전체에게 메시지 전송
			if len(errs) != 0 {
				fmt.Println(errs)
			}
			time.Sleep(1 * time.Millisecond) // 과부하 방지를 위한 1ms 휴식
			// users.Ch <- str                  // ch 채널로 string data 전달
		}
	}
}
