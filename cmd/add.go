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
	"path/filepath"

	"github.com/fatih/color"
	"github.com/novafex/goral/decl"
	"github.com/novafex/goral/fs"
	"github.com/novafex/goral/utils"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add name...",
	Short: "Makes a new declaration file of the given name",
	Long: `Use to scaffold a templated declaration file with the given name. By
default it will use YAML for the format. You can specify more than one name to
create multiple templates.`,
	Args: cobra.MinimumNArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		if len(name) == 0 {
			cobra.CheckErr("missing argument for command \"add\"")
		}

		filename := utils.ToKebabCase(name)
		base := filepath.Join(gtd, filename)
		ext := fs.GetExtensionOrder()[0]
		path := fs.CombineBaseExt(base, ext)
		if fs.FileExists(path) {
			cobra.CheckErr(fmt.Errorf("declaration %s already exists", filename))
		}

		if err := fs.Write(decl.ExampleDeclaration, base, ext); err != nil {
			cobra.CheckErr(err)
		}
		color.HiGreen("Created new declaration at %s", path)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
