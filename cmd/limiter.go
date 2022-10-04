package main

import (
	"os"

	"limiter/cmd/root"
	"limiter/cmd/serve"
	"limiter/internal"
	"limiter/pkg/logger"
)

func main() {
	app := internal.NewApplication()

	rootCmd := root.Cmd(app)
	rootCmd.AddCommand(serve.Cmd(app))

	if err := rootCmd.Execute(); err != nil {
		logger.Log().Errorf("An error occurred: %s", err.Error())
		os.Exit(1)
	}
}
