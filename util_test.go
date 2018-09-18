package flatstorage

import "testing"

func Test_pathExists(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Positive Test",
			args: args{
				path: "/dev/null",
			},
			want: true,
		},
		{
			name: "Negative Test",
			args: args{
				path: "/dev/thiswontexist",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pathExists(tt.args.path); got != tt.want {
				t.Errorf("pathExists() = %v, want %v", got, tt.want)
			}
		})
	}
}
