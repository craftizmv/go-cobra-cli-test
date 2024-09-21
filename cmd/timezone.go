/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"time"
)

// timezoneCmd represents the timezone command
var timezoneCmd = &cobra.Command{
	Use:   "timezone",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		timezone := args[0]
		location, _ := time.LoadLocation(timezone)
		dateFlag, _ := cmd.Flags().GetString("date")
		var date string

		if dateFlag != "" {
			date = time.Now().In(location).Format(dateFlag)
		} else {
			date = time.Now().In(location).Format(time.RFC3339)[:10]
		}
		fmt.Printf("Current date in %v: %v\n", timezone, date)
	},
}

func init() {
	rootCmd.AddCommand(timezoneCmd)
	timezoneCmd.Flags().StringP("date", "d", "", "Date for which to get the time (format: yyyy-mm-dd)")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// timezoneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// timezoneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
