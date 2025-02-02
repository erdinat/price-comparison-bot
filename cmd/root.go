package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "myapp",
	Short: "MyApp is a CLI application",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to MyApp!")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
