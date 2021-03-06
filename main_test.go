package main

import "testing"

func Test_createCommand(t *testing.T) {
	type args struct {
		file    string
		version float64
		dir     string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "create command",
			args: args{file: "sample.py", version: 3.7},
			want: "CMD=\"python sample.py\" docker-compose -f py37/docker-compose.yml up",
		},
		{
			name: "create command",
			args: args{file: "sample.py", version: 3.7, dir: "sample"},
			want: "CMD=\"cd sample;pip install -r requirements.txt;python sample.py\" docker-compose -f py37/docker-compose.yml up",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createCommand(tt.args.file, tt.args.version, tt.args.dir); got != tt.want {
				t.Errorf("createCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}
