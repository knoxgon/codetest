# IBAN Verifier

[IBAN Wiki Page](https://github.com/knoxgon/codetest/wiki)

Verifies a given IBAN. It checks by first converting the whole IBAN into integers and calculates the modulo by 97. If the result number is 1, the IBAN might be valid. Any other numbers are not valid.

### Install

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
