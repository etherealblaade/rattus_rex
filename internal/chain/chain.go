package chain

import (
	"fmt"
	"github.com/etherealblaade/rattus_rex/internal/api"
	"os"
)

type ModelChain struct {
	DeepseekClient   *api.Client
	OpenRouterClient *api.Client
	ShowReasoning    bool
	History          []api.Message
	DeepseekModel    string
	OpenRouterModel  string
}

func NewModelChain() *ModelChain {
	return &ModelChain{
		DeepseekClient:   api.NewClient("https://api.deepseek.com", os.Getenv("DEEPSEEK_API_KEY")),
		OpenRouterClient: api.NewClient("https://openrouter.ai/api/v1", os.Getenv("OPENROUTER_API_KEY")),
		ShowReasoning:    true,
		History:          make([]api.Message, 0),
		DeepseekModel:    "deepseek-reasoner",
		OpenRouterModel:  "openai/gpt-4o-mini",
	}
}

func (mc *ModelChain) Process(input string) (string, error) {
	reasoning, err := mc.DeepseekClient.CreateCompletion(input, mc.DeepseekModel)
	if err != nil {
		return "", fmt.Errorf("deepseek error: %w", err)
	}

	if len(reasoning.Choices) == 0 {
		return "", fmt.Errorf("no response from deepseek")
	}

	response, err := mc.OpenRouterClient.CreateCompletion(
		input+"\n"+reasoning.Choices[0].Message.Content,
		mc.OpenRouterModel,
	)
	if err != nil {
		return "", fmt.Errorf("openrouter error: %w", err)
	}

	if len(response.Choices) == 0 {
		return "", fmt.Errorf("no response from openrouter")
	}

	return response.Choices[0].Message.Content, nil
}
