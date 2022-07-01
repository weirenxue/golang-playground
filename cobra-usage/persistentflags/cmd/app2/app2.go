package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func main() {
	cmd := &cobra.Command{
		Use: "app",
		Run: func(cmd *cobra.Command, args []string) {
			name, err := cmd.PersistentFlags().GetString("name")
			fmt.Printf("PersistentFlags().GetString(\"name\") => name: %s, err: %v\n", name, err)
			name, err = cmd.Flags().GetString("name")
			fmt.Printf("Flags().GetString(\"name\") => name: %s, err: %v\n", name, err)
		},
	}

	sub1 := &cobra.Command{
		Use: "sub1",
		Run: func(cmd *cobra.Command, args []string) {
			name, err := cmd.PersistentFlags().GetString("name")
			fmt.Printf("PersistentFlags().GetString(\"name\") => name: %s, err: %v\n", name, err)
			name, err = cmd.Flags().GetString("name")
			fmt.Printf("Flags().GetString(\"name\") => name: %s, err: %v\n", name, err)
		},
	}

	sub1sub1 := &cobra.Command{
		Use: "sub1",
		Run: func(cmd *cobra.Command, args []string) {
			name, err := cmd.PersistentFlags().GetString("name")
			fmt.Printf("PersistentFlags().GetString(\"name\") => name: %s, err: %v\n", name, err)
			name, err = cmd.Flags().GetString("name")
			fmt.Printf("Flags().GetString(\"name\") => name: %s, err: %v\n", name, err)
		},
	}

	sub2 := &cobra.Command{
		Use: "sub2",
		Run: func(cmd *cobra.Command, args []string) {
			name, err := cmd.PersistentFlags().GetString("name")
			fmt.Printf("PersistentFlags().GetString(\"name\") => name: %s, err: %v\n", name, err)
			name, err = cmd.Flags().GetString("name")
			fmt.Printf("Flags().GetString(\"name\") => name: %s, err: %v\n", name, err)
		},
	}

	cmd.AddCommand(sub1, sub2)
	sub1.AddCommand(sub1sub1)

	cmd.PersistentFlags().String("name", "John", "your name")
	sub1.PersistentFlags().String("name", "May", "sub1 your name")

	cmd.Execute()
}
