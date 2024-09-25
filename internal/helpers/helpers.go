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

func PrepareFilters(filters []string) map[string][]string {
	f := make(map[string][]string)

	if len(filters) > 0 {
		for _, filter := range filters {
			parts := strings.Split(filter, "=")
			if len(parts) == 2 {
				k := strings.ToLower(parts[0])
				v := strings.ToLower(parts[1])
				f[k] = append(f[k], v)
			}
		}
	}

	return f
}

func HasTags(filterTags, objTags []string) bool {

	if len(filterTags) > 0 {
		for _, v := range filterTags {
			if slices.Contains(objTags, v) {
				return true
			}
		}
	}

	return false
}

func HasGroup(filterGroups []string, group string) bool {
	if len(filterGroups) > 0 {
		group = strings.ToLower(group)
		return slices.ContainsFunc(filterGroups, func(g string) bool {
			return strings.ToLower(g) == group
		})
	}
	return false
}
