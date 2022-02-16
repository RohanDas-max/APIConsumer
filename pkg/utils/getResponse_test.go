package utils

import (
	"context"
	"net/http"
	"testing"
	"time"
)

func TestGet(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 7*time.Second)
	defer cancel()
	type args struct {
		ctx context.Context
		url string
	}
	tests := []struct {
		name string
		args args

		wantErr bool
	}{
		{
			name: "fail(passing wrong url)",
			args: args{
				ctx: ctx,
				url: "asdasdasd",
			},

			wantErr: true,
		},
		{
			name: "pass(passing github api url)",
			args: args{
				ctx: ctx,
				url: "http://google.com",
			},

			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Get(tt.args.ctx, tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_checkStatus(t *testing.T) {
	resp, _ := http.Get("http://google.com")

	type args struct {
		h *http.Response
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "fail(passing nil http.response)",
			args: args{
				h: &http.Response{},
			},
			wantErr: true,
		},
		{
			name: "pass",
			args: args{
				h: resp,
			},
			wantErr: false,
		},
	}
	defer resp.Body.Close()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := checkStatus(tt.args.h)
			if (err != nil) != tt.wantErr {
				t.Errorf("checkStatus() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}
