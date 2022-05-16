package reverse_string

import "testing"

func TestReverseString(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name       string
		args       args
		wantOutput string
	}{
		{
			name:       "simple string",
			args:       args{"bobby"},
			wantOutput: "ybbob",
		},
		{
			name:       "special string",
			args:       args{"bobby \n hill"},
			wantOutput: "llih \n ybbob",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutput := ReverseString(tt.args.input); gotOutput != tt.wantOutput {
				t.Errorf("ReverseString() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}
