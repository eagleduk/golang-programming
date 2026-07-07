## functions

- 함수는 _파라메터(매개변수)_ 와 함께 정의 한다.
- 함수를 사용하고자 할 때에는 정의되어 있는 _인수(값)_ 를 전달한다.
- 값을 반환할 때에는 파이썬처럼 단일 변수로 여러개를 반환할 수 있다.

  ```go

  func 함수명<리시버> ([파라메터명 [타입]]...) 리턴타입 {}

  func foo(s string) (string, bool) {}

  ```

- `...` 기호를 사용하여 갯수가 정해지지 않은 매개변수(**0개 이상**)를 정의할 수 있다.
- 가변 매개변수는 함수의 제일 마지막 매개변수가 되어야 한다.

  ```go
      func foo(s ...int) {}
  ```

## defer

- 함수가 종료될 때 수행하고자 하는 키워드
- 예를 들어, 파일을 여는 명령어 바로 아래에 _defer_ 를 추가한 파일을 닫는 명령러를 작성한다.

  ```go
  func main() {
      foo()
      defer bar()
      foobar()
  }
  // foo() -> foobar() -> bar()
  ```

## Method

- 구조체에 정의되는 함수
- 구조체를 만들고 함수를 연결 하는 식으로 추가한다.
  ```go
  type pserson struct {
      firstname string
      lastname string
  }
  func (a person) fullname() {
      return a.firstname + " " + a.lastname
  }
  ```

## Interface

- 구조체를 메소드로 엮는다.
- 동일한 메소드가 있으면 같은 인터페이스 타입으로 정의한다.
  ```go
  type in interface {
      say()
  }
  ```
