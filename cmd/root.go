/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "webblock-cli",
	Short: "Blocks local access to certain websites.",
	Long:  `To be worked on later...`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
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
	cobra.OnInitialize(initConfig)
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is /.config.example.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// // Find home directory.
		// home, err := os.UserHomeDir()
		// if home == "/root" {
		// 	fmt.Println("You are using root privileges, please pass in custom config file path as a flag.")
		// }
		// fmt.Printf("Looking in home at '%s' for default config file\n", home)
		// cobra.CheckErr(err)

		// // Search config in home directory with name ".webblock-cli" (without extension).
		// viper.AddConfigPath(home)
		// viper.SetConfigType("yaml")
		// viper.SetConfigName(".webblock-cli")

		// Find current directory.
		currDir, err := os.Getwd()
		cobra.CheckErr(err)
		fmt.Printf("Looking in current directory at '%s' for default config file\n", currDir)

		// Search config in current directory with name ".config.example" (without extension).
		viper.AddConfigPath(currDir)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".config.example")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
