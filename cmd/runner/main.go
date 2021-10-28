package main

import (
	"flag"
	"fmt"

	containerRuntime "github.com/danmrichards/container-runner/internal/runtime"
)

var runtimeFlag, workingDir, entrypoint string

func main() {
	flag.StringVar(
		&runtimeFlag,
		"runtime",
		string(containerRuntime.Docker),
		fmt.Sprintf(
			"name of the runtime to use for the container, allowed values (%s)",
			containerRuntime.StringList(),
		),
	)
	flag.StringVar(&workingDir, "workingdir", "", "path/to the working directory")
	flag.StringVar(&entrypoint, "entrypoint", "", "path/to the container entrypoint")

	flag.Parse()
}
