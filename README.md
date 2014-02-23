goQPR
=====

Quoted printable encoder for Go


Usage
=====
```
import (
    "fmt"
    "github.com/turretIO/goQPR"
)

func main() {
    encoder := qpr.NewQPEncoder()
    enc, err := encoder.Encode([]byte("encode this to quoted printable"))
    if err != nil {
        fmt.Println(err)
    }    
    fmt.Println(enc)
}
```


