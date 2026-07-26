## run

- go 파일을 컴파일 하기 전에 실행한다.
  ```command
  go run [].go
  ```

## build

- go 파일을 build 하면 컴파일된 파일이 생성된다.
  ```command
  go build
  ```
- 빌드한 파일은 바로 실행할 수 있다.
  ```command
  .\[].go
  ```

## install

- build 된 파일을 po path에 설치한다.
  ```command
  go install
  ```
- install 한 파일은 `$path` 에 등록되기 때문에 어디서든 바로 실행할 수 있다.
  ```command
  .\[].go
  ```
