package errors

import (
	"fmt"
	"runtime"
)

func F(err error) error {
	_, file, line, _ := runtime.Caller(1)
	return fmt.Errorf("%s:%d %w", file, line, err)
}
