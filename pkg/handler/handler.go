package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/rohandas-max/ghCrwaler/pkg/utils"
)

type data struct {
	Id        int    `json:"id"`
	User      string `json:"login"`
	Followers int    `json:"followers"`
	Following int    `json:"following"`

	Repo string `json:"repos_url"`
	Orgs string `json:"organizations_url"`
}
type repo struct {
	Name string `json:"name"`
}
type org struct {
	Name        string `json:"login"`
	Description string `json:"description"`
}
type response struct {
	data data
	repo []repo
	Org  []org
}

func Handler(ctx context.Context, username string) error {
	url := "http://api.github.com/users/" + username
	var d data
	var r []repo
	var o []org
	if byte, err := utils.Get(ctx, url); err == nil {
		json.Unmarshal(byte, &d)

		repoByte, err := utils.Get(ctx, d.Repo)
		if err != nil {
			return err
		}
		json.Unmarshal(repoByte, &r)

		orgByte, err := utils.Get(ctx, d.Orgs)
		if err != nil {
			return err
		}
		json.Unmarshal(orgByte, &o)

		if err := write(username, response{
			data: d,
			repo: r,
			Org:  o,
		}); err != nil {
			return err
		}

	} else {
		return err
	}

	return nil
}

//function to write output in a file
func write(filename string, response response) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	s := fmt.Sprintf("DETAILS:: Id:%d, UserName:%s, Follower:%d, Following:%d \nREPO LIST:", response.data.Id, response.data.User, response.data.Followers, response.data.Following)
	f.WriteString(s)
	for _, s := range response.repo {
		s := fmt.Sprint("\t", s.Name)
		f.WriteString(s)
	}
	f.WriteString("\nOrganizations:\n")
	for _, s := range response.Org {
		name := fmt.Sprintf("\tName:%s\n", s.Name)
		f.WriteString(name)
		desc := fmt.Sprintf("\tDescription:%s\n", s.Description)
		f.WriteString(desc)
	}
	fmt.Println("stored in " + filename)
	return nil
}
