package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type AiClient struct {
	Url    string
	Model  string
	AiType string
	ApiKey string
}

var AiType = map[string]*AiMetadata{
	"1": {
		Url:   "https://generativelanguage.googleapis.com/v1beta/openai/chat/completions",
		Model: "gemini-2.5-pro",
	}, // "gemini"
}

type AiMetadata struct {
	Url   string `json:"url"`
	Model string `json:"model"`
}

func NewAiClient(config *Config) *AiClient {
	metadata := AiType[config.AiType]

	return &AiClient{
		Url:    metadata.Url,
		Model:  metadata.Model,
		ApiKey: config.AiApiKey,
		AiType: config.AiType,
	}
}

// analyzeCodeWithAI sends the diff to OpenAI and returns AI feedback
func (c *AiClient) AnalyzeCode(diff string, customPrompt string) (string, error) {

	persona := "You're an expert software engineering assistant tasked with performing a detailed yet concise code review.\n. "

	modelPersona := "You're an expert software engineering assistant. Your tone is energetic, friendly, and slightly playful."
	context := "Code To Be Reviewed:" + diff
	protocol := "Instructions:\n" +
		"- Do not give positive comments or compliments.\n" +
		"- Provide comments and suggestions ONLY if there is something to improve.\n" +
		"- Use the given description only for the overall context and only comment the code.\n" +
		"- IMPORTANT: NEVER suggest adding comments to the code.\n" +
		"- Structure your feedback clearly and concisely using the following markdown format: \n" +
		formatMarkdown

	finalPrompt := "[MODEL]" +
		modelPersona +
		"[CONTEXT]" +
		context +
		"[PROTOCOL]" +
		protocol

	data := map[string]interface{}{
		"model": c.Model,
		"messages": []map[string]string{
			{"role": "system", "content": persona},
			{"role": "user", "content": finalPrompt},
		},
		"stream":     false,
		"max_tokens": 1000000,
	}

	// intercept custom prompt
	if len(customPrompt) > 0 {
		data["messages"] = []map[string]string{
			{"role": "system", "content": persona},
			{"role": "user", "content": customPrompt},
		}
	}

	jsonData, _ := json.Marshal(data)

	req, _ := http.NewRequest("POST", c.Url, bytes.NewBuffer(jsonData))
	req.Header.Set("Authorization", "Bearer "+c.ApiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("[AnalyzeCode] failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("[AnalyzeCode] API error: %s", resp.Status)
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("[AnalyzeCode] failed to decode response: %w", err)
	}

	return c.resolveResponse(result), nil
}

func (c *AiClient) resolveResponse(result map[string]interface{}) string {

	// custom ai reviewer model
	if c.AiType != "1" {
		message := result["message"].(map[string]interface{})
		return message["content"].(string)
	}

	choices := result["choices"].([]interface{})
	message := choices[0].(map[string]interface{})["message"].(map[string]interface{})
	return message["content"].(string)
}
