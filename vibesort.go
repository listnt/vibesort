package vibesort

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/invopop/jsonschema"
	openai "github.com/openai/openai-go/v2"
	"github.com/openai/openai-go/v2/option"
)

type SortedNumbers struct {
	Nums []int64
}

type vibesort struct {
	key string
}

func GenerateSchema[T any]() interface{} {
	reflector := jsonschema.Reflector{
		AllowAdditionalProperties: false,
		DoNotReference:            true,
	}
	var v T
	schema := reflector.Reflect(v)
	return schema
}

var SortedNumbersSchema = GenerateSchema[SortedNumbers]()

func NewVibesort(key string) *vibesort {
	return &vibesort{key: key}
}

func (v *vibesort) vibesort(nums []int64) ([]int64, error) {
	client := openai.NewClient(
		option.WithAPIKey(v.key),
	)

	// Define JSON schema for sorted numbers
	schemaParam := openai.ResponseFormatJSONSchemaJSONSchemaParam{
		Name:   "sorted_numbers",
		Schema: SortedNumbersSchema,
		Strict: openai.Bool(true),
	}

	chatCompletion, err := client.Chat.Completions.New(context.Background(), openai.ChatCompletionNewParams{
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(fmt.Sprintf("Response in form\n[1 2 3 ...] \nSort these numbers\n  %v", nums)),
		},
		Model: "gpt-5",

		ResponseFormat: openai.ChatCompletionNewParamsResponseFormatUnion{
			OfJSONSchema: &openai.ResponseFormatJSONSchemaParam{
				JSONSchema: schemaParam,
			},
		},
	})
	if err != nil {
		return nil, err
	}

	var sortedNumbers SortedNumbers
	_ = json.Unmarshal([]byte(chatCompletion.Choices[0].Message.Content), &sortedNumbers)

	return sortedNumbers.Nums, nil
}
