# AI-Play Chatting

## Stack

- Golang

## Running the Development Server

```
cd brandnew
go run main.go // Auto-reload disabled
// Or
air // Auto-reload enabled
```

For Air reference : https://github.com/cosmtrek/air

## Directory Structure

```
Chatting
  |-- Dockerfile
  |-- README.md
  |-- brandnew // Directory for the currently deployed chat server
  |   |-- main.go
  |   `-- websocket
  |       |-- client.go
  |       |-- init.go
  |       |-- pool.go
  |       `-- websocket.go
  |-- client // Client part of the previous chat server
  |   |-- client
  |   |   |-- client.go
  |   |   `-- init.go
  |   `-- main
  |       |-- init.go
  |       `-- main.go
  |-- go.mod
  |-- go.sum
  `-- server // Server part of the previous chat server
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
