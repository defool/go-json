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

	// buf := bytes.NewBuffer(bs)
	// err = json.NewDecoder(buf).Decode(val2) // unsupported now
	err = json.Unmarshal(bs, val2)
	checkError(err)
	fmt.Printf("data is %+v\n", val2) // data is &{FieldName:111 IntField:22 Slice:[]}
}
