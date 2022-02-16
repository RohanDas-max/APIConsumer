package controller

import (
	"context"
	"os"
	"testing"
	"time"
)

func TestController(t *testing.T) {
	type args struct {
		ctx      context.Context
		username string
		t        time.Duration
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "pass ",
			args: args{
				ctx:      context.Background(),
				username: "rohandas-max",
				t:        1 * time.Second,
			},
			wantErr: false,
		},
		{
			name: "faile empty username",
			args: args{
				ctx:      context.Background(),
				username: "",
				t:        1 * time.Second,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Controller(tt.args.ctx, tt.args.username, tt.args.t); (err != nil) != tt.wantErr {
				t.Errorf("Controller() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
		os.Remove(tt.args.username + ".txt")
	}
}
