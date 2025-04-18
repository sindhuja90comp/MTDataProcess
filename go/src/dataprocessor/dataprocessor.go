```go
package dataprocessor

func Hello() string {
	return "Hello from dataprocessor"
}
```
* **go/main.go**
```go
package main

import (
	"fmt"
	"MTDataProcess/go/dataprocessor"
)

func main() {
	fmt.Println(dataprocessor.Hello())
}
```
* If this works, the problem is in your original `main.go` or `resultsaver.go`