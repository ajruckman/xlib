package xlib

import (
    "fmt"
    "strings"
)

func PrintIndent(indent int, v ...interface{}) {
    fmt.Print(strings.Repeat("\t", indent))
    fmt.Println(v...)
}

func Printfln(format string, a ...interface{}) {
    fmt.Printf(format+"\n", a)
}
