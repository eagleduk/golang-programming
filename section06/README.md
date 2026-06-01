## Variable

- 변수를 전언할 때, 한번에 선언할 수 있다.

```go

var (
    a = 11
    b = 1.232
    c = "hello"
)

```

# iota

- 변수를 그룹으로 선언할 때, 증분되는 임의의 상수의 예약어
- 상수에서만 사용 가능하다.

```go
const (
	a = iota
	b = iota
	c = iota
)

const (
	aa = iota
	bb = iota
	cc = iota
)

func main() {
	fmt.Println(a)  // 0
	fmt.Println(b)  // 1
	fmt.Println(c)  // 2

	fmt.Println(aa) // 0
	fmt.Println(bb) // 1
	fmt.Println(cc) // 2
}

```
