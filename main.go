package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/fatih/color"
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
	if len(os.Args) < 2 {
		color.Red("Please introduce a valid GitHub username.")
		os.Exit(1)
	}
	userName := os.Args[1]

	// Styles with fatih/color
	titleColor := color.New(color.FgHiMagenta).Add(color.Bold)
	labelColor := color.New(color.FgHiBlue)
	valueColor := color.New(color.FgHiWhite)

	// Create channels for coordination
	photoUrlChan := make(chan string)
	doneChan := make(chan bool)

	// Start concurrent operations
	go func() {
		photoUrl := listUserProfile(userName, titleColor, labelColor, valueColor)
		photoUrlChan <- photoUrl
	}()

	go func() {
		listRepos(userName, labelColor, valueColor)
		doneChan <- true
	}()

	// Print ASCII art first (if exists from previous run)
	fmt.Print("\n")
	imgToAscii()

	// Wait for user profile
	photoUrl := <-photoUrlChan

	// Wait for repos to complete
	<-doneChan
	fmt.Print("\n")

	// Create a channel for image processing
	imgDoneChan := make(chan bool)

	// Download and process image concurrently
	go func() {
		downloadPhoto(photoUrl)
		imgToAscii()
		imgDoneChan <- true
	}()

	// Wait for image processing to complete
	<-imgDoneChan
	deletePhoto()
}

func listRepos(userName string, labelColor, valueColor *color.Color) {
	url := fmt.Sprintf("https://api.github.com/users/%s/repos?per_page=3", userName)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("Error: ", err)
	}

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

	icon := color.HiYellowString("󰊢")
	labelColor.Printf(" %s ", icon)
	labelColor.Print(color.HiBlueString("repos"))
	valueColor.Print("  ")

	for i, repo := range repos {
		if i == 0 {
			color.HiWhite(repo.Name)
		} else {
			valueColor.Printf(" • %s", color.HiWhiteString(repo.Name))
		}
	}
	fmt.Print("\n")
}

func listUserProfile(userName string, titleColor, labelColor, valueColor *color.Color) string {
	url := fmt.Sprintf("https://api.github.com/users/%s", userName)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("Error", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response: ", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading: ", err)
	}

	var user UserData
	if err := json.Unmarshal(body, &user); err != nil {
		log.Println("Error when unmarshal json", err)
	}

	// Print username as title
	titleColor.Printf("\n  %s\n", strings.ToUpper(user.UserName))
	titleColor.Printf("  %s\n", strings.Repeat("─", len(user.UserName)))

	return user.Photo
}

func imgToAscii() {
	cmd := exec.Command("image2ascii", "-f", "./userPhoto.jpg", "-w", "35", "-g", "20")
	out, err := cmd.Output()
	if err != nil {
		return
	}
	color.HiCyan(string(out))
}

func downloadPhoto(photoUrl string) {
	// Create temporary file
	userPhotoFile, err := os.Create("userPhoto.jpg")
	if err != nil {
		log.Printf("Couldn't create file image -> %v", err)
		return
	}
	defer userPhotoFile.Close()

	// Download the image from URL
	resp, err := http.Get(photoUrl)
	if err != nil {
		log.Printf("Couldn't download image -> %v", err)
		return
	}
	defer resp.Body.Close()

	// Write downloaded data to file
	_, err = io.Copy(userPhotoFile, resp.Body)
	if err != nil {
		log.Printf("Error transfering photo data -> %v", err)
		return
	}
}

func deletePhoto() {
	if err := os.Remove("./userPhoto.jpg"); err != nil && !os.IsNotExist(err) {
		log.Printf("Error deleting photo: %v", err)
	}
}
