package run

import (
	"os/exec"
)

// isDepsInstalled checks if the given dependencies are available in the system PATH.
// It returns a boolean indicating if all are present, and a slice of missing dependencies.
func isDepsInstalled(deps []string) (bool, []string) {
	var missing []string

	for _, dep := range deps {
		_, err := exec.LookPath(dep)
		if err != nil {
			missing = append(missing, dep)
		}
	}

	return len(missing) == 0, missing
}
