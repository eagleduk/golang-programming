## Testing

- 테스트 하고자 하는 패키지와 동일한 폴더에 위치
- 테스트 하고자 하는 패키지의 `_test.go` 
- `func Test[함수]` 테스트 하고자 하는 함수
- `go test` 로 수행

- 테스트 문서도 일반 API 문서 작성하듯이 작성이 가능하다.

## golint

- 스타일 검사 도구
- **현재는 사용하지 않는 듯**
- `gofmt`: 파일 포멧 정리
- `go vet`: 파일 잠재 오류 분석
- `staticcheck (설치 필요, 서드파티)`: 스타일뿐 아니라 잠재적인 버그, 비효율, 잘못된 API 사용까지 검사

## benchmark

- 작성한 함수의 퍼포먼스를 확인할 수 있다.

    ```zsh
    go test -bench .
    ```

## coverage

- 작성한 함수의 사용량을 알 수 있다.

    ```zsh
    go test -cover
    ```

- 분석 내용을 파일로 내보낸다.

    ```zsh
    go test -coverprofile=c.out
    ```
    
- 내보낸 파일을 html 형식으로 확인한다.

    ```zsh
    go tool cover -html=c.out
    ```
