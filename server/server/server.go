package server

import (
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/AI-Play/Chatting/server/user_manage"
)

type Server struct {
	socket net.Listener
	// address        net.IPAddr
	gerneratedTime time.Time
	users          *user_manage.Users
}

func NewServer(users *user_manage.Users) *Server {
	return &Server{users: users}
}

func (s *Server) ServerStart(wg *sync.WaitGroup, network, address string) {
	defer wg.Done()
	l, err := net.Listen(network, address)
	if err != nil {
		fmt.Println("Listen 중 에러 발생! :", err)
	} else {
		s.socket = l
		// s.address =
		s.gerneratedTime = time.Now()
		fmt.Println("[SERVER] 서버 가동 중!")
	}
	defer l.Close()

	for {
		fmt.Println("[SERVER] 클라이언트 연결 받는 중")
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Accept 중 에러 발생! : ", err)
		} else {
			// conn 정보를 필드로 갖는 유저 객체를 생성해서
			// handler에 전달할 예정
			newUser := user_manage.NewUser(conn, s.users)
			s.users.UserList = append(s.users.UserList, newUser)
			go newUser.UserHandler()
			fmt.Println("유저 접속: ", conn)
		}
	}

}

// func handler(conn net.Conn) {
// 	recv := make([]byte, 4096)

// 	for {
// 		n, err := conn.Read(recv)
// 		if err != nil {
// 			if err == io.EOF {
// 				fmt.Println("connection is closed from client : ", conn.RemoteAddr().String())
// 			}
// 			fmt.Println("Failed to receive data : ", err)
// 			break
// 		}

// 		if n > 0 {
// 			fmt.Println(string(recv[:n]))
// 			conn.Write(recv[:n])
// 		}
// 	}
// }

// func getIP() {
// 	ifaces, err := net.Interfaces()
// 	if err != nil {
// 		fmt.Println("IP 획득 중 에러 발생! net.Interface", err)
// 		return
// 	}
// 	for _, i := range ifaces {
// 		addrs, err := i.Addrs()
// 		// handle err
// 		for _, addr := range addrs {
// 			var ip net.IP
// 			switch v := addr.(type) {
// 			case *net.IPNet:
// 				ip = v.IP
// 			case *net.IPAddr:
// 				ip = v.IP
// 			}
// 			// process IP address
// 		}
// 	}
// }
