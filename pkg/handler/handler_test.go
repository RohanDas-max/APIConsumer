package handler

import (
	"context"
	"os"
	"testing"
)

func TestHandler(t *testing.T) {
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
			name: "pass with username rohandas-max",
			args: args{
				ctx:      context.Background(),
				username: "rohandas-max",
			},
			wantErr: false,
		},
		{
			name: "fail with wrong username",
			args: args{
				ctx:      context.Background(),
				username: "rohandas-max123456789",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Handler(tt.args.ctx, tt.args.username); (err != nil) != tt.wantErr {
				t.Errorf("Handler() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
		os.Remove(tt.args.username + ".txt")
	}

}

func Test_write(t *testing.T) {
	type args struct {
		filename string
		response response
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "pass with filename and with response",
			args: args{
				filename: "random",
				response: response{},
			},
			wantErr: false,
		},
		{
			name: "pass with filename and without response",
			args: args{
				filename: "random",
				response: response{
					data: data{
						Id:        1,
						User:      "rohan",
						Followers: 1,
						Following: 1,
						Repo:      "random.random",
						Orgs:      "random.random",
					},
					repo: []repo{{
						Name: "rohan",
					}},
					Org: []org{{
						Name:        "random org",
						Description: "lorem ipsumlorem ipsum",
					}},
				},
			},
			wantErr: false,
		},
		{
			name: "pass without filename",
			args: args{
				filename: "",
				response: response{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := write(tt.args.filename, tt.args.response); (err != nil) != tt.wantErr {
				t.Errorf("write() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
		os.Remove(tt.args.filename + ".txt")
	}
}
