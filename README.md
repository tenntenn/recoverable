# recoverable

[![pkg.go.dev][gopkg-badge]][gopkg]

`recoverable` recovers a panic and convert to an error.

```go
func example() {
	f := recoverable.Func(func() {
		panic("example")
	})

	if err := f(); err != nil {
		v, ok := recoverable.Recovered(err)
		if ok {
			fmt.Println("Panic with", v)
		}
	}
}
```

<!-- links -->
[gopkg]: https://pkg.go.dev/github.com/tenntenn/recoverable
[gopkg-badge]: https://pkg.go.dev/badge/github.com/tenntenn/recoverable?status.svg
