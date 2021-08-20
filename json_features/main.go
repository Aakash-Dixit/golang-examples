package main

import (
	"examples/json_features/marshal"
	"examples/json_features/parse"
	"examples/json_features/patch"
)

func main() {
	parse.Parse()
	marshal.Marshal()
	patch.Patch()
}
