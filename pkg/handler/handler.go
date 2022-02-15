package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/rohandas-max/ghCrwaler/pkg/utils"
)

type data struct {
	User      string `json:"login"`
	Followers int    `json:"followers"`
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

func Handler(ctx context.Context, userName string, t time.Duration) {
	select {
	case <-ctx.Done():
		log.Fatal(ctx.Err().Error())
	case <-time.After(t):
		url := "http://api.github.com/users/" + userName
		resp := utils.Get(ctx, url)
		defer resp.Body.Close()
		if resp.StatusCode == http.StatusOK {
			var d data
			json.NewDecoder(resp.Body).Decode(&d)

			res := utils.Get(ctx, d.Repo)
			defer res.Body.Close()
			var r []repo
			json.NewDecoder(res.Body).Decode(&r)

			resO := utils.Get(ctx, d.Orgs)
			defer resO.Body.Close()
			var O []org
			if err := json.NewDecoder(resO.Body).Decode(&O); err != nil {
				panic(err)
			}

			write(userName+".txt", response{
				data: data{
					User:      d.User,
					Followers: d.Followers,
					Repo:      d.Repo,
					Orgs:      d.Orgs,
				},
				repo: r,
				Org:  O,
			})
		} else {
			log.Fatal(http.StatusNotFound)
		}
	}
}

//function to write output in a file
func write(name string, response response) {
	f, err := os.Create(name)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	s := fmt.Sprintf("DETAILS:: UserName:%s, Follower:%s, \nREPO LIST:", response.data.User, strconv.Itoa(response.data.Followers))
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
}
