package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func main() {
	var cfg string
	var host string
	cobra.OnInitialize(func() {
		viper.SetConfigFile(cfg)
		if err := viper.ReadInConfig(); err != nil {
			return
		}
	})

	cmd := &cobra.Command{
		Use: "app",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("in command Run: host=%s\n", host)
			fmt.Printf("in command Run: viper.GetString(\"host\")=%s\n", viper.GetString("host"))
		},
	}

	cmd.Flags().StringVar(&cfg, "cfg", "", "Config file")

	cmd.Flags().StringVar(&host, "host", "localhost", "server host")
	viper.BindPFlag("host", cmd.Flags().Lookup("host"))

	cmd.Execute()
}
