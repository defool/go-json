package main

import (
	"fmt"
	"strings"

	"github.com/defool/go-json"
)

type abc struct {
	ValAbc string
	Arr    []string

	Obj *abc
}

// fieldNameCase 默认小驼峰命名方式
func fieldNameCase(name string) string {
	if len(name) == 0 {
		return name
	}
	return strings.ToLower(name[:1]) + name[1:]
}

func main() {
	json.SetDefaultJsonTagNaming(fieldNameCase)
	val := &abc{ValAbc: "111", Arr: nil}
	bs, err := json.MarshalWithOption(val, json.NoNullSlice())
	fmt.Println("out", string(bs), err) // "valAbc":"111","arr":[]}
}
