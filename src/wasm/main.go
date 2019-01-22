package main

import (
	"fmt"
	"syscall/js"
	"time"
)

func main() {
	fmt.Println("loading")
	js.Global().Set("add", js.FuncOf(add))
	wait()
}

func wait() {
	done := make(chan bool)
	js.Global().Get("process").Call("on", "SIGTERM", js.FuncOf(func(js.Value, []js.Value) interface{} {
		done <- true
		return nil
	}))
	for {
		select {
		case <-done:
			return
		default:
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func add(this js.Value, args []js.Value) interface{} {
	return js.ValueOf(args[0].Int() - args[1].Int())
}
