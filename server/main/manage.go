package main

import (
	"crypto/sha512"
	"fmt"
	"time"
)

// 입력된 명령어가 '/'로 시작되는지 확인하는 함수
func serverCommandChecker(str string) bool {
	for {
		if str[0] == '/' {
			return serverCommandMux(str)
		} else {
			fmt.Println("올바르지 않은 명령어:", str)
			return true
		}
	}
}

// 입력된 명령어 대로 동작하는 Mux 함수
func serverCommandMux(str string) bool {
	switch str[1] {
	case 'a':
		return sendToAllUser(str[3:]) // 접속한 모든 유저에게 메시지 전달
	case 'e':
		return exitServer(str[3:]) // 서버 종료
	default:
		fmt.Println("올바르지 않은 명령어: ", str)
		return true
	}
}

//////////////////////////////////////
// 새로운 함수를 아래에 작성한 다음   //
// serverCommandMux에 case 추가     //
/////////////////////////////////////

// 서버 관리자가 전체 유저에 메시지 보내는 함수
func sendToAllUser(str string) bool {
	str = fmt.Sprintln("[SERVER]", str)
	errs := users.SendAll(str) // 유저 전체에게 메시지 전송
	if len(errs) != 0 {
		fmt.Println(errs)
	}
	time.Sleep(1 * time.Millisecond) // 과부하 방지를 위한 1ms 휴식
	// users.Ch <- str                  // ch 채널로 string data 전달
	return true
}

// 서버 종료 함수
func exitServer(str string) bool {
	// db에서 읽어올 해쉬
	password := sha512.Sum512([]byte("1234"))

	// password가 db에 저장된 해시와 같을 경우 서버 종료
	if password == sha512.Sum512([]byte(str)) {
		return false
	} else {
		fmt.Println("잘못된 비밀번호")
		return true
	}
}
