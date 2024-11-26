/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	block "github.com/hnsia/webblock-cli/pkg"
	"github.com/spf13/cobra"
	"github.com/txn2/txeh"
)

// unblockCmd represents the unblock command
var unblockCmd = &cobra.Command{
	Use:   "unblock",
	Short: "Unblock given sites",
	Long:  `Unblock given sites.`,
	Run: func(cmd *cobra.Command, args []string) {
		hosts, err := txeh.NewHostsDefault()
		if err != nil {
			panic(err)
		}
		sites, _ := cmd.Flags().GetStringSlice("sites")
		block.CleanBlocks(hosts, sites)
	},
}

func init() {
	rootCmd.AddCommand(unblockCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// unblockCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// unblockCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	unblockCmd.PersistentFlags().StringSliceP("sites", "s", []string{}, "sites to unblock")
	unblockCmd.MarkFlagRequired("sites")
}
