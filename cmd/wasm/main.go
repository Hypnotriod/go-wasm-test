package main

import (
	"errors"
	"fmt"
	"syscall/js"
	"time"
)

func main() {
	js.Global().Set("goWasmTest", js.FuncOf(goWasmTest))
	fmt.Println("Test from Go WASM")
	<-make(chan struct{})
}

func goWasmTest(this js.Value, args []js.Value) interface{} {
	if len(args) != 1 || args[0].Type() != js.TypeFunction {
		return errors.New("Invalid argument")
	}
	go func() {
		time.Sleep(1 * time.Second)
		args[0].Invoke("Message from goroutine delayed by a second")
	}()
	return js.Undefined()
}
