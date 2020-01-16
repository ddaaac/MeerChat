# MeerChat
Chatting Shell with Golang

## 하나 목표
각자 *브랜치*를 파서 작동하는 채팅(서버) 구현하기.
- 다음모임때 코드 리뷰를 해서 채팅서버 통합하기.
- P2P던, 채팅방만들기던 메세지를 주고받을 수 있어야 함.
- 터미널이던 웹이던 뭐든 작동해야함.

다음 모임 날짜
- 21일 화요일 저녁 6시에 모여서 같이 밥먹고 합시다.

## 둘 목표
Shell CUI 구현하기. shell 명령어 되면서 상단/우측/좌측/하단/ 임의의 그림이 보이게
- 다다음모임때 코드 리뷰해서 통합하기

## 구현할 기능
### Server side
- [x] 주소:포트가 주어지면 그에 따른 웹소켓 서버 개설 기능
- [x] URL로부터 Query parameter를 읽어오는 기능(id, password)
- [x] 받아온 id와 password로 room을 만들거나 가져오는 기능
    - [x] 이미 존재하는 방이라면 password가 올바른지 체크하는 기능
    - [x] 패스워드가 틀리다면 모든 메세지는 meer로 보이는 기능
    - [x] 패스워드가 틀리다면 보내는 메세지가 다른 사람에게 meer로 보이는 기능
- [x] client로부터 메세지를 읽어오는 기능
- [x] 받은 메세지를 해당 room 안의 다른 client에게 뿌리는 기능
- [x] client와 연결이 끊어졌다면 해당 client를 초기화하는 기능
- [x] room에 client가 없다면 room을 초기화하는 기능
### Client side
- [x] 주소:포트, id, password가 주어지면 그에 따른 방에 웹소켓 연결 기능
- [x] 콘솔 창에 메세지를 입력하여 전송하는 기능
- [x] 콘솔 창에 room의 다른 client가 보낸 메세지를 받는 기능
- [ ] room을 선택하는 기능