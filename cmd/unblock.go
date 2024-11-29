/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	block "github.com/hnsia/webblock-cli/pkg"
	"github.com/hnsia/webblock-cli/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
		// sites, _ := cmd.Flags().GetStringSlice("sites")
		// block.CleanBlocks(hosts, sites)
		viperSites := viper.GetStringSlice("sites")
		viperSites = utils.AddWww(viperSites)
		fmt.Println(viperSites)
		block.CleanBlocks(hosts, viperSites)
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
	unblockCmd.PersistentFlags().StringP("hosts-file", "f", "/etc/hosts", "a custom hostfile path other than /etc/hosts")
	unblockCmd.PersistentFlags().StringSliceP("sites", "s", []string{}, "sites to unblock")
	unblockCmd.MarkFlagRequired("sites")
	viper.BindPFlag("sites", unblockCmd.PersistentFlags().Lookup("sites"))
	viper.BindPFlag("hostsfile", unblockCmd.PersistentFlags().Lookup("hostsfile"))
}
