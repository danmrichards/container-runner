// Package all imports all runtimes
package all

import (

	// Import all known runtimes so they register themselves.
	_ "github.com/danmrichards/container-runner/internal/runtime/containerd"
	_ "github.com/danmrichards/container-runner/internal/runtime/docker"
)
