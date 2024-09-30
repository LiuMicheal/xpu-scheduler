package main

import (
	"xpu-scheduler/pkg/plugins"

	"math/rand"
	"os"
	"time"

	"k8s.io/component-base/cli"
	"k8s.io/component-base/logs"
	"k8s.io/kubernetes/cmd/kube-scheduler/app"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	command := app.NewSchedulerCommand(
		// Register plugin
		app.WithPlugin(plugins.Name, plugins.New),
	)
	go plugins.Cachestart()
	logs.InitLogs()
	defer logs.FlushLogs()

	code := cli.Run(command)
	os.Exit(code)

}
