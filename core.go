package vibes

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/openai/openai-go/v2"
	"github.com/openai/openai-go/v2/option"

	"github.com/invopop/jsonschema"
)

// Inits package, have to be called before any other methods
func InitVibes(key string) {
	g_key = key
}

type AnyArray[T any] struct {
	Data []T
}

type SingleValue[T any] struct {
	Data T
}

// GenerateSchema is a function to generate openAI json schema used in structured output
func GenerateSchema[T any]() interface{} {
	reflector := jsonschema.Reflector{
		AllowAdditionalProperties: false,
		DoNotReference:            true,
	}
	var v T
	schema := reflector.Reflect(v)
	return schema
}

var g_key string

func call[T any](prompt string) (T, error) {
	var zero T
	if g_key == "" {
		return zero, fmt.Errorf("OpenAI key is not initialized")
	}

	client := openai.NewClient(
		option.WithAPIKey(g_key),
	)

	singleValueSchema := GenerateSchema[SingleValue[T]]()

	schemaParam := openai.ResponseFormatJSONSchemaJSONSchemaParam{
		Name:   "single_value",
		Schema: singleValueSchema,
		Strict: openai.Bool(true),
	}

	chatCompletion, err := client.Chat.Completions.New(context.Background(), openai.ChatCompletionNewParams{
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(prompt),
		},
		Model: "gpt-5",

		ResponseFormat: openai.ChatCompletionNewParamsResponseFormatUnion{
			OfJSONSchema: &openai.ResponseFormatJSONSchemaParam{
				JSONSchema: schemaParam,
			},
		},
	})
	if err != nil {
		return zero, err
	}

	var singleValue SingleValue[T]
	_ = json.Unmarshal([]byte(chatCompletion.Choices[0].Message.Content), &singleValue)

	return singleValue.Data, nil
}

func callArray[T any](prompt string) ([]T, error) {
	if g_key == "" {
		return nil, fmt.Errorf("OpenAI key is not initialized")
	}

	client := openai.NewClient(
		option.WithAPIKey(g_key),
	)

	jsonSchema := GenerateSchema[AnyArray[T]]()

	schemaParam := openai.ResponseFormatJSONSchemaJSONSchemaParam{
		Name:   "values",
		Schema: jsonSchema,
		Strict: openai.Bool(true),
	}

	chatCompletion, err := client.Chat.Completions.New(context.Background(), openai.ChatCompletionNewParams{
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(prompt),
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

	var values AnyArray[T]
	_ = json.Unmarshal([]byte(chatCompletion.Choices[0].Message.Content), &values)

	return values.Data, nil
}
