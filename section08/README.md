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
