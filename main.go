package main

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
)

type RepoData struct {
	Name string `json:"name"`
	Url  string `json:"html_url"`
}

type UserData struct {
	UserName string `json:"login"`
	Photo    string `json:"avatar_url"`
}

// Escala de grises para mapear a caracteres ASCII
var asciiChars = []string{"@", "#", "S", "%", "?", "*", "+", ";", ":", ",", "."}

func main() {
	userName := "sponkurtus2"

	// Colores básicos usando fatih/color
	pink := color.New(color.FgHiMagenta)

	// Ejemplo de salida estilizada
	pink.Println("Hola, bienvenido a mi neofetch en Go!")
	pink.Println("------------------------------")

	listRepos(userName)
	// listUserProfile(userName, token)
	downloadPhoto("https://avatars.githubusercontent.com/u/74841175?v=4")
	imgToAscii("./userPhoto.jpg")

}

func listRepos(userName string) {
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
		repos = repos[:3]
	}

	// Definir colores usando fatih/color
	nameColor := color.New(color.FgCyan).Add(color.Bold) // Nombre del repositorio
	urlColor := color.New(color.FgYellow)                // URL del repositorio
	separatorColor := color.New(color.FgMagenta)         // Separador para estética

	for i, repo := range repos {
		nameColor.Printf("%d. Name -> %s\n", i+1, repo.Name)
		urlColor.Printf("	URL -> %s\n", repo.Url)

		if i < len(repos)-1 {
			separatorColor.Println("------------------------------")
		}
	}

}

func listUserProfile(userName string) {
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
		log.Println("Errror while reading: ", err)
	}

	var user UserData
	if err := json.Unmarshal(body, &user); err != nil {
		log.Println("Error when unmarshal json", err)
		return
	}

	// Define colors
	greenColor := color.New(color.FgGreen)

	greenColor.Println(user.UserName, user.Photo)
}

func imgToAscii(photoFile string) {
	// cmd := exec.Command("image2ascii -f userPhoto.jpg -w 35 -g 15")
	cmd := exec.Command("image2ascii", "-f", photoFile, "-w", "35", "-g", "15")

	out, err := cmd.Output()
	if err != nil {
		log.Fatal("Error doing the command -> ", err)
	}
	fmt.Println(string(out))

	deletePhoto()
}

func downloadPhoto(photoUrl string) {

	// Create temp file
	userPhotoFile, err := os.Create("userPhoto.jpg")
	if err != nil {
		log.Fatal("Couldn't create file image -> ", err)
	}
	defer userPhotoFile.Close()

	// Get the photo data
	resp, err := http.Get(photoUrl)
	if err != nil {
		log.Fatal("Couldn't download image -> ", err)
	}
	defer resp.Body.Close()

	// Write downloaded data into our file
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
