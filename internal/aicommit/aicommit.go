package aicommit

import (
	"context"
	gogpt "github.com/sashabaranov/go-gpt3"
)

type AiCommit struct {
	ID   int
	Diff []byte
}

func New(id int, diff []byte) *AiCommit {
	return &AiCommit{ID: id, Diff: diff}
}

func (a *AiCommit) GenerateCommitMessage() (string, error) {
	gptClient := gogpt.NewClient("YOUR_OPENAI_API_KEY")
	ctx := context.Background()
	prompt := "Lisez le git diff suivant pour plusieurs fichiers :" + string(a.Diff) + "Générez 1 à 3 paragraphes pour expliquer cette différence à un humain sans mentionner les changements eux-mêmes."
	req := gogpt.CompletionRequest{
		Model:     gogpt.GPT3TextDavinci002,
		MaxTokens: 100,
		Prompt:    prompt,
	}
	resp, err := gptClient.CreateCompletion(ctx, req)
	if err != nil {
		return "", err
	}
	return resp.Choices[0].Text, nil
}
