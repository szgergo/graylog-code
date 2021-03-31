package util

import (
	"io"
	"com/github/graylog-code/error"
)

func CloseResource(toClose io.ReadCloser) {
	err := toClose.Close()
	error.Check(err)
}
