// Copyright © 2019 InVisionApp


package cmd

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/schigh/mksb/mksb"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var inFile string
var inDelimiter int
var outDelimiter int
var sbName string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "mksb",
	Short: "Convert files or streams to golang string builders",
	Long: `mksb will convert plain text to a set of golang strings.Builder commands`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		if inFile != "" {
			lines, err := mksb.GetLinesInFile(inFile, inDelimiter)
			if err != nil {
				_,_ = fmt.Fprintf(os.Stderr, "Error encountered reading file '%s': %v", inFile, err)
				os.Exit(1)
			}

			sb := mksb.WrapSB(sbName, lines, outDelimiter)
			fmt.Printf("%v", sb.String())
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() { 
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.mksb.yaml)")

	rootCmd.PersistentFlags().StringVar(&inFile, "file", "", "absolute path to file")
	rootCmd.PersistentFlags().IntVar(&inDelimiter, "rd", 10, "default line READ DELIMITER")
	rootCmd.PersistentFlags().IntVar(&outDelimiter, "wd", 10, "default line WRITE DELIMITER")
	rootCmd.PersistentFlags().StringVar(&sbName, "sbName", "sb", "string builder name")

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
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".mksb" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".mksb")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
