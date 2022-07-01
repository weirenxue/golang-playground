package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func main() {
	cmd := &cobra.Command{
		Use: "app",
		Run: func(cmd *cobra.Command, args []string) {
			name, _ := cmd.Flags().GetString("name")
			fmt.Printf("Changed(\"name\"): %#v\nvalue: %s\n", cmd.Flags().Changed("name"), name)
		},
	}

	cmd.Flags().String("name", "John", "your name")

	cmd.Execute()
}
