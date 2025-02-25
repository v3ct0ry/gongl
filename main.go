package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/google/uuid"
)

type NGLMessage struct {
	Username string `json:"username"`
	Question string `json:"question"`
	DeviceID string `json:"deviceId"`
}

func main() {
	fmt.Print("\033]0;goNGL by V3CT0RY | @v3ct0ry\007")
	clearConsole()
	bPrint("Welcome!\n")

	username := getValidUsername()

	deviceId := uuid.New().String()
	bPrint(fmt.Sprintf("Device ID generated: %s (used as identifier)\n", deviceId))

	message := bInput("Enter the desired message you want to send:")

	nglMsg := NGLMessage{
		Username: username,
		Question: message,
		DeviceID: deviceId,
	}
	jsonData, err := json.Marshal(nglMsg)
	if err != nil {
		log.Fatal("Error marshalling JSON:", err)
	}

	resp, err := http.Post("https://ngl.link/api/submit", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal("Error sending request:", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading response:", err)
	}

	if resp.StatusCode == http.StatusOK {
		bPrint("Message sent successfully!\n")
	} else {
		bPrint("Error sending message!\n")
	}
	bPrint("Response Body: " + string(body) + "\n")

	bPrint("Press ENTER to close...\n")
	bufio.NewReader(os.Stdin).ReadString('\n')
}

func getValidUsername() string {
	for {
		username := bInput("Enter the username:")
		resp, err := http.Get("https://ngl.link/@" + username)
		if err != nil {
			log.Println("Error checking username:", err)
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusOK {
			clearConsole()
			bPrint(fmt.Sprintf("Username @%s is valid\n", username))
			return username
		} else {
			clearConsole()
			bPrint("❌ Invalid/Unknown Username! Try again.\n")
		}
	}
}

func bPrint(msg string) {
	fmt.Print("\t\t" + msg)
}

func bInput(quest string) string {
	bPrint(quest + "\n\t\t-> ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func clearConsole() {
	fmt.Print("\033[H\033[2J\n\n\033[35m")
	banner := `
		 ██████╗  ██████╗ ███╗   ██╗ ██████╗ ██╗
		██╔════╝ ██╔═══██╗████╗  ██║██╔════╝ ██║ 
		██║  ███╗██║   ██║██╔██╗ ██║██║  ███╗██║     
		██║   ██║██║   ██║██║╚██╗██║██║   ██║██║     
		╚██████╔╝╚██████╔╝██║ ╚████║╚██████╔╝███████╗
		 ╚═════╝  ╚═════╝ ╚═╝  ╚═══╝ ╚═════╝ ╚══════╝
		 			=> V3CT0RY | @v3ct0ry`

	fmt.Println(banner)
	fmt.Printf("\033[36m\n\n")
}
