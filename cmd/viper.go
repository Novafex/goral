package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/novafex/goral/fs"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if optCfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(optCfgFile)
	} else {
		// Find home directory to add as backup
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)
		viper.AddConfigPath(home)

		// Use the current working directory
		viper.AddConfigPath(cwd)

		// Set the config name
		viper.SetConfigName(CONFIG_NAME)

		// Figure out the type by checking if one we like exists
		_, cfgType := fs.FindPathWithExtensions(filepath.Join(cwd, CONFIG_NAME))
		viper.SetConfigType(cfgType)
	}

	viper.SetEnvPrefix("GORAL")
	viper.AutomaticEnv()

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}

// sets the default viper options, which in turn, is the default config file
// options
func setDefaultConfig() {
	viper.SetDefault("version", "1.0.0")
}
