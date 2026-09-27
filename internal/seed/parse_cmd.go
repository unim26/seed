package seed

import (
	"errors"
	"fmt"
	"os"

	"github.com/unim26/seed/internal/colors"
	"github.com/unim26/seed/internal/services"
)

// process coomand
func ProcessCommand() (string, error) {

	args := os.Args
	if len(args) < 2 {
		return "", errors.New("Invalid argument provided")
	}

	cmd := args[1]
	switch cmd {
	case "snapshot":
		msg, err := services.CreateSnapshot()
		if err != nil {
			return "", err
		}
		return msg, nil
	case "build":
		msg, err := services.BuildFromSnapshot()
		if err != nil {
			return "", err
		}
		return msg, nil
	default:
		fmt.Printf("%snot a valid command%s, valid usages\n", colors.Red, colors.Reset)
		fmt.Printf("\nseed <command>\n")
		fmt.Printf("\tcommand [ snapshot | build]")
		return "", nil
	}

}
