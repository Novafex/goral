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
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init [path]",
	Short: "Initialize Goral within the current working directory",
	Long: `Initializes the current folder (project) to use Goral by creating the
initially needed files and folders.`,

	Args: cobra.MaximumNArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		color.Cyan("Welcome to Goral\n\n")
		
		// Find the base CWD
		var pathAdd string
		if len(args) == 1 {
			pathAdd = filepath.Clean(args[0])
		}
		cwd, _ := os.Getwd()
		wd := filepath.Join(cwd, pathAdd)

		debugPrint("Current working directory: %s\n", wd)

		initStepStructure(wd)
		initStepGitIgnore(wd)

		color.HiGreen("\nCompleted initialization!\nTry using \"goral add\" next to make your first declaration.")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}


func initStepStructure(wd string) {
	// Check for existing folder
	parentDir := filepath.Join(wd, "goral")
	if pathExists(parentDir) {
		debugPrint("Goral directory exists\n")

		// Check existing configuration
		if pathExists(filepath.Join(parentDir, "goral.config.yaml")) || pathExists(filepath.Join(parentDir, "goral.config.json")) {
			debugPrint("Found configuration file\n")
			color.Green(`Looks like Goral is already scaffolded...`)
			color.Yellow(`TIP: You can check existing setups using "goral verify"`)
			return
		}
		debugPrint("No configuration file")

		// Make a blank configuration
		newPath := filepath.Join(parentDir, "goral.config.yaml")
		if err := os.WriteFile(newPath, []byte(""), 0771); err != nil {
			color.HiRed("failed to write config file to %s", newPath)
			panic(err)
		}
		color.HiGreen(`Created new configuration file at %s...`, newPath)
	}

	// Scaffold the whole thing
	if err := os.MkdirAll(parentDir, 1777); err != nil {
		color.HiRed("failed to create new directory %s", parentDir)
		panic(err)
	}

	// Make a blank configuration
	newPath := filepath.Join(parentDir, "goral.config.yaml")
	if err := os.WriteFile(newPath, []byte(""), 0771); err != nil {
		color.HiRed("failed to write config file to %s", newPath)
		panic(err)
	}
	color.Green(`Created new configuration file at %s...`, newPath)
}

func initStepGitIgnore(wd string) {
	const GI_RULE = "\n# Goral generated files\n*_goral.go\n"

	prompt := promptui.Prompt{
		Label: "Should we append a Git ignore rule to remove Goral generated files?",
		IsConfirm: true,
	}
	answer, _ := prompt.Run()
	if isYesAnswer(answer) {
		debugPrint("adding line to .gitignore\n")

		pathGI := filepath.Join(wd, ".gitignore")
		f, err := os.OpenFile(pathGI, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		if err != nil {
			color.HiRed("could not open file %s", pathGI)
			panic(err)
		}
		defer f.Close()

		f.WriteString(GI_RULE)

		color.Green("Setup .gitignore file for you...")
	}
}