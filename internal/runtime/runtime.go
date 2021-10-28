package runtime

import "strings"

// Runtime represents a type of container runtime.
type Runtime string

const (
	// Docker is the runtime powered by Docker.
	Docker Runtime = "docker"

	// ContainerD is the runtime powered by ContainerD.
	ContainerD Runtime = "containerd"
)

// Valid is the list of valid runtimes.
var Valid = []Runtime{Docker, ContainerD}

// StringList returns the list of valid runtimes.
func StringList() string {
	validStr := make([]string, 0, len(Valid))
	for _, v := range Valid {
		validStr = append(validStr, string(v))
	}

	return strings.Join(validStr, ",")
}
