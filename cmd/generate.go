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
	"time"

	"github.com/fatih/color"
	"github.com/novafex/goral/fs"
	"github.com/novafex/goral/gen"
	"github.com/novafex/goral/utils"
	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:     "generate [declaration name]",
	Aliases: []string{"build", "compile"},
	Short:   "Generate Go files from declarations",
	Long: `Generate usable Go files from the declarations within the "./goral/*"
directory. Each declaration file (.yaml | .json) is treated as an individual
declaration for the purposes of building.

You can optionally provide a name for the declaration and only it will be
generated instead of all of them.`,

	Run: func(cmd *cobra.Command, args []string) {
		var err error
		var paths []string

		if len(args) == 0 {
			// Find them all
			debugPrint("looking for any and all files in %s\n", gtd)
			paths, err = fs.FindAllWithExtensions(gtd)
			if err != nil {
				color.HiRed("Failed to find declarations in %s", gtd)
				color.Red(`Ensure the project is valid with "goral verify" or "goral init"`)
				cobra.CheckErr(err)
			}
		} else {
			// Use the provided names
			debugPrint("using the provided arguments\n")
			paths = make([]string, len(args))
			var path string
			for i, str := range args {
				path = filepath.Join(gtd, utils.ToKebabCase(str))
				debugPrint("looking for declaration %s\n", path)
				if ok, ext := fs.FindPathWithExtensions(path); ok {
					debugPrint("found declaration %s.%s\n", path, ext)
					paths[i] = fs.CombineBaseExt(path, ext)
				}
			}
		}

		// Early return if no paths found
		if len(paths) == 0 {
			color.Yellow("No declaration files where found in directory %s", gtd)
			return
		}
		paths = utils.CleanStringSlice(paths)

		// Output for fun
		color.Blue("Found %d declarations:", len(paths))
		for i, _ := range paths {
			fmt.Printf("\t%s\n", paths[i])
		}

		// Check if the args don't match
		if len(args) > 0 && len(paths) != len(args) {
			color.Red("Mismatch on found paths, compared to arguments!")
			color.Yellow("You asked for %d declarations, we only found %d", len(args), len(paths))
			return
		}

		// Perf time check
		start := time.Now()

		// Run the processing, possibly in parallel mode
		var errs []error
		if ok, _ := cmd.Flags().GetBool("parallel"); ok {
			errs = gen.ParallelProcessDeclarationFiles(paths)
		} else {
			// Run in single mode
			errs = gen.ProcessDeclarationFiles(paths)
		}

		errs = utils.CleanErrorSlice(errs)
		if len(errs) > 0 {
			// Something went wrong.
			color.HiRed("%d errors occured while processing...", len(errs))
			for i, err := range errs {
				color.Red("%s - %s", paths[i], err.Error())
			}
			os.Exit(1)
		}

		perf := time.Since(start).Milliseconds()
		color.HiGreen("Completed generation of %d declarations in %dms", len(paths), perf)
	},
}

func init() {
	generateCmd.Flags().BoolP("parallel", "p", false, "if provided all the declarations will be split into concurrent threads")

	rootCmd.AddCommand(generateCmd)
}
