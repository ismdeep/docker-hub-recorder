package handler

import (
	"github.com/ismdeep/rand"
	"testing"
)

func TestImageAdd(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "",
			args: args{
				name: "test-" + rand.HexStr(20),
			},
			wantErr: false,
		},
		{
			name: "",
			args: args{
				name: "alpine",
			},
			wantErr: false,
		},
		{
			name: "",
			args: args{
				name: "ubuntu",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ImageAdd(tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("ImageAdd() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
