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

	unblockCmd.PersistentFlags().StringP("hosts-file", "f", "/etc/hosts", "a custom hostfile path other than /etc/hosts")
	unblockCmd.PersistentFlags().StringSliceP("sites", "s", []string{}, "sites to unblock")
	unblockCmd.MarkFlagRequired("sites")
	viper.BindPFlag("sites", unblockCmd.PersistentFlags().Lookup("sites"))
	viper.BindPFlag("hostsfile", unblockCmd.PersistentFlags().Lookup("hostsfile"))
}
