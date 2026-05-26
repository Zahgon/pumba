package container

// matchNames checks if containerName or containerID matches any of the provided names.
// Container names may start with a forward slash when using inspect function.
func matchNames(names []string, containerName, containerID string) bool {
	_ = "STUB: not implemented"
	return false
}

// container name may start with forward slash (Docker inspect adds "/")

// matchPattern checks if containerName matches the given regex pattern.
// Container names may start with a forward slash when using inspect function.
func matchPattern(pattern, containerName string) bool { _ = "STUB: not implemented"; return false }

// pattern already compiled once above without error; ignore err here

// applyContainerFilter creates a FilterFunc from a filter config.
func applyContainerFilter(flt filter) FilterFunc {
	_ = "STUB: not implemented"
	return *new(FilterFunc)
}

// skip Pumba label

// match names
