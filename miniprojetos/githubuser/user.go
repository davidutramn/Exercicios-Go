package main

import (
	"encoding/json"
	"net/http"
)

const GitHubUsersURL = "https://api.github.com/users/"

type GitHubUser struct {
	Login           string `json:"login"`
	ID              int    `json:"id"`
	NodeID          string `json:"node_id"`
	AvatarURL       string `json:"avatar_url"`
	GravatarID      string `json:"gravatar_id"`
	Type            string `json:"type"`
	SiteAdmin       bool   `json:"site_admin"`
	Name            string `json:"name"`
	Company         string `json:"company"`
	Blog            string `json:"blog"`
	Location        string `json:"location"`
	Email           string `json:"email"`
	Hireable        string `json:"hireable"`
	Bio             string `json:"bio"`
	TwitterUsername string `json:"twitter_username"`
	PublicRepos     int    `json:"public_repos"`
	PublicGists     int    `json:"public_gists"`
	Followers       int    `json:"followers"`
	Following       int    `json:"following"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`

	URL               string `json:"url"`
	HTMLURL           string `json:"html_url"`
	FollowersURL      string `json:"followers_url"`
	FollowingURL      string `json:"following_url"`
	StarredURL        string `json:"starred_url"`
	SubscriptionsURL  string `json:"subscriptions_url"`
	OrganizationsURL  string `json:"organizations_url"`
	ReposURL          string `json:"repos_url"`
	EventsURL         string `json:"events_url"`
	ReceivedEventsURL string `json:"received_events_url"`
}

func getUser(username string) (*GitHubUser, error) {
	resp, err := http.Get(GitHubUsersURL + username)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var user GitHubUser
	err = json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
