package runtime

import (
	"fmt"
	"strings"
)

var runtimes = map[string]Creator{}

// Register registers a runtime.
func Register(rt Runtime, f Creator) {
	rts := strings.ToLower(string(rt))
	if _, ok := runtimes[rts]; ok {
		panic(fmt.Sprintf("duplicate runtime %q", rts))
	}
	runtimes[rts] = f
}

// New returns an instantiated runtime.
func New(rt Runtime) (Runner, error) {
	f, ok := runtimes[string(rt)]
	if !ok {
		return nil, fmt.Errorf("invalid runtime: %q", string(rt))
	}

	return f()
}
