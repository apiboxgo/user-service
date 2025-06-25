package utils

import (
	"fmt"
	"runtime"
)

func Dump(expression ...any) {

	_, file, line, ok := runtime.Caller(1)
	for _, expression := range expression {

		fmt.Println(fmt.Sprintf("\n\n==========\n%#v \n------------ \nfile: %s \nline:%d \nisOk:%t\n\n==========\n\n", expression, file, line, ok))
	}
}
