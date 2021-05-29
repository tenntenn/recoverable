package recoverable

import (
	"fmt"
	"runtime"
)

type Caller struct {
	PC   uintptr
	File string
	Line int
}

func callstack() []*Caller {
	var callers []*Caller
	for skip := 3; ; skip++ {
		pc, file, line, ok := runtime.Caller(skip)
		if !ok {
			break
		}

		callers = append(callers, &Caller{
			PC:   pc,
			File: file,
			Line: line,
		})
	}
	return callers
}

func (c *Caller) Func() *runtime.Func {
	return runtime.FuncForPC(c.PC)
}

type errRecovered struct {
	value     interface{}
	callstack []*Caller
}

func (err *errRecovered) Error() string {
	return fmt.Sprintf("panic with %s", err.value)
}

func (err *errRecovered) Recovered() interface{} {
	return err.value
}

func (err *errRecovered) CallStack() []*Caller {
	return err.callstack
}

// RecoveredValue returns recovered value from the error.
// If the error implements bellow interface,
// RecoveredValue returns the recovered value and true.
//     interface {
//          Recovered() interface{}
//     }
func Recovered(err error) (interface{}, bool) {
	rerr, ok := err.(interface {
		Recovered() interface{}
	})

	if !ok {
		return nil, false
	}
	return rerr.Recovered(), true
}

// CallStack returns a call stack of paniced function.
// If CallStack the error implements bellow interface,
//  returns the recovered a call stack.
//     interface {
//          CallStack []*Caller
//     }
func CallStack(err error) []*Caller {
	rerr, ok := err.(interface {
		CallStack() []*Caller
	})

	if ok {
		return rerr.CallStack()
	}

	return nil
}

// Func converts the given function to a function
// which returns an error when a panic
// have occured in the given function.
// The recovered value can get from the error with RecoveredValue.
func Func(f func()) func() error {
	return func() (rerr error) {
		defer func() {
			if r := recover(); r != nil {
				rerr = &errRecovered{value: r, callstack: callstack()}
			}
		}()
		f()
		return nil
	}
}

// Func converts the given function to a function
// which returns an error when a panic
// have occured in the given function.
// The recovered value can get from the error with RecoveredValue.
func FuncWithErr(f func() error) func() error {
	return func() (rerr error) {
		defer func() {
			if r := recover(); r != nil {
				rerr = &errRecovered{value: r, callstack: callstack()}
			}
		}()
		if err := f(); err != nil {
			return err
		}
		return nil
	}
}
