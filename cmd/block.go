/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// blockCmd represents the block command
var blockCmd = &cobra.Command{
	Use:   "block",
	Short: "The block subcommand will be creating the blocks",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Block certain URLs in /etc/hosts")
		fmt.Println("Or where the users configured!")
		hostsfile, _ := cmd.Flags().GetString("hosts-file")
		fmt.Println(hostsfile)
	},
}

func init() {
	rootCmd.AddCommand(blockCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	blockCmd.PersistentFlags().String("hosts-file", "/etc/hosts", "A custom hostfile path other than /etc/hosts")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// blockCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
