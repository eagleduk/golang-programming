## loop

```go
for [init]; [condition]; [post]; {}
```

- ~~go 에는 while 문이 없다.~~ while 키워드를 사용하지 않는 것이지 없는것이 아니다.
- 조건절이 없을 수도 있다.

```go
for [condition] {}
```

## conditional

```go
if [condition] {}
```

#### if

- `두줄 이상의 코드를 한줄로 작성할 때 (;) 를 추가한다.` 따라서 초기화와 상태 검증를 한줄에 작성할 수 있다.

```go
if x := 43; x == 3 {}
```

#### else, else if

```go
if [condition] {}
else if [condition] {}
else [condition] {}
```

#### switch

```go
switch [variable] {}
```

- 조건이 없는 `switch` 를 작성할 수 있다. => `switch true 가 된다.`
- 기본적으로 `일치하는 경우에 대해서만 수행하고 끝` 이다.
- 순차적으로 경우를 이어서 수행하고 싶으면 `fallthrough` 를 추가해주어야 한다.
- 여러가지 경우에 대해서는 `,` 로 구분해 주어야 한다.
