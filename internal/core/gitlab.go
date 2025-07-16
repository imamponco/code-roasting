package core

import (
	"fmt"
	"github.com/go-resty/resty/v2"
)

type GitlabClient struct {
	REST      *resty.Client
	ProjectId string
	MrId      int
}

func NewGitlabClient(config *Config) *GitlabClient {
	client := resty.New().
		SetBaseURL(config.GitLabBaseURL).
		SetAuthToken(config.GitlabAccessToken)

	return &GitlabClient{
		REST:      client,
		ProjectId: config.GitlabProjectId,
		MrId:      config.GitlabMrId,
	}
}

// getOpenMergeRequests fetches all open merge requests
func (c *GitlabClient) GetOpenMergeRequests() ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	resp, err := c.REST.R().
		SetResult(&result).
		Get(fmt.Sprintf("/projects/%s/merge_requests?state=opened", c.ProjectId))
	if err != nil {
		return nil, fmt.Errorf("[GetOpenMergeRequests] failed to get merge requests: %w", err)
	}

	if resp.IsError() {
		return nil, fmt.Errorf("[GetOpenMergeRequests] API error: %s", resp.String())
	}

	return result, nil
}

// postComment posts a comment on a merge request
func (c *GitlabClient) PostComment(comment string) error {
	body := map[string]string{"body": comment}
	resp, err := c.REST.R().
		SetBody(body).
		Post(fmt.Sprintf("/projects/%s/merge_requests/%d/notes", c.ProjectId, c.MrId))
	if err != nil {
		return fmt.Errorf("[PostComment] failed to post comment: %w", err)
	}

	if resp.IsError() {
		return fmt.Errorf("[PostComment] API error: %s", resp.String())
	}

	return nil
}

// getMergeRequestDiff fetches the diff of a merge request
func (c *GitlabClient) GetMergeRequestDiff() (string, error) {
	// Define a structure to hold the API response
	var result struct {
		Changes []struct {
			Diff string `json:"diff"`
		} `json:"changes"`
	}

	resp, err := c.REST.R().
		SetResult(&result). // Automatically unmarshals JSON into `result`
		Get(fmt.Sprintf("/projects/%s/merge_requests/%d/changes", c.ProjectId, c.MrId))
	if err != nil {
		return "", fmt.Errorf("[GetMergeRequestDiff] failed to get merge request diff: %w", err)
	}

	if resp.IsError() {
		return "", fmt.Errorf("[GetMergeRequestDiff] API error: %s", resp.String())
	}

	// Combine all diffs into a single string
	diff := ""
	for _, change := range result.Changes {
		diff += change.Diff
	}
	return diff, nil
}
