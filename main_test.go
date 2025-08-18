package main

import "testing"

func Test_greeting(t *testing.T) {
	type args struct {
		name *string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "with name",
			args: args{
				name: pointer("John"),
			},
			want: "Hello, John!",
		},
		{
			name: "without name",
			args: args{
				name: nil,
			},
			want: "Hello!",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			greeting(tt.args.name)
		})
	}
}

func pointer[T any](value T) *T {
	return &value
}
