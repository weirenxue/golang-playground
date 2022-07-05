package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func main() {
	var cfg string
	cobra.OnInitialize(func() {
		fmt.Printf("in cobra.OnInitialize, cfg=%s\n", cfg)
	})

	cmd := &cobra.Command{
		Use:               "app",
		PersistentPreRun:  func(cmd *cobra.Command, args []string) { fmt.Println("in command PersistentPreRun") },
		PreRun:            func(cmd *cobra.Command, args []string) { fmt.Println("in command PreRun") },
		Run:               func(cmd *cobra.Command, args []string) { fmt.Println("in command Run") },
		PostRun:           func(cmd *cobra.Command, args []string) { fmt.Println("in command PostRun") },
		PersistentPostRun: func(cmd *cobra.Command, args []string) { fmt.Println("in command PersistentPostRun") },
	}

	cmd.Flags().StringVar(&cfg, "cfg", "", "Config file")

	cmd.Execute()
}
