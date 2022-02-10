package user_manage

import (
	"fmt"
	"io"
	"net"
	"time"
)

var id int = 0

type User struct {
	conn           net.Conn
	gerneratedTime time.Time
	Id             int
	Name           string
	users          *Users
}

func NewUser(conn net.Conn, users *Users) *User {
	newUser := &User{conn: conn, gerneratedTime: time.Now(), Id: id, users: users}
	id++
	return newUser
}
func (u User) String() string {
	return fmt.Sprintf("%d: %s (생성일: %v)", u.Id, u.Name, u.gerneratedTime)
}

// 유저 고루틴을 도는 함수
func (u *User) UserHandler() {
	recv := make([]byte, 4096)
	for {
		n, err := u.conn.Read(recv)
		if err != nil {
			if err == io.EOF {
				fmt.Println("connection is closed from client : ", u.conn.RemoteAddr().String())
				// 유저 객체 정보 db에 저장 후 삭제를 수행할 공간 //
				break
			}
			fmt.Println("Failed to receive data : ", err)
		}
		if n > 0 {
			// 글자가 들어오면 어떤 동작을 수행함.
			// Mux 함수를 만들어서 커맨드에 따라 동작을 수행하도록 하거나
			// SendAll을 불러서 수행하면 될듯.
			if recv[0] == byte('/') {
				//Call commandMux(recv[1])
				//만약 첫 글자가 /라면 명령어가 들어오는 것이다.
				//명령어를 commandMux로 넘겨서 함수를 동작시킬 수 있다.
				u.commandMux(fmt.Sprintln(u.conn.RemoteAddr().String(), ":", string(recv[:n])))

			} else {
				str := fmt.Sprintln(u.conn.RemoteAddr().String(), ":", string(recv[:n]))
				fmt.Println(str)
				u.conn.Write([]byte(str))
			}
		}
	}
}
func (u *User) Send(msg string) { u.conn.Write([]byte(msg)) }
func (u *User) commandMux(msg string) []error {
	switch msg[1] {
	case 'a':
		fmt.Println(msg)
		errs := u.users.SendAll(msg) // 유저 전체에게 메시지 전송
		if len(errs) != 0 {
			fmt.Println(errs)
		}
		return errs
	default:
		fmt.Println("올바르지 않은 명령어")
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
	bMsg := []byte(msg)
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
