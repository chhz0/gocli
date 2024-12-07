package main

import (
	"fmt"
	"strings"

	"github.com/chhz0/gocli"
	"github.com/spf13/cobra"
)

var echoTimes int

func newPrintCmd() *gocli.SimpleCommand {
	return &gocli.SimpleCommand{
		Use:   "print [string to print]",
		Short: "Print anything to the screen",
		Long: `print is for printing anything back to the screen.
For many years people have printed back to the screen.`,
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) error {
			fmt.Println("Print: " + strings.Join(args, " "))
			return nil
		},
	}
}

func newEchoCmd() *gocli.SimpleCommand {
	return &gocli.SimpleCommand{
		Use:   "echo [string to echo]",
		Short: "Echo anything to the screen",
		Long: `echo is for echoing anything back.
Echo works a lot like print, except it has a child command.`,
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) error {
			fmt.Println("Print: " + strings.Join(args, " "))
			return nil
		},
	}
}

func newTimesCmd() *gocli.SimpleCommand {
	return &gocli.SimpleCommand{
		Use:   "times [# times] [string to echo]",
		Short: "Echo anything to the screen more times",
		Long: `echo things multiple times back to the user by providing
a count and a string.`,
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) error {
			for i := 0; i < echoTimes; i++ {
				fmt.Println("Echo: " + strings.Join(args, " "))
			}
			return nil
		},
	}
}

func main() {
	x := gocli.New(&gocli.SimpleCommand{
		Use:   "gocli",
		Short: "gocli is a simple cli tool",
		Run: func(cmd *cobra.Command, args []string) error {
			return nil
		},

		Subcommands: []*gocli.SimpleCommand{
			newPrintCmd(),
			newEchoCmd().AppendCommands(newTimesCmd()),
		},
	})

	x.Execute()
}
