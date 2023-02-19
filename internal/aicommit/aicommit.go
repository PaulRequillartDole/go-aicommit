package aicommit

import (
	"context"
	"github.com/joho/godotenv"
	gogpt "github.com/sashabaranov/go-gpt3"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

type AiCommit struct {
	ID   int
	Diff []byte
}

var (
	_, b, _, _ = runtime.Caller(0)
	// Root folder of this project
	Root = filepath.Join(filepath.Dir(b), "../..")
)

func New(id int, diff []byte) *AiCommit {
	return &AiCommit{ID: id, Diff: diff}
}

func goDotEnvVariable(key string) string {

	environmentPath := filepath.Join(Root, ".env")
	err := godotenv.Load(environmentPath)
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

func (a *AiCommit) GenerateCommitMessage() (string, error) {
	gptClient := gogpt.NewClient(goDotEnvVariable("OPENAI_API_KEY"))
	ctx := context.Background()
	prompt := goDotEnvVariable("BEFORE_PROMPT") + string(a.Diff)
	req := gogpt.CompletionRequest{
		Model:            gogpt.GPT3TextDavinci003,
		MaxTokens:        300,
		Prompt:           prompt,
		Temperature:      0.7,
		TopP:             1,
		FrequencyPenalty: 0,
		PresencePenalty:  0,
		Stream:           false,
	}
	resp, err := gptClient.CreateCompletion(ctx, req)
	if err != nil {
		return "", err
	}
	return strings.ReplaceAll(resp.Choices[0].Text, "\n", ""), nil
}
