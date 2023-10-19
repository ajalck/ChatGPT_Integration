package models

type PromptParams struct {
	Model        string  `json:"model"`
	Prompt       string  `json:"prompt"`
	MaxTokens    int     `json:"max_tokens,omitempty`
}
