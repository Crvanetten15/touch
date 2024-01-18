/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "touch",
	Short: "PowerShell adaptaion of `touch` from bash",
	Long:  `Run Touch [filename.txt] to make a file`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 0 {
			for _, arg := range args {
				command := "New-Item"
				params := []string{"-ItemType", "File", "-Path", fmt.Sprintf("./%s", arg)}
				cmd := exec.Command("powershell", append([]string{command}, params...)...)
				output, err := cmd.CombinedOutput()
				if err != nil {
					fmt.Println("Error executing command:", err)
				} else {
					fmt.Println(string(output))
				}
			}
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.touch.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
