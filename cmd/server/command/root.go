package command

import (
	"github.com/kenlabs/permastar/pkg/system"
	"github.com/spf13/cobra"

	"github.com/kenlabs/permastar/pkg/option"
)

var Opt *option.DaemonOptions

var ExampleUsage = `
# Init permastar configs(default path is ~/.permastar/config.yaml).
permastar-server init

# Start permastar http api server
permastar-server daemon
`

func NewRoot() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:        "permastar-server",
		Short:      "permastar server cli",
		Example:    ExampleUsage,
		SuggestFor: []string{"permastar-server"},
	}

	Opt = option.New(rootCmd)
	msg, err := Opt.Parse()
	if err != nil {
		system.Exit(1, err.Error())
	}
	if msg != "" {
		system.Exit(0, msg)
	}

	childCommands := []*cobra.Command{
		InitCmd(),
		DaemonCmd(),
	}
	rootCmd.AddCommand(childCommands...)

	return rootCmd
}
