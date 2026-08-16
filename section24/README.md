## Error

- 타 언어와는 다르게 try,catch 로 예외처리를 하지 않는다
- 에러가 발생하는 부분에서 바로 에러처리를 한다.

## Log

#### fmt.Print

#### Log.Println

- 기본적인 로그 출력
- `fmt.Println` 에서 로그 발생 시간이 앞에 추가된다.

#### Log.Write

- 기본 로그 출력을 파일에 저장한다.

#### Log.Fatalln

- 에러가 발생하면 바로 프로그램을 종료한다.
- 발생 시점 이후의 모든것은 무시한다.

#### Log.Panic

- `goroutine` 이 있다면 역순으로 정지할 때 까지 수행
- `defer` 가 있다면 30분 후에 수행
- `recover` 로 제어 가능