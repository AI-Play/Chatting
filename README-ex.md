# Chatting

시작하기 전에 고려해볼 사항...
1. 에러 핸들링(관련 패키지 제작 고려...)
2. 채팅 log에 대한 처리. 
    client가 채팅 log 서버를 보는 형태 vs 서버에서 클라이언트에게 채팅 정보를 날리는 형태
3. 전체적인 flow chart 그리기


대략적인 구조도

main server(일단 AWS EC2) 

1. main thread
  -> 서버를 바인딩 한다.
  
  -> listen thread 를 생성한다.
  
  -> 명령어로 thread관리, user관리, 채팅창 관리 등등을 수행한다. -> 각각 관리하는 프로그램 또는 패키지를 짜야할듯?

2. listen thread(1개)
  -> 스레드 하나를 써서 접속자를 확인하는 루프를 돈다
  
  -> 접속자의 정상 접근이 확인되면 user thread를 생성한다 -> 유저 스레드로
  
  -> 다른 접속자가 확인될 때 까지 반복한다
  
  -> 서버가 종료될 때 같이 종료한다.
  
3. user thread(개수 많음)
  -> 접속자의 정보를 받아서 db에 있는지 확인하고, 있을 경우 db 정보를 load, 없으면 새로운 객체 생성

  -> 유저로 부터 데이터를 받는 루프를 돈다
  
  -> 데이터가 들어오면 해당 명령어에 맞는 행동을 한다
  
  -> 해당 유저의 연결이 끊어지면 유저 객체 정보를 db에 저장하고 객체를 삭제한다.
  
4. (Optional) chat room thread(어느정도 가능)
  -> 유저로 부터 받은 명령중에 new chat room 명령을 받아서 생성한다.
  
  -> 접속 유저가 0 명이면 스레드를 종료하고 해당 chat room 객체를 삭제한다.
  

```
Chatting/
  │
  ├── server/
  │     │
  │     ├── main/
  │     │     ├── init.go   : 서버 시작에 필요한 initialize(미구현)
  │     │     └── main.go   : Entry Point Server 객체를 생성하고 GoRoutine을 생성한다.
  │     │     └── manage.go : 명령어 입력 및 실시간 서버 관리 함수
  │     │
  │     ├── server/
  │     │     ├── init.go   : 임포트에 필요한 initialize(미구현)
  │     │     └── server.go : 서버 객체, 서버 구동 함수, 유저 객체를 생성 GoRoutine을 생성한다.
  │     │
  │     └── user_manage/
  │           ├── init.go       : 임포트에 필요한 initialize(미구현)
  │           └── userobject.go : 유저 객체 함수. 유저 커맨드를 받아서 동작하는 함수를 관리한다.
  │
  └── client/
        │
        ├── main/
        │     ├── init.go : 클라이언트 시작에 필요한 initialize(미구현)
        │     └── main.go : 서버 연결 고루틴 실행 및 명령어 입력 관리
        │     
        └── client/
              ├── init.go   : 임포트에 필요한 initialize(미구현)
              └── client.go : 클라이언트 함수 및 객체 관리(미구현)
```