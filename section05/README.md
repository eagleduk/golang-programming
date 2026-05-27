## exercises 1

- `package main` , `func main() {}` 필수
- `:=` 단축 선언
- 출력을 하고자 하면 `"fmt"` 패키지 import

## exercises 2

- 변수에 타입만 선언하면 컴파일러가 값을 지정하는데 이를 `제로값(Zero Value)` 이라 한다.

## exercises 3

- `fmt.Sprintf(format string, a ...any)` format 을 지정하여 문자열을 반환한다.

## exercises 4

- [underlying types](https://go.dev/ref/spec#Underlying_types)
- 사용자 임의의 타입을 만들 수 있다.

## exercises 5

- 원본 타입이 같더라도, `underlying types` 값을 와 원시 타입의 변수에 대입할 수 없다.

## Quiz

- 단축 선언 연산자는 `특정 코드 블록 내에서만 사용. 패키지 레벨에서 사용 불가`
- 리턴값을 사용하지 않을 것이라고 컴파일러에게 알려주는 문자 `_`
