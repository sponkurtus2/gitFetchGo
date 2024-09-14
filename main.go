package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type RepoData struct {
	Name string `json:"name"`
	Url  string `json:"html_url"`
}

type UserData struct {
	UserName string `json:"login"`
	Photo    string `json:"avatar_url"`
}

func main() {
	userName := "sponkurtus2"
	token := "Bearer" + "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"

	// listRepos(userName, token)
	listUserProfile(userName, token)
}

func listRepos(userName, token string) {
	url := fmt.Sprintf("https://api.github.com/users/%s/repos", userName)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("Error: ", err)
	}

	req.Header.Add("Authorization", token)

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response...", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading...", err)
	}

	var repos []RepoData
	if err := json.Unmarshal(body, &repos); err != nil {
		log.Println("Error when Unmarshal json", err)
		return
	}

	if len(repos) > 3 {
		repos = repos[:3]
	}

	log.Println(repos)
}

func listUserProfile(userName, token string) {
	url := fmt.Sprintf("https://api.github.com/users/%s", userName)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("Error", err)
	}

	req.Header.Add("Authorization", token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response: ", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Errror while reading: ", err)
	}

	var user UserData
	if err := json.Unmarshal(body, &user); err != nil {
		log.Println("Error when unmarshal json", err)
		return
	}

	log.Println(user)
}
