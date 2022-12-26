# go-json
This is a temporary library forked from [go-json](https://github.com/goccy/go-json) to compatible with grpc-gateway([protojson](https://pkg.go.dev/google.golang.org/protobuf/encoding/protojson)).

## Compatible Feature

- Support both camelCase and snake_case in default JSON tag 
- Encoding nil slice from `null` to `[]`.
- Auto parse string to number.

## Example
```
package main

import (
	"fmt"

	"github.com/defool/go-json"
)

type Example struct {
	FieldName string
	IntField  uint32
	Slice     []string
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	val := &Example{FieldName: "111", Slice: nil, IntField: 2}
	bs, err := json.Marshal(val)
	checkError(err)
	fmt.Println("out is", string(bs), err) // out is {"fieldName":"111","intField":2,"slice":[]}

	val2 := &Example{}
	bs = []byte(`{"field_name":"111","intField":"22"}`)
	err = json.Unmarshal(bs, val2)
	checkError(err)
	fmt.Printf("data is %+v\n", val2) // data is &{FieldName:111 IntField:22 Slice:[]}
}

```