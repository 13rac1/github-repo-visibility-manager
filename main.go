package main

import (
	"context"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"gopkg.in/yaml.v2"
)

var pFalse = false
var pTrue = true

// Config stores the settings for which repositories to make public or private.
type Config struct {
	Organization string
	Repos        []struct {
		Name    string
		Private bool
	}
}

func main() {
	log.Print("Private/Public Github Repository Maintainer")
	ctx := context.Background()

	if len(os.Args) < 2 {
		log.Fatal("config file required")
	}

	// Load settings
	configFile := strings.TrimSpace(os.Args[1])
	if configFile == "" {
		log.Fatal("config file required")
	}

	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	config := Config{}
	err = yaml.Unmarshal([]byte(data), &config)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	githubToken := os.Getenv("GITHUB_TOKEN")
	if githubToken == "" {
		log.Fatal("GITHUB_TOKEN environment variable missing")
	}
	log.Print("found GITHUB_TOKEN")

	// Setup AuthN
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: githubToken},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	log.Printf("Organization: %s", config.Organization)

	for _, repo := range config.Repos {
		log.Printf("Repository: %s", repo.Name)

		repoRecord, _, err := client.Repositories.Get(ctx, config.Organization, repo.Name)
		if err != nil {
			log.Fatalf("error: %v", err)
		}

		log.Printf("found: %s", *repoRecord.FullName)

		if repo.Private == *repoRecord.Private {
			log.Printf("private already set to: %v, skipping...", repo.Private)
			continue
		}

		if repo.Private {
			log.Print("setting private")
			repoRecord.Private = &pTrue
		} else {
			log.Print("setting public")
			repoRecord.Private = &pFalse
		}

		repoUpdate, _, err := client.Repositories.Edit(ctx, config.Organization, repo.Name, repoRecord)
		if err != nil {
			log.Fatalf("error: %v", err)
		}

		log.Printf("private set to: %v", *repoUpdate.Private)
	}
	log.Print("done")
}
