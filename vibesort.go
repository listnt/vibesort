package vibesort

import (
	"context"
	"fmt"
	openai "github.com/openai/openai-go/v2"
	"github.com/openai/openai-go/v2/option"
	"strconv"
	"strings"
)

type vibesort struct {
	key string
}

func NewVibesort(key string) *vibesort {
	return &vibesort{key: key}
}

func (v *vibesort) vibesort(nums []int64) ([]int64, error) {
	client := openai.NewClient(
		option.WithAPIKey(v.key),
	)

	chatCompletion, err := client.Chat.Completions.New(context.Background(), openai.ChatCompletionNewParams{
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(fmt.Sprintf("Response in form\n[1 2 3 ...] \nSort these numbers\n  %v", nums)),
		},
		Model: "gpt-5",
	})
	if err != nil {
		return nil, err
	}

	s := strings.Trim(chatCompletion.Choices[0].Message.Content, "[]")
	parts := strings.Fields(s)

	nums = make([]int64, len(parts))
	for i, p := range parts {
		n, err := strconv.Atoi(p)
		if err != nil {
			return nil, err
		}
		nums[i] = int64(n)
	}

	return nums, nil
}
