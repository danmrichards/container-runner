package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	containerRuntime "github.com/danmrichards/container-runner/internal/runtime"
	_ "github.com/danmrichards/container-runner/internal/runtime/all"
)

var (
	runtimeFlag      containerRuntime.Runtime
	workloadDir, cmd string
)

func main() {
	flag.Var(
		&runtimeFlag,
		"runtime",
		fmt.Sprintf(
			"name of the runtime to use for the container, allowed values (%s)",
			containerRuntime.StringList(),
		),
	)
	flag.StringVar(&workloadDir, "workloaddir", "", "path/to the directory containing your workload")
	flag.StringVar(&cmd, "cmd", "", "command to run when starting the container")

	flag.Parse()

	r, err := containerRuntime.New(runtimeFlag)
	if err != nil {
		log.Fatal(err)
	}

	if err = r.Run(context.Background(), workloadDir, cmd); err != nil {
		log.Fatal(err)
	}

	fmt.Println("container running via", runtimeFlag)
}
