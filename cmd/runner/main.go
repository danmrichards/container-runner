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

	cid, err := r.Run(context.Background(), workloadDir, cmd)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("container running with ID %q via %s\n", cid, runtimeFlag)
}
