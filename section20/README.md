## 동시성

- 함수 앞에 `go` 를 추가하면 병렬로 실행(?)
  : CPU 갯수가 멀티인데 수행이 안된다?

## 초기화 함수

- `init` 함수를 추가하면 main 함수 수행 전에 수행된다.
- 각각 파일에 존재하고 있어도 에러발생하지 않는다.

## WaitGroup

- waitGroup 에 추가한(`WaitGroup.Add(n)`) 만큼 루틴이 추가된다.
- `WaitGroup.Wait` 메소드로 추가된 루틴이 종료되길 기다린다.
- `WaitGroup.Done` 메소드로 루틴을 1개 종료한다.

## race

- go 루틴끼리 양보하면서 경쟁(?)
- 다중 프로세서에 따른 다중 수행 가능
- 현재 멀티 프로세서 때문인지 `WaitGroup` 을 사용하지 않아도 `go` 함수 수행으로도 가능하다(?)
- `runtime.Gosched() ` go 루틴이 실행하면서 프로세서 양보 한다고 한다.

```command
go run -race [].go
```
