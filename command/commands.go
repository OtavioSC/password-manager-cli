package command

import (
	"flag"
	"fmt"

	"github.com/otaviosc/password-manager-cli/internal/types"
)

var commands = []types.Command{
	{
		Name:        "store",
		Description: "Store a password securely",
		Execute:     StorePassword,
	},
}

func InitializeCLI() {
	flag.Usage = func() {
		fmt.Println("Usage: password-manager-cli [command] [options]")
	}

	flag.Parse()

	if len(flag.Args()) == 0 {
		flag.Usage()
		return
	}

	cmd := flag.Arg(0)
	for _, command := range commands {
		if command.Name == cmd {
			fmt.Println("Executing command:", command.Name)
			if err := command.Execute(flag.Args()[1:]); err != nil {
				fmt.Printf("Error executing command '%s': %v\n", cmd, err)
			}
			return
		}
	}
}

func ListCommands() {
	for _, cmd := range commands {
		fmt.Printf("%s: %s\n", cmd.Name, cmd.Description)
	}
}
