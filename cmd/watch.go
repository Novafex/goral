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

	"github.com/fatih/color"
	"github.com/fsnotify/fsnotify"
	"github.com/novafex/goral/fs"
	"github.com/novafex/goral/gen"
	"github.com/spf13/cobra"
)

// watchCmd represents the watch command
var watchCmd = &cobra.Command{
	Use:   "watch",
	Short: "Watches the Goral directory for changes",
	Long: `Watches the Goral directory for changes to declaration files. When
a change is detected it will re-generate the declaration.`,

	Run: func(cmd *cobra.Command, args []string) {
		watcher, err := fsnotify.NewWatcher()
		cobra.CheckErr(err)
		defer watcher.Close()

		done := make(chan bool)

		go func() {
			for {
				select {
				case evt := <-watcher.Events:
					handleWatchEvent(evt)

				case err := <-watcher.Errors:
					color.Red("error during watching: %s", err.Error())
					done <- true
				}
			}
		}()

		if err := watcher.Add(gtd); err != nil {
			cobra.CheckErr(err)
		}

		color.Blue("Watching folder %s for changes...", gtd)
		<-done
	},
}

func init() {
	rootCmd.AddCommand(watchCmd)
}

func handleWatchEvent(evt fsnotify.Event) {
	switch evt.Op {
	case fsnotify.Create:
		handleWatchCreate(evt)
	case fsnotify.Remove:
		handleWatchRemove(evt)
	case fsnotify.Rename:
		handleWatchRename(evt)
	case fsnotify.Write:
		handleWatchChange(evt)
	}
}

func handleWatchCreate(evt fsnotify.Event) {
	debugPrint("New file %s created\n", evt.Name)

	if ok, _ := fs.HasExtension(evt.Name); ok {
		fmt.Printf("Detected new file %s, waiting for content\n", evt.Name)
	}
}

func handleWatchRemove(evt fsnotify.Event) {
	debugPrint("file %s was removed\n", evt.Name)

	// TODO: Remove corresponding go outputs for this file
}

func handleWatchRename(evt fsnotify.Event) {
	debugPrint("file %s was renamed\n", evt.Name)

	// TODO: Remove corresponding go outputs for this file
}

func handleWatchChange(evt fsnotify.Event) {
	if ok, _ := fs.HasExtension(evt.Name); ok {
		fmt.Printf("Change in %s detected\n", evt.Name)

		if err := gen.ProcessDeclarationFile(evt.Name); err != nil {
			color.Red("Error processing declaration file %s\n%s", evt.Name, err.Error())
		} else {
			color.Green("Generated output for declaration %s", evt.Name)
		}
	} else {
		debugPrint("file %s changed but it is not a declaration\n", evt.Name)
	}
}
