package main

import (
	"bytes"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"strings"
)

func init() {
	log.SetFlags(0)
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("PR number or URL required")
	}

	var body string
	if len(os.Args) > 2 {
		body = " " + strings.Join(os.Args[2:], " ")
	} else {
		body = emojis[rand.Intn(len(emojis))]
	}

	var buf bytes.Buffer
	pr := os.Args[1]
	cmd := exec.Command("gh", "pr", "review", pr, "--body", body, "--approve")
	cmd.Stderr = &buf

	err := cmd.Run()
	if err != nil {
		if s := buf.String(); s != "" {
			log.Fatal(s)
		}
		log.Fatal(err)
	}

	log.Printf("PR %s approved with %s", pr, body)
}

var emojis = []string{"✅"}
