package main

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

// The *Run functions are executed in the following order:
//   * PersistentPreRun()
//   * PreRun()
//   * Run()
//   * PostRun()
//   * PersistentPostRun()

func main() {
	cmd := &cobra.Command{
		Use: "app",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("This is PersistentPreRunE")
			// return errors.New("this error is from PersistentPreRunE")
			return nil
		},
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			fmt.Println("This is PersistentPreRun")
		},
		PreRun: func(cmd *cobra.Command, args []string) {
			fmt.Println("This is PreRun")
		},
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("This is Run")
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("This is RunE")
			// return errors.New("this error is from RunE")
			return nil
		},
		PostRun: func(cmd *cobra.Command, args []string) {
			fmt.Println("This is PreRun")
		},
		PersistentPostRun: func(cmd *cobra.Command, args []string) {
			fmt.Println("This is PersistentPostRun")
		},
		PersistentPostRunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("This is PersistentPostRunE")
			return errors.New("this error is from PersistentPostRunE")
		},
	}
	err := cmd.Execute()
	fmt.Println(err)
}
