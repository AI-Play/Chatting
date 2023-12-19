# AI-Play Chatting

## Stack

- Golang

## 개발 서버 작동 방법

```
cd brandnew
go run main.go // Auto-reload X
// 또는
air // Auto-reload O
```

Air 관련 참고 : https://github.com/cosmtrek/air

## 디렉토리 구조도

```
Chatting
  |-- Dockerfile
  |-- README.md
  |-- brandnew // 현재 배포 중인 채팅 서버의 디렉토리
  |   |-- main.go
  |   `-- websocket
  |       |-- client.go
  |       |-- init.go
  |       |-- pool.go
  |       `-- websocket.go
  |-- client // 이전 채팅 서버의 클라이언트 부분
  |   |-- client
  |   |   |-- client.go
  |   |   `-- init.go
  |   `-- main
  |       |-- init.go
  |       `-- main.go
  |-- go.mod
  |-- go.sum
  `-- server // 이전 채팅 서버의 서버 부분
      |-- main
      |   |-- init.go
      |   |-- main.go
      |   `-- manage.go
      |-- server
      |   |-- init.go
      |   `-- server.go
      `-- user_manage
          |-- init.go
          `-- userobject.go
```
