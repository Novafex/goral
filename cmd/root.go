/*
Copyright © 2023 Novafex

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

const (
	// Specifies the target directory from the root project folders perspective
	TARGET_DIR = "goral"

	// Target name of the configuration file
	CONFIG_NAME = "goral"
)

var (
	// current working directory
	cwd string

	// goral target directory
	gtd string

	// declares the accepted extensions for config, in order
	configExtensions = []string{
		"toml",
		"yaml",
		"json",
	}

	// configuration file target (set by flag)
	optCfgFile string

	// whether to use Debug mode
	optDebug bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "goral",
	Version: "0.1.0-alpha",
	Short: "CRUD API generator for Go",
	Long: `Automatically generates Go code designed for REST-ish API generation
based on manifest files. With Goral entire back-end services can be generated
using a few configuration (manifest) files written in YAML/JSON.

Some features include:
 * Automatic struct generation for resources
 * Endpoint binding for web servers
 * SQL generation scripts for building the schema
 * Exporting TypeScript declarations
 * And much more...
 
Initialize a project/repository using "goral init" to scaffold the folder
structure and configurations. From there "goral add" will create new manifests,
and "goral generate" will re-generate the Go files.

See "goral help" for more information on available commands and options.`,
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
	var err error

	// Setup the current working directory
	cwd, err = os.Getwd()
	if err != nil {
		// Cobra will exit from this
		cobra.CheckErr(fmt.Errorf("failed to get current working directory: %s", err.Error()))
	}

	// Build target directory
	gtd = filepath.Join(cwd, TARGET_DIR)

	// Cobra can initialize the configuration
	cobra.OnInitialize(initConfig)

	// === Configuration settings ===
	setDefaultConfig()

	// === Persistent flags ===
	rootCmd.PersistentFlags().StringVar(&optCfgFile, "config", "", "config file (default is ./goral.yaml)")
	rootCmd.PersistentFlags().BoolVarP(&optDebug, "debug", "D", false, "enables debug logging")
}
