package helpers

import (
	"slices"
	"strings"
)

var validRoles = []string{
	"owner",
	"full",
	"full,billing",
	"partial",
	"partial,billing",
}

var validEnv = []string{
	"development",
	"staging",
	"production",
}

func ValidRoles() []string {
	return validRoles
}

func IsValidRole(role string) bool {
	return slices.Contains(validRoles, strings.ToLower(role))
}

func IsValidEnvironments() []string {
	return validEnv
}

func IsValidEnvironment(environment string) bool {
	return slices.Contains(validEnv, strings.ToLower(environment))
}
