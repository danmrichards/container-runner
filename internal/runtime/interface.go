package runtime

import "context"

// Runner is the interface implemented by container runtimes.
type Runner interface {
	// Run runs a container with the given working directory and entrypoint.
	Run(ctx context.Context, workingDir, cmd string) error
}

// Creator is a function that creates a Runner.
type Creator func() (Runner, error)
