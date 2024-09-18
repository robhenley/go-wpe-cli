package installs

import "strings"

func isValidEnvironment(environment string) bool {
	switch strings.ToLower(environment) {
	case "production", "staging", "development":
		return true
	default:
		return false
	}
}
