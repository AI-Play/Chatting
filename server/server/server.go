package server

import (
	"fmt"
	"net"
	"time"

	"github.com/AI-Play/Chatting/server/user_manage"
)

type Server struct {
	socket         net.Listener       // 서버가 사용중인 소켓
	gerneratedTime time.Time          // 서버가 생성된 시간
	users          *user_manage.Users // 전체 유저 객체 리스트
}

func NewServer(users *user_manage.Users) *Server {
	return &Server{users: users}
}

func (s *Server) ServerStart(network, address string) {
	// 시스템 서버 Listen.
	l, err := net.Listen(network, address)
	if err != nil {
		fmt.Println("Listen 중 에러 발생! :", err)
	} else {
		s.socket = l
		s.gerneratedTime = time.Now()
		fmt.Println("[SERVER] 서버 가동 중!")
	}
	defer l.Close() // 함수가 종료되면 시스템에 자원 반납

	for {
		fmt.Println("[SERVER] 클라이언트 연결 받는 중")
		conn, err := l.Accept() // 클라이언트가 연결을 시도할 때까지 대기
		if err != nil {
			fmt.Println("Accept 중 에러 발생! : ", err)
		} else {
			// conn 정보를 필드로 갖는 유저 객체를 생성해서
			// UserHandler 함수로 고루틴 실행
			fmt.Println("[SERVER] 클라이언트 연결 성공!")
			fmt.Println(conn)
			newUser := user_manage.NewUser(conn, s.users)        // 유저 객체 생성
			s.users.UserList = append(s.users.UserList, newUser) // 전체 유저 리스트에 유저 객체 추가
			go newUser.UserHandler()                             // UserHandler 함수 GoRoutine 실행
			fmt.Println(s.users.UserList)
			fmt.Println("유저 접속: ", conn)
		}
	}

}
