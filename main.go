package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const Username = "TacoBot"

var (
	token        string
	levelPattern = regexp.MustCompile(`.*level (\d+).*`)
)

func HandleWebhook(w http.ResponseWriter, r *http.Request) {
	if r.FormValue("token") != token {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("invalid token"))
		return
	}
	username := r.FormValue("user_name")
	if username == "slackbot" {
		return
	}
	text := strings.ToLower(r.FormValue("text"))
	fmt.Printf("%s: %s\n", username, text)

	match := levelPattern.FindStringSubmatch(text)
	if len(match) < 2 {
		// Bark twice if someone mentions taco!
		if strings.Contains(text, "taco") {
			Respond(w, "Bark! Bark!")
		}
		return
	}
	// Print the solution to a level if someone says "level %d".
	level64, err := strconv.ParseInt(match[1], 10, 8)
	if err != nil {
		return
	}
	level := int(level64)
	if !(1 <= level && level <= len(Levels)) {
		return
	}
	path := Solve(Levels[level-1])
	text = fmt.Sprintf("The solution to level %d is %s", level, path[0])
	for _, dir := range path[1:] {
		text += fmt.Sprintf(", %s", dir)
	}
	text += "."
	Respond(w, text)
}

func Respond(w http.ResponseWriter, text string) {
	fmt.Printf("%s: %s\n", Username, text)
	buf, err := json.Marshal(map[string]string{
		"user_name": Username,
		"text":      text,
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		fmt.Fprintf(os.Stderr, err.Error())
		return
	}
	w.Write(buf)
}

func main() {
	if token = os.Getenv("TOKEN"); token == "" {
		log.Fatal("TOKEN not specified")
	}
	http.HandleFunc("/", HandleWebhook)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
