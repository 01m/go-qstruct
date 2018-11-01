# qstruct
qstruct is Go library for decoding url values into struct

# Usage

```go
import "github.com/01m/qstruct/qstruct"
```

```go
import (
	"fmt"
	"net/url"

	"github.com/01m/qstruct/qstruct"
)

type structA struct {
	A string `json:"a"`
}

func main() {
	u := url.Values{}
	u.Set("a", "123")

	var q = structA{}
	err := qstruct.Decode(u, &q)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", q)
}
```
