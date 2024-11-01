package main

import (
	"github.com/fatih/color"
	"github.com/ysl2/go-hello-world/pkg/greeter"
	"github.com/ysl2/go-hello-world/utils/helper"
)

func main() {
	message := greeter.Greet()
	helper.PrintMessage(message)
	color.Green("Hello, World!")
}
