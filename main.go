package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

//  let api_key: &str = "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";
//
//    let url = format!(
//        "https://api.github.com/users/${user_name}"
//    );

type ResponseData struct {
	Results []interface{} `json:"results"`
}

type RepoData struct {
	Name string `json:"name"`
	Url  string `json:"html_url"`
}

func main() {
	userName := "sponkurtus2"
	url := fmt.Sprintf("https://api.github.com/users/%s", userName)

	// var token = "Bearer" + "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"
	token := "Bearer" + "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"

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

	_, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading", err)
	}
	// log.Println(string([]byte(body)))

	list_repos(userName, token)
}

func list_repos(userName, token string) {
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
