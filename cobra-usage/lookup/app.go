package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func main() {
	cmd := &cobra.Command{
		Use: "app",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Lookup(\"the-only-flag\"): %#v\n", cmd.Flags().Lookup("the-only-flag"))
			fmt.Printf("Lookup(\"not-exist-flag\"): %#v\n", cmd.Flags().Lookup("not-exist-flag"))
		},
	}

	cmd.Flags().StringSlice("the-only-flag", []string{}, "Your string slice")

	cmd.Execute()
}
