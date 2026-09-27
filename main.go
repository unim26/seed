package main

import (
	"fmt"

	"github.com/unim26/seed/internal/colors"
	"github.com/unim26/seed/internal/seed"
)

func main() {
	msg, err := seed.ProcessCommand()
	if err != nil {
		fmt.Printf("%s%v%s", colors.Red, err, colors.Reset)
		return
	}

	fmt.Printf("%s%s%s", colors.Green, msg, colors.Reset)
}
