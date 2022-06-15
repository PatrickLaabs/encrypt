/*
Copyright © 2022 Patrick Laabs <patrick.laabs@me.com>

*/
package cmd

import (
	"crypto/sha512"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// sha512Cmd represents the sha512 command
var sha512Cmd = &cobra.Command{
	Use:   "sha512",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		data := []byte(os.Args[2])
		hash := sha512.Sum512(data)
		fmt.Printf("%s\n", data)
		fmt.Printf("%x\n", hash)
	},
}

func init() {
	rootCmd.AddCommand(sha512Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sha512Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sha512Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
