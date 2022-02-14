package user_manage

import (
	"fmt"
	"io"
	"net"
	"time"
)

var id int = 0

type User struct {
	conn           net.Conn  // 클라이언트 연결 인터페이스
	gerneratedTime time.Time // 유저 객체가 생성된 시간
	Id             int       // db에서 PK 읽기
	Name           string    // db에서 Name 읽기
	users          *Users    // User 객체를 포함하는 유저리스트
}

// 유저 객체 생성자
func NewUser(conn net.Conn, users *Users) *User {
	newUser := &User{conn: conn, gerneratedTime: time.Now(), Id: id, users: users}
	id++
	return newUser
}

// 유저 객체 String 메소드. python의 __str__ 메소드 와 같다.
func (u User) String() string {
	return fmt.Sprintf("%d: %s (생성일: %v)", u.Id, u.Name, u.gerneratedTime)
}

// 유저 고루틴을 도는 메소드
func (u *User) UserHandler() {
	recv := make([]byte, 4096)
	for {
		n, err := u.conn.Read(recv)
		if err != nil {
			if err == io.EOF {
				fmt.Println("connection is closed from client : ", u.conn.RemoteAddr().String())
				// 유저 객체 정보 db에 저장 후 삭제를 수행할 공간
				u.conn.Close()                       // 연결을 끊고
				for i, v := range u.users.UserList { // 객체를 삭제한다
					if v == u {
						fmt.Println(i, v, u)
					}
				}
				break
			}
			fmt.Println("Failed to receive data : ", err)
		}
		if n > 0 {
			// 첫 글자 '/'로 들어오면 명령을 수행함.
			if recv[0] == byte('/') {
				u.commandMux(recv[:n])
			} else {
				str := fmt.Sprint(u.Name, ": ", string(recv[:n]))
				fmt.Println(str)
				u.conn.Write([]byte(str))
			}
		}
	}
}
func (u *User) Send(msg string) { u.conn.Write([]byte(msg)) }
func (u *User) commandMux(recv []byte) []error {
	msg := fmt.Sprint(u.Name, ": ", string(recv[3:]))
	switch recv[1] {
	case byte('a'):
		fmt.Println(msg)
		errs := u.users.SendAll(msg) // 유저 전체에게 메시지 전송
		if len(errs) != 0 {
			fmt.Println(errs)
		}
		return errs
	default:
		fmt.Printf("올바르지 않은 명령어: %v\n", string(recv[:1]))
		u.Send("명령어가 올바르지 않습니다.")
		return nil
	}

}

////////////////
// User Slice //
////////////////
type Users struct {
	UserList []*User
	// Ch       chan string
}

func (u *Users) SendAll(msg string) []error {
	bMsg := []byte("[ALL]" + msg)
	errs := make([]error, 0, 10)
	for _, v := range u.UserList {
		n, err := v.conn.Write(bMsg)
		if err != nil {
			fmt.Println(v.Id, " 유저에게 전송 실패", n)
			errs = append(errs, err)
		}
	}
	return errs
}
