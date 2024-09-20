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
	// userName := "sponkurtus2"
	if len(os.Args) < 2 {
		fmt.Println("Please introduce a valid name.")
		os.Exit(1)
	}
	userName := os.Args[1]

	// Styles with fatih/color
	headerColor := color.New(color.FgHiMagenta).Add(color.Bold) // Header
	labelColor := color.New(color.FgHiWhite)                    // Label
	valueColor := color.New(color.FgHiCyan)                     // Values

	listRepos(userName, labelColor, valueColor)

	photoUrl := listUserProfile(userName, labelColor, valueColor)
	downloadPhoto(photoUrl)
	imgToAscii()

	headerColor.Println(strings.Repeat("─", 40))
}

func listRepos(userName string, labelColor, valueColor *color.Color) {
	url := fmt.Sprintf("https://api.github.com/users/%s/repos", userName)

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

	if len(repos) > 3 {
		// Only show the first 3 repos
		repos = repos[:3]
	}

	// Show repos on terminal
	labelColor.Println("Repositories:")
	for i, repo := range repos {
		labelColor.Printf("  %d. %s\n", i+1, "Name")
		valueColor.Printf("     → %s\n", repo.Name)
		labelColor.Printf("     %s\n", "URL")
		valueColor.Printf("     → %s\n", repo.Url)

		if i < len(repos)-1 {
			labelColor.Println(strings.Repeat("─", 30))
		}
	}
}

func listUserProfile(userName string, labelColor, valueColor *color.Color) string {
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

	// Mostrar la información del usuario al estilo neofetch
	labelColor.Println("User profile")
	labelColor.Printf("  %s\n", "User")
	valueColor.Printf("     → %s\n", user.UserName)

	// Retornar la URL de la foto
	return user.Photo
}

func imgToAscii() {
	// Execute command to convert and print ascii img
	cmd := exec.Command("image2ascii", "-f", "./userPhoto.jpg", "-w", "50", "-g", "30")

	out, err := cmd.Output()
	if err != nil {
		log.Fatal("Error doing the command -> ", err)
	}
	fmt.Println(string(out))

	// Deletes temp img
	deletePhoto()
}

func downloadPhoto(photoUrl string) {
	// Crear archivo temporal
	userPhotoFile, err := os.Create("userPhoto.jpg")
	if err != nil {
		log.Fatal("Couldn't create file image -> ", err)
	}
	defer userPhotoFile.Close()

	// Descargar la imagen desde la URL
	resp, err := http.Get(photoUrl)
	if err != nil {
		log.Fatal("Couldn't download image -> ", err)
	}
	defer resp.Body.Close()

	// Escribir los datos descargados en nuestro archivo
	_, err = io.Copy(userPhotoFile, resp.Body)
	if err != nil {
		log.Fatal("Error transfering photo data -> ", err)
	}
}

func deletePhoto() {
	file := os.Remove("./userPhoto.jpg")
	if file != nil {
		log.Fatal(file)
	}
}
