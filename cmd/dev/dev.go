// Package dev provides the dev command for running development infrastructure
package dev

import (
	"context"
	"os/signal"
	"os/user"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/bketelsen/tiny/cmd"
	"github.com/bketelsen/tiny/enats"
	natsserver "github.com/nats-io/nats-server/v2/server"
	"github.com/urfave/cli/v2"
)

func init() {
	cmd.Register(
		&cli.Command{
			Name:   "dev",
			Usage:  "run a development nats server",
			Action: Run,
			Flags:  Flags,
		},
	)
}

func Run(c *cli.Context) error {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	dataPath := c.String("data-dir")
	expandedPath := ExpandPath(dataPath)
	ns, err := enats.New(ctx, enats.WithNATSServerOptions(&natsserver.Options{
		JetStream: true,
		StoreDir:  expandedPath,
	}))
	if err != nil {
		panic(err)
	}

	ns.NatsServer.Start()

	ns.WaitForServer()
	go func() {
		<-ctx.Done()
		ns.NatsServer.Shutdown()
	}()

	ns.NatsServer.WaitForShutdown()

	return nil
}

var Flags = []cli.Flag{
	&cli.StringFlag{
		Name:    "data-dir",
		Usage:   "persistent data storage directory",
		EnvVars: []string{"TINY_DATA_DIR"},
		Value:   "~/.config/tiny/data",
	},
}

// ExpandPath is a helper function to expand a relative or home-relative path to an absolute path.
//
// eg. ~/.someconf -> /home/alec/.someconf
// copied from github.com/alecthomas/kong
func ExpandPath(path string) string {
	if filepath.IsAbs(path) {
		return path
	}
	if strings.HasPrefix(path, "~/") {
		user, err := user.Current()
		if err != nil {
			return path
		}
		return filepath.Join(user.HomeDir, path[2:])
	}
	abspath, err := filepath.Abs(path)
	if err != nil {
		return path
	}
	return abspath
}
