package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func main() {
	cmd := &cobra.Command{
		Use: "app",
		Run: func(cmd *cobra.Command, args []string) {
			slice, err := cmd.Flags().GetStringSlice("slice")
			if err != nil {
				fmt.Println(err)
			}
			fmt.Printf("GetStringSlice: %#v\n", slice)
			array, err := cmd.Flags().GetStringArray("array")
			if err != nil {
				fmt.Println(err)
			}
			fmt.Printf("GetStringArray: %#v\n", array)
		},
	}

	cmd.Flags().StringSlice("slice", []string{}, "Your string slice")
	cmd.Flags().StringArray("array", []string{}, "Your array slice")

	cmd.Execute()
}
