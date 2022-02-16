package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/rohandas-max/ghCrwaler/pkg/utils"
)

type data struct {
	Id        int    `json:"id"`
	User      string `json:"login"`
	Followers int    `json:"followers"`
	Following int    `json:"following"`
	Repo      string `json:"repos_url"`
	Orgs      string `json:"organizations_url"`
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
	var data data
	var repo []repo
	var org []org
	if byte, err := utils.Get(ctx, url); err == nil {
		json.Unmarshal(byte, &data)
		repoByte, err := utils.Get(ctx, data.Repo)
		if err != nil {
			return err
		}
		json.Unmarshal(repoByte, &repo)

		orgByte, err := utils.Get(ctx, data.Orgs)
		if err != nil {
			return err
		}
		json.Unmarshal(orgByte, &org)

		if err := writeToFile(username, response{
			data: data,
			repo: repo,
			Org:  org,
		}); err != nil {
			return err
		}

	} else {
		return err
	}

	return nil
}

//function to write output in a text file
func writeToFile(filename string, response response) error {
	if filename == "" {
		return errors.New("no filename entered")
	}
	f, err := os.Create(filename + ".txt")
	if err != nil {
		return err
	}
	defer f.Close()

	var repoName []string
	var orgName []string
	var orgDes []string
	for _, n := range response.repo {
		repoName = append(repoName, n.Name)
	}
	for _, r := range response.Org {
		orgName = append(orgName, r.Name)
		orgDes = append(orgDes, r.Description)
	}

	s := fmt.Sprintf("DETAILS:: Id:%d, UserName:%s, Follower:%d, Following:%d \nREPO LIST:%v,\nOrganization:\n\tName:%v\n\tDescription:%v",
		response.data.Id, response.data.User, response.data.Followers,
		response.data.Following, repoName, orgName, orgDes)

	f.WriteString(s)
	fmt.Println("stored in " + filename + ".txt")
	return nil
}
