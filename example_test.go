package recoverable_test

import (
	"fmt"

	. "github.com/tenntenn/recoverable"
	"golang.org/x/sync/errgroup"
)

func ExampleFunc() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Panic:", r)
		}
	}()

	f := Func(func() {
		panic("example")
	})

	var eg errgroup.Group
	eg.Go(f)

	if err := eg.Wait(); err != nil {
		v, ok := Recovered(err)
		if ok {
			panic(v)
		}
		fmt.Println("Error:", err)
	}

	// Output: Panic: example
}
