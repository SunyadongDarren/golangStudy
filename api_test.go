package golangStudy

import "testing"

func TestHelloGolangStudy(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: `a`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HelloGolangStudy(tt.args.name); got != tt.want {
				t.Errorf("HelloGolangStudy() = %v, want %v", got, tt.want)
			}
		})
	}
}
