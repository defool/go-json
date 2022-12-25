# go-json
This is a temporary library forked from [go-json](https://github.com/goccy/go-json) to compatible with grpc-gateway([protojson](https://pkg.go.dev/google.golang.org/protobuf/encoding/protojson)).

## Compatible Feature

- Default field JSON tag naming from CamelCase to camelCase.
- Encoding nil slice from `null` to `[]`.

## Example
```
package main

import (
	"fmt"

	"github.com/defool/go-json"
)

type Example struct {
	FieldName string
	Slice     []string
}

func main() {
	val := &Example{FieldName: "111", Slice: nil}
	bs, err := json.Marshal(val)
	fmt.Println("out is", string(bs), err) // out is {"fieldName":"111","slice":[]}

	val2 := &Example{}
	json.Unmarshal(bs, val2)
	fmt.Println("data is", val2) // data is &{111 []}
}
```