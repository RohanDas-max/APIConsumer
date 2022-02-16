package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

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
	resp := utils.Get(ctx, url)
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		var d data
		json.NewDecoder(resp.Body).Decode(&d)

		// reading values from repository_url
		res := utils.Get(ctx, d.Repo)
		defer res.Body.Close()
		var r []repo
		json.NewDecoder(res.Body).Decode(&r)

		// reading values from Organizations_url
		resO := utils.Get(ctx, d.Orgs)
		defer resO.Body.Close()
		var O []org
		json.NewDecoder(resO.Body).Decode(&O)

		write(username+".txt", response{
			data: data{
				Id:        d.Id,
				User:      d.User,
				Followers: d.Followers,
				Following: d.Following,
				Repo:      d.Repo,
				Orgs:      d.Orgs,
			},
			repo: r,
			Org:  O,
		})
	} else {
		return errors.New(strconv.Itoa(http.StatusNotFound))

	}
	return nil
}

//function to write output in a file
func write(filename string, response response) {
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
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
}
