package main

import (
	"context"
	"os"

	"github.com/denisgrosek/changelock/internal/preflightcli"
)

var (
	version = "dev"
	commit  = "unknown"
	date    = "unknown"
)

func main() {
	app, err := preflightcli.NewApp(os.Getenv, preflightcli.DefaultRuntime(preflightcli.VersionInfo{
		Version: version,
		Commit:  commit,
		Date:    date,
	}))
	if err != nil {
		_, _ = os.Stderr.WriteString(err.Error() + "\n")
		os.Exit(preflightcli.ExitExecution)
	}
	os.Exit(app.Run(context.Background(), os.Args[1:], os.Stdout, os.Stderr))
}
