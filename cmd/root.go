package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gencld",
	Short: "Gencld is a command tool to generate CLD over protocal buffers",
	Long:  "Cencld is a command tool to generate CLD over protocal buffers",
}

// Execute runs gencld.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
