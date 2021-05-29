package recoverable_test

import (
	"fmt"

	. "github.com/tenntenn/recoverable"
)

func ExampleFunc() {
	f := Func(func() {
		panic("example")
	})

	if err := f(); err != nil {
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

	if err := f(); err != nil {
		callstack := CallStack(err)
		if len(callstack) >= 1 {
			fmt.Println(callstack[0].Func().Name())
		}
	}

	// Output: github.com/tenntenn/recoverable_test.ExampleCallStack.func1
}
