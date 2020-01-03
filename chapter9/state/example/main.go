package main

import (
	"context"
	"fmt"

	"github.com/jaakidup/go-cookbook/chapter9/state"
)

func main() {
	in := make(chan *state.WorkRequest, 10)
	out := make(chan *state.WorkResponse, 10)
	ctx := context.Background()

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	go state.Processor(ctx, in, out)

	req := &state.WorkRequest{Operation: state.Add, Value1: 3, Value2: 4}
	in <- req

	req2 := &state.WorkRequest{Operation: state.Subtract, Value1: 5, Value2: 2}
	in <- req2

	req3 := &state.WorkRequest{Operation: state.Multiply, Value1: 9, Value2: 9}
	in <- req3

	req4 := &state.WorkRequest{Operation: state.Divide, Value1: 8, Value2: 2}
	in <- req4

	req5 := &state.WorkRequest{Operation: state.Divide, Value1: 8, Value2: 0}
	in <- req5

	for i := 0; i < 5; i++ {
		resp := <-out
		fmt.Printf("Request: %v; Result: %v, Error: %v", resp.Wr, resp.Result, resp.Err)
		println()
	}
}
