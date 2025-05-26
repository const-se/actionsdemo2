package main

import "testing"

func Test_greeting(t *testing.T) {
	type args struct {
		name string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "success",
			args: args{
				name: "John",
			},
			want: "Hello, John!",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			greeting(tt.args.name)
		})
	}
}
