package helpers

import "testing"

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
