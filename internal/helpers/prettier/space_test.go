// BEGIN: 8f7e6d5c3b2a
package prettier

import "testing"

func TestRemoveSpaces(t *testing.T) {
	tests := []struct {
		name string
		args string
		want string
	}{
		{
			name: "removes spaces",
			args: "hello world",
			want: "helloworld",
		},
		{
			name: "removes tabs",
			args: "hello\tworld",
			want: "helloworld",
		},
		{
			name: "removes newlines",
			args: "hello\nworld",
			want: "helloworld",
		},
		{
			name: "removes all spaces",
			args: "   hello \t\n world   ",
			want: "helloworld",
		},
		{
			name: "handles empty string",
			args: "",
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveSpaces(tt.args); got != tt.want {
				t.Errorf("removeSpaces() = %v, want %v", got, tt.want)
			}
		})
	}
}

// END: 8f7e6d5c3b2a
