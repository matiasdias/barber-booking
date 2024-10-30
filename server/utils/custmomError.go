package utils

import (
	"fmt"
	"runtime"
)

type CustomError struct {
	Code  int
	Msg   string
	Cause error
	Trace []string
}

func New(code int, msg string, cause error) *CustomError {
	trace := captureStackTrace()
	return &CustomError{Code: code, Msg: msg, Cause: cause, Trace: trace}
}

// Error implementa o método Error() da interface error
func (e *CustomError) Error() string {
	return fmt.Sprintf("%s", e.Msg)
}

// captureStackTrace captura o rastreamento de pilha atual para fins de debug
func captureStackTrace() []string {
	var trace []string
	pc := make([]uintptr, 10)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])

	for frame, more := frames.Next(); more; frame, more = frames.Next() {
		trace = append(trace, fmt.Sprintf("%s:%d %s", frame.File, frame.Line, frame.Function))
	}
	return trace
}
