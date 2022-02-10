package server

import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

type Server struct {
	socket net.Listener
	// address        net.IPAddr
	gerneratedTime time.Time
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) ServerStart(wg *sync.WaitGroup) {
	defer wg.Done()
	l, err := net.Listen("tcp", ":8000")
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
			go handler(conn)
			fmt.Println("유저 접속: ", conn)
		}
	}

}

func handler(conn net.Conn) {
	recv := make([]byte, 4096)

	for {
		n, err := conn.Read(recv)
		if err != nil {
			if err == io.EOF {
				fmt.Println("connection is closed from client : ", conn.RemoteAddr().String())
			}
			fmt.Println("Failed to receive data : ", err)
			break
		}

		if n > 0 {
			fmt.Println(string(recv[:n]))
			conn.Write(recv[:n])
		}
	}
}

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
