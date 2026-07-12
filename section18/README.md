## json

- `encodeing/json`

#### marshal

- struct type 을 json 형태로 변경해준다.
- json 의 필드명이 _대문자_ 로 시작해야 한다.

#### unmarshal

- json 형태의 byte 를 struct type 으로 변환해 준다.
- struct 를 정의할 때 tag 를 이용해서 원하는 키에 매핑해준다.
- json 을 받을 구조체의 필드명은 무조건 _대문자_ 로 시작해야 한다.
- json 데이터를 받고자 하는 변수의 **주소값** 을 넘겨주어야 한다.

## sort

## bcrypt

- 공식 외부 라이브러리 사용
  ```command
  go mod init [내 프로젝트명]
  go get golang.org/x/crypto/bcrypt
  ```
