package main

import (
	"monkey/core"
	"os"
)


func main() {
	instance := core.Repl {}
	instance.Run(os.Stdin, os.Stdout)
}
