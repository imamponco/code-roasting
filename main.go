package main

import (
	"encoding/json"
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"log"
)

func main() {
	config := new(Config)
	err := envconfig.Process("", config)
	if err != nil {
		log.Fatalf("Error fetching data config: %v", err)
	}

	if config.Debug {
		configBytes, err := json.Marshal(config)
		if err != nil {
			return
		}
		log.Printf("Debug config: %v", string(configBytes))
	}

	gitlab := NewGitlabClient(config)
	ai := NewAiClient(config)

	// Step 1: Fetch open merge requests
	mrs, err := gitlab.GetOpenMergeRequests()
	if err != nil {
		log.Fatalf("Error fetching merge requests: %v", err)
	}

	if len(mrs) == 0 {
		fmt.Println("No open merge requests found.")
		return
	}

	// Step 2: Get the diff of the MR
	diff, err := gitlab.GetMergeRequestDiff()
	if err != nil {
		log.Fatalf("Error fetching merge request diff: %v", err)
	}

	// Step 3: Analyze the diff with AI
	feedback, err := ai.AnalyzeCode(diff, config.AiCustomPrompt)
	if err != nil {
		log.Fatalf("Error analyzing code with AI: %v", err)
	}

	// Step 4: Post the feedback as a comment
	err = gitlab.PostComment(feedback)
	if err != nil {
		log.Fatalf("Error posting comment: %v", err)
	}

	fmt.Printf("AI feedback posted on MR #%d\n", config.GitlabMrId)
}
