package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("Dial 중 에러 발생! : ", err)
	}
	defer conn.Close()

	ch := make(chan string)

	go func(c net.Conn, ch chan string) {
		defer close(ch)

		for {
			var send string = <-ch
			_, err = c.Write([]byte(send))
			if err != nil {
				fmt.Println("Failed to write data : ", err)

			}
		}
	}(conn, ch)

	go func(c net.Conn) {
		recv := make([]byte, 4096)

		for {
			n, err := c.Read(recv)
			if err != nil {
				fmt.Println("Failed to Read data : ", err)
				break
			}

			fmt.Println("[RESPONSE]", string(recv[:n]))
		}
	}(conn)

	// main goroutine (입력을 전달한다.)
	bScan := bufio.NewScanner(os.Stdin) // 입력을 받는 스캐너
	var str string                      // 입력 받을 스트링 변수
	for {
		tf := bScan.Scan() // 입력 받는 함수
		if !tf {
			fmt.Println("데이터 입력 중 에러 발생!")
			bufio.NewReader(os.Stdin).ReadString('\n')
		} else {
			// fmt.Println(str, n)
			str = bScan.Text()               // 입력 받은 값을 string으로 저장
			ch <- str                        // ch 채널로 string data 전달
			time.Sleep(1 * time.Millisecond) // 과부하 방지를 위한 1ms 휴식
		}
	}
}
