package cmd

import "testing"

func Test_getWorld(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "success",
			want: "world is mine",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getWorld(); got != tt.want {
				t.Errorf("getWorld() = %v, want %v", got, tt.want)
			}
		})
	}
}
