## Channel

- Channel 을 생성한다.
- Channel 에는 값을 넣고 꺼내올 수 있다.
  ```go
  chan <- [v]
  <-chan
  ```
- 값을 넣을 때는 `Goroutine` 을 사용하거나 생성시 넣을 값의 갯수를 추가한다.
