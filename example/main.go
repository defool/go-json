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
