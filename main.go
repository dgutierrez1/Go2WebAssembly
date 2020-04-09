package main

import (
	"syscall/js"
)

func add(this js.Value, i []js.Value) interface{} {
	result := js.ValueOf(i[0].Int() + i[1].Int())
	js.Global().Set("output", result)
	println(result.String())
	return result
}

func subtract(this js.Value, i []js.Value) interface{} {
	result := js.ValueOf(i[0].Int() - i[1].Int())
	js.Global().Set("output", result)
	println(result.String())
	return result
}

func registerCallbacks() {
	js.Global().Set("add", js.FuncOf(add))
	js.Global().Set("subtract", js.FuncOf(su))
}

func main() {
	c := make(chan struct{}, 0)

	println("Go-WebAssembly Initialized")
	registerCallbacks()

	<-c
}
