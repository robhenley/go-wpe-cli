package helpers

import (
	"reflect"
	"testing"
)

type test struct {
	name string
	args string
	want bool
}

func TestIsValidEnvironment(t *testing.T) {

	tests := []test{
		{
			name: "Valid Environment",
			args: "development",
			want: true,
		},
		{
			name: "Invalid Environment",
			args: "invalid",
			want: false,
		},
		{
			name: "Empty Environment",
			args: "",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidEnvironment(tt.args); got != tt.want {
				t.Errorf("IsValidEnvironment() = %v, want %v", got, tt.want)
			}
		})
	}

}

func TestIsValidRole(t *testing.T) {
	tests := []test{
		{
			name: "Valid Role",
			args: "owner",
			want: true,
		},
		{
			name: "Invalid Role",
			args: "invalid",
			want: false,
		},
		{
			name: "Empty Role",
			args: "",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidRole(tt.args); got != tt.want {
				t.Errorf("IsValidRole() = %v, want %v", got, tt.want)
			}
		})
	}

}

func TestHasTags(t *testing.T) {
	type args struct {
		filters   []string
		objValues []string
	}

	type test struct {
		name string
		args args
		want bool
	}

	tests := []test{
		{
			name: "Object contains tag",
			args: args{
				filters:   []string{"Client", "Pending"},
				objValues: []string{"Client"},
			},
			want: true,
		},
		{
			name: "Object does not contains tag",
			args: args{
				filters:   []string{"Client", "Pending"},
				objValues: []string{"Team"},
			},
			want: false,
		},
		{
			name: "Filters are empty",
			args: args{
				filters:   []string{},
				objValues: []string{"Team"},
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasTags(tt.args.filters, tt.args.objValues); got != tt.want {
				t.Errorf("HasTags() = %v, want %v", got, tt.want)

			}
		})
	}

}

func TestHasGroup(t *testing.T) {
	type args struct {
		filters []string
		group   string
	}
	type test struct {
		name string
		args args
		want bool
	}

	tests := []test{
		{
			name: "Object has the group",
			args: args{
				filters: []string{
					"clients",
					"team",
					"bob",
					"theme",
					"gone",
					"however",
				},
				group: "team",
			},
			want: true,
		},
		{
			name: "Object does not have group",
			args: args{
				filters: []string{
					"clients",
					"team",
				},
				group: "deleted",
			},
			want: false,
		},
		{
			name: "Object does not have group",
			args: args{
				filters: []string{
					"clients",
					"team",
				},
				group: "deleted",
			},
			want: false,
		},
		{
			name: "Group comparison is case insensitive",
			args: args{
				filters: []string{
					"Clients",
					"Team",
					"Deleted",
				},
				group: "deleted",
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasGroup(tt.args.filters, tt.args.group); got != tt.want {
				t.Errorf("HasGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPrepareFilters(t *testing.T) {
	type args struct {
		filters []string
	}

	type test struct {
		name string
		args args
		want map[string][]string
	}

	tests := []test{
		{
			name: "Passed a single group",
			args: args{
				filters: []string{"group=my team"},
			},
			want: map[string][]string{
				"group": {
					"my team",
				},
			},
		},
		{
			name: "Passed multiple groups",
			args: args{
				filters: []string{
					"group=my team",
					"group=admins",
				},
			},
			want: map[string][]string{
				"group": {
					"my team",
					"admins",
				},
			},
		},
		{
			name: "Passed multiple groups and tags",
			args: args{
				filters: []string{
					"group=my team",
					"group=admins",
					"tag=event",
					"tag=live",
				},
			},
			want: map[string][]string{
				"group": {
					"my team",
					"admins",
				},
				"tag": {
					"event",
					"live",
				},
			},
		},
		{
			name: "Passed empty filters",
			args: args{
				filters: []string{},
			},
			want: map[string][]string{},
		},
		{
			name: "Comparison is case insensitive",
			args: args{
				filters: []string{
					"group=My Team",
					"tag=Event",
				},
			},
			want: map[string][]string{
				"group": {
					"my team",
				},
				"tag": {
					"event",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PrepareFilters(tt.args.filters)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PrepareFilters() = %v, want %v", got, tt.want)
			}
		})
	}

}
