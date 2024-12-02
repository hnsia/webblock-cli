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

// blockCmd represents the block command
var blockCmd = &cobra.Command{
	Use:   "block",
	Short: "The block subcommand will be creating the blocks",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Block certain URLs in /etc/hosts")
		fmt.Println("Or where the users configured!")
		hosts, err := txeh.NewHostsDefault()
		if err != nil {
			panic(err)
		}
		// hostsfile, _ := cmd.Flags().GetString("hosts-file")
		// sites, _ := cmd.Flags().GetStringSlice("sites")
		// block.BlockSites(hosts, sites)
		viperSites := viper.GetStringSlice("sites")
		viperSites = utils.AddWww(viperSites)
		fmt.Println(viperSites)
		block.BlockSites(hosts, viperSites)
	},
}

func init() {
	rootCmd.AddCommand(blockCmd)

	blockCmd.PersistentFlags().StringP("hosts-file", "f", "/etc/hosts", "a custom hostfile path other than /etc/hosts")
	blockCmd.PersistentFlags().StringSliceP("sites", "s", []string{}, "sites to block")
	viper.BindPFlag("hostsfile", blockCmd.PersistentFlags().Lookup("hostsfile"))
	viper.BindPFlag("sites", blockCmd.PersistentFlags().Lookup("sites"))
}
