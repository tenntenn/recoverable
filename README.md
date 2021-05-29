# recoverable

[![pkg.go.dev][gopkg-badge]][gopkg]

`recoverable` recovers a panic and convert to an error.

```go
func example() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Panic:", r)
		}
	}()

	f := recoverable.Func(func() {
		panic("example")
	})

	var eg errgroup.Group
	eg.Go(f)

	if err := eg.Wait(); err != nil {
		v, ok := recoverable.Recovered(err)
		if ok {
			panic(v)
		}
		fmt.Println("Error:", err)
	}

	// Output: Panic: example
}
```

<!-- links -->
[gopkg]: https://pkg.go.dev/github.com/tenntenn/recoverable
[gopkg-badge]: https://pkg.go.dev/badge/github.com/tenntenn/recoverable?status.svg
