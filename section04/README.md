## Control flow

- `package`, `func main()` 가 필수로 있어야 한다.
- 변수를 선언하면 반드시 사용해야 한다.(`_`를 사용하여 컴파일시 무시하라는 명령을 내릴 수 있다.)

## [변수](https://go.dev/ref/spec#Identifiers)

- 변수 선언시 `문자 + 문자 | 유니코드` 형식으로 선언한다.
- `go`에서 사용하는 [예약어](https://go.dev/ref/spec#Keywords)는 사용할 수 없다.
- `var` 변수를 선언하고 값을 지정한다.
  - 함수 밖에서 선언할 떄 사용된다.
  - 값을 선언하지 않으면 기본값(integer: 0, string: "", bool: false, float: 0.0, etc: nil)
    ```go
        var x = 32
        var z int
        func main() {
            fmt.Println(x)
        }
    ```
- `:=` 단축 선언 연산자 변수를 선언하고 값을 지정한다.
  - 함수 내에서 변수를 사용할 떄 사용한다.
    ```go
        func main() {
            x := 22 
            fmt.Println(x)
            x = 33
            fmt.Println(x)
        }
    ```

## [타입](https://go.dev/ref/spec#Types)

- 정적 타입만 선언할 수 있다.(string 타입을 선언하고 int 값을 대입할 수 없다.)
- 원시 자료형(string, int, bool)
- 복합 자료형
- 변수의 타입을 출력할 수 있다.
    ```go
    fmt.Printf("%T", y)
    ```
- 나만의 커스텀 타입을 만들 수 있다
    ```go
        type abc int
        var a abc

        func main {
            fmt.Printf("%T", a)     // main.abc
        }
    ```
- 타입 변환(치환): ***TYPE(변수)*** 함수로 변환할 수 있다.
    ```go
    int(a)
    string(b)
    ```