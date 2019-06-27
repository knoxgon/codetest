# codetest

### Install iban library

```go
go get -u github.com/knoxgon/codetest/ibanpkg
```

### Example

```go
import (
  "fmt"
  
  "github.com/knoxgon/codetest/ibanpkg"
)

...

isCorrect := ibanpkg.ControlIban("DE89 3704 0044 0532 0130 00")
fmt.Println(isCorrect) //<- True

isFalse := ibanpkg.ControlIban("XB12 33F4 0044 0532 0130 00")
fmt.Println(isFalse) //<- True
```
