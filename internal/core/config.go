package core

// Config defines structure of Application configuration
type Config struct {
	NodeId            string `envconfig:"NODE_ID"`
	WorkDir           string `envconfig:"WORK_DIR"`
	Debug             bool   `envconfig:"DEBUG" default:"false"`
	GitLabBaseURL     string `envconfig:"GITLAB_BASE_URL"`
	GitlabProjectId   string `envconfig:"GITLAB_PROJECT_ID"`
	GitlabMrId        int    `envconfig:"GITLAB_MR_ID"`
	GitlabAccessToken string `envconfig:"GITLAB_ACCESS_TOKEN"`
	AiType            string `envconfig:"AI_TYPE" default:"1"`
	AiApiKey          string `envconfig:"AI_API_KEY"`
	AiCustomPrompt    string `envconfig:"AI_CUSTOM_PROMPT"`
}
