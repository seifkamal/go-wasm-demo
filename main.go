package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	c := make(chan struct{}, 0)

	fmt.Println("Go WebAssembly Initialized")

	js.Global().Set("greet", js.FuncOf(greet))

	<-c
}

func greet(this js.Value, args []js.Value) interface{} {
	name := args[0].String()
	return fmt.Sprintf("Hello, %s!", name)
}
