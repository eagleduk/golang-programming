## pointer

- 메모리 주소를 확인(`&`)
  ```go
      a := 44
      fmt.Println(a)
      fmt.Println($a)
  ```
- 메모리 주소가 가지고 있는 데이터를 확인(`*`)
  ```go
      a := 44
      fmt.Println(a)
      fmt.Println(*&a)
  ```
