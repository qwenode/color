package color

import "testing"

func TestSuccessMessage(t *testing.T) {
	type args struct {
		format string
		a      []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			args: args{format: "success"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SuccessMessage(tt.args.format, tt.args.a)
		})
	}
}
