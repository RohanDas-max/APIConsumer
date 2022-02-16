package controller

import (
	"context"
	"os"
	"testing"
	"time"
)

func TestController(t *testing.T) {

	ctx, cancel := context.WithTimeout(context.Background(), 7*time.Second)
	defer cancel()
	type args struct {
		ctx      context.Context
		username string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "pass ",
			args: args{
				ctx:      ctx,
				username: "rohandas-max",
			},
			wantErr: false,
		},
		{
			name: "faile empty username",
			args: args{
				ctx:      ctx,
				username: "",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Controller(tt.args.ctx, tt.args.username); (err != nil) != tt.wantErr {
				t.Errorf("Controller() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
		os.Remove(tt.args.username + ".txt")
	}
}
