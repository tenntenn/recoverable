package recoverable_test

import (
	"fmt"

	. "github.com/tenntenn/recoverable"
	"golang.org/x/sync/errgroup"
)

func ExampleFunc() {
	f := Func(func() {
		panic("example")
	})

	var eg errgroup.Group
	eg.Go(f)

	if err := eg.Wait(); err != nil {
		v, ok := Recovered(err)
		if ok {
			fmt.Println("Panic with", v)
		}
	}

	// Output: Panic with example
}

func ExampleCallStack() {
	f := Func(func() {
		panic("example")
	})

	var eg errgroup.Group
	eg.Go(f)

	if err := eg.Wait(); err != nil {
		callstack := CallStack(err)
		if len(callstack) >= 1 {
			fmt.Println(callstack[0].Func().Name())
		}
	}

	// Output: github.com/tenntenn/recoverable_test.ExampleCallStack.func1
}
