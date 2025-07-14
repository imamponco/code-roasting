package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "code-roasting",
		Usage: "AI-powered code review for GitLab merge requests",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "gitlab-base-url",
				Usage:   "GitLab base URL",
				EnvVars: []string{"GITLAB_BASE_URL"},
			},
			&cli.StringFlag{
				Name:    "gitlab-project-id",
				Usage:   "GitLab project ID",
				EnvVars: []string{"GITLAB_PROJECT_ID"},
			},
			&cli.IntFlag{
				Name:    "gitlab-mr-id",
				Usage:   "GitLab merge request ID",
				EnvVars: []string{"GITLAB_MR_ID"},
			},
			&cli.StringFlag{
				Name:    "gitlab-access-token",
				Usage:   "GitLab access token",
				EnvVars: []string{"GITLAB_ACCESS_TOKEN"},
			},
			&cli.StringFlag{
				Name:    "ai-type",
				Value:   "1",
				Usage:   "AI type",
				EnvVars: []string{"AI_TYPE"},
			},
			&cli.StringFlag{
				Name:    "ai-api-key",
				Usage:   "AI API key",
				EnvVars: []string{"AI_API_KEY"},
			},
			&cli.StringFlag{
				Name:    "ai-custom-prompt",
				Usage:   "Custom prompt for the AI",
				EnvVars: []string{"AI_CUSTOM_PROMPT"},
			},
			&cli.BoolFlag{
				Name:    "debug",
				Usage:   "Enable debug mode",
				EnvVars: []string{"DEBUG"},
			},
		},
		Action: func(c *cli.Context) error {
			config := &Config{
				GitLabBaseURL:     c.String("gitlab-base-url"),
				GitlabProjectId:   c.String("gitlab-project-id"),
				GitlabMrId:        c.Int("gitlab-mr-id"),
				GitlabAccessToken: c.String("gitlab-access-token"),
				AiType:            c.String("ai-type"),
				AiApiKey:          c.String("ai-api-key"),
				AiCustomPrompt:    c.String("ai-custom-prompt"),
				Debug:             c.Bool("debug"),
			}
			return run(config)
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func run(config *Config) error {
	if config.Debug {
		configBytes, err := json.Marshal(config)
		if err != nil {
			return fmt.Errorf("[run] failed to marshal config: %w", err)
		}
		log.Printf("Debug config: %v", string(configBytes))
	}

	gitlab := NewGitlabClient(config)
	ai := NewAiClient(config)

	// Step 1: Fetch open merge requests
	mrs, err := gitlab.GetOpenMergeRequests()
	if err != nil {
		return fmt.Errorf("[run] failed to get open merge requests: %w", err)
	}

	if len(mrs) == 0 {
		fmt.Println("No open merge requests found.")
		return nil
	}

	// Step 2: Get the diff of the MR
	diff, err := gitlab.GetMergeRequestDiff()
	if err != nil {
		return fmt.Errorf("[run] failed to get merge request diff: %w", err)
	}

	// Step 3: Analyze the diff with AI
	feedback, err := ai.AnalyzeCode(diff, config.AiCustomPrompt)
	if err != nil {
		return fmt.Errorf("[run] failed to analyze code: %w", err)
	}

	// Step 4: Post the feedback as a comment
	err = gitlab.PostComment(feedback)
	if err != nil {
		return fmt.Errorf("[run] failed to post comment: %w", err)
	}

	fmt.Printf("AI feedback posted on MR #%d\n", config.GitlabMrId)
	return nil
}
