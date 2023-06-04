package command

import (
	"fmt"
	"github.com/kenlabs/permastar/pkg/api/core"
	"github.com/kenlabs/permastar/pkg/api/v1/server"
	ipfs_shell "github.com/kenlabs/permastar/pkg/ipfs-shell"
	"github.com/kenlabs/permastar/pkg/util/log"
	"os"
	"os/signal"
	"syscall"

	logging "github.com/ipfs/go-log/v2"
	"github.com/spf13/cobra"
)

var logger = log.NewSubsystemLogger()

func DaemonCmd() *cobra.Command {
	const failedError = "run daemon failed: \n\t%v\n"
	return &cobra.Command{
		Use:   "daemon",
		Short: "start permastar server",
		RunE: func(cmd *cobra.Command, args []string) error {
			err := setLogLevel()
			if err != nil {
				return fmt.Errorf(failedError, err)
			}

			c, err := initCore()
			if err != nil {
				return err
			}
			apiServer, err := server.NewAPIServer(Opt, c)
			apiServer.MustStartAllServers()

			quit := make(chan os.Signal)
			signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
			<-quit
			fmt.Println("shutting down servers...")

			return nil
		},
	}
}

func setLogLevel() error {
	logLevel, err := logging.LevelFromString(Opt.LogLevel)
	if err != nil {
		return err
	}
	logging.SetAllLoggers(logLevel)

	return nil
}

func initCore() (*core.Core, error) {
	c := &core.Core{}

	c.IPFSNodeAPI = ipfs_shell.NewIPFSNodeShell(Opt.IPFSNode.GatewayIP + ":" + Opt.IPFSNode.GatewayPort)

	return c, nil
}
