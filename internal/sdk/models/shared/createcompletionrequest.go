// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/antonbabenko/terraform-provider-openai/v2/internal/sdk/internal/utils"
)

// CreateCompletionRequestModel - ID of the model to use. You can use the [List models](/docs/api-reference/models/list) API to see all of your available models, or see our [Model overview](/docs/models/overview) for descriptions of them.
type CreateCompletionRequestModel string

const (
	CreateCompletionRequestModelTextDavinci003 CreateCompletionRequestModel = "text-davinci-003"
	CreateCompletionRequestModelTextDavinci002 CreateCompletionRequestModel = "text-davinci-002"
	CreateCompletionRequestModelTextDavinci001 CreateCompletionRequestModel = "text-davinci-001"
	CreateCompletionRequestModelCodeDavinci002 CreateCompletionRequestModel = "code-davinci-002"
	CreateCompletionRequestModelTextCurie001   CreateCompletionRequestModel = "text-curie-001"
	CreateCompletionRequestModelTextBabbage001 CreateCompletionRequestModel = "text-babbage-001"
	CreateCompletionRequestModelTextAda001     CreateCompletionRequestModel = "text-ada-001"
)

func (e CreateCompletionRequestModel) ToPointer() *CreateCompletionRequestModel {
	return &e
}
func (e *CreateCompletionRequestModel) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "text-davinci-003":
		fallthrough
	case "text-davinci-002":
		fallthrough
	case "text-davinci-001":
		fallthrough
	case "code-davinci-002":
		fallthrough
	case "text-curie-001":
		fallthrough
	case "text-babbage-001":
		fallthrough
	case "text-ada-001":
		*e = CreateCompletionRequestModel(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateCompletionRequestModel: %v", v)
	}
}

type PromptType string

const (
	PromptTypeStr                   PromptType = "str"
	PromptTypeArrayOfStr            PromptType = "arrayOfStr"
	PromptTypeArrayOfInteger        PromptType = "arrayOfInteger"
	PromptTypeArrayOfArrayOfInteger PromptType = "arrayOfArrayOfInteger"
)

// Prompt - The prompt(s) to generate completions for, encoded as a string, array of strings, array of tokens, or array of token arrays.
//
// Note that <|endoftext|> is the document separator that the model sees during training, so if a prompt is not specified the model will generate as if from the beginning of a new document.
type Prompt struct {
	Str                   *string
	ArrayOfStr            []string
	ArrayOfInteger        []int64
	ArrayOfArrayOfInteger [][]int64

	Type PromptType
}

func CreatePromptStr(str string) Prompt {
	typ := PromptTypeStr

	return Prompt{
		Str:  &str,
		Type: typ,
	}
}

func CreatePromptArrayOfStr(arrayOfStr []string) Prompt {
	typ := PromptTypeArrayOfStr

	return Prompt{
		ArrayOfStr: arrayOfStr,
		Type:       typ,
	}
}

func CreatePromptArrayOfInteger(arrayOfInteger []int64) Prompt {
	typ := PromptTypeArrayOfInteger

	return Prompt{
		ArrayOfInteger: arrayOfInteger,
		Type:           typ,
	}
}

func CreatePromptArrayOfArrayOfInteger(arrayOfArrayOfInteger [][]int64) Prompt {
	typ := PromptTypeArrayOfArrayOfInteger

	return Prompt{
		ArrayOfArrayOfInteger: arrayOfArrayOfInteger,
		Type:                  typ,
	}
}

func (u *Prompt) UnmarshalJSON(data []byte) error {

	var str string = ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = PromptTypeStr
		return nil
	}

	var arrayOfStr []string = []string{}
	if err := utils.UnmarshalJSON(data, &arrayOfStr, "", true, true); err == nil {
		u.ArrayOfStr = arrayOfStr
		u.Type = PromptTypeArrayOfStr
		return nil
	}

	var arrayOfInteger []int64 = []int64{}
	if err := utils.UnmarshalJSON(data, &arrayOfInteger, "", true, true); err == nil {
		u.ArrayOfInteger = arrayOfInteger
		u.Type = PromptTypeArrayOfInteger
		return nil
	}

	var arrayOfArrayOfInteger [][]int64 = [][]int64{}
	if err := utils.UnmarshalJSON(data, &arrayOfArrayOfInteger, "", true, true); err == nil {
		u.ArrayOfArrayOfInteger = arrayOfArrayOfInteger
		u.Type = PromptTypeArrayOfArrayOfInteger
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for Prompt", string(data))
}

func (u Prompt) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.ArrayOfStr != nil {
		return utils.MarshalJSON(u.ArrayOfStr, "", true)
	}

	if u.ArrayOfInteger != nil {
		return utils.MarshalJSON(u.ArrayOfInteger, "", true)
	}

	if u.ArrayOfArrayOfInteger != nil {
		return utils.MarshalJSON(u.ArrayOfArrayOfInteger, "", true)
	}

	return nil, errors.New("could not marshal union type Prompt: all fields are null")
}

type CreateCompletionRequestStopType string

const (
	CreateCompletionRequestStopTypeStr        CreateCompletionRequestStopType = "str"
	CreateCompletionRequestStopTypeArrayOfStr CreateCompletionRequestStopType = "arrayOfStr"
)

// CreateCompletionRequestStop - Up to 4 sequences where the API will stop generating further tokens. The returned text will not contain the stop sequence.
type CreateCompletionRequestStop struct {
	Str        *string
	ArrayOfStr []string

	Type CreateCompletionRequestStopType
}

func CreateCreateCompletionRequestStopStr(str string) CreateCompletionRequestStop {
	typ := CreateCompletionRequestStopTypeStr

	return CreateCompletionRequestStop{
		Str:  &str,
		Type: typ,
	}
}

func CreateCreateCompletionRequestStopArrayOfStr(arrayOfStr []string) CreateCompletionRequestStop {
	typ := CreateCompletionRequestStopTypeArrayOfStr

	return CreateCompletionRequestStop{
		ArrayOfStr: arrayOfStr,
		Type:       typ,
	}
}

func (u *CreateCompletionRequestStop) UnmarshalJSON(data []byte) error {

	var str string = ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = CreateCompletionRequestStopTypeStr
		return nil
	}

	var arrayOfStr []string = []string{}
	if err := utils.UnmarshalJSON(data, &arrayOfStr, "", true, true); err == nil {
		u.ArrayOfStr = arrayOfStr
		u.Type = CreateCompletionRequestStopTypeArrayOfStr
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for CreateCompletionRequestStop", string(data))
}

func (u CreateCompletionRequestStop) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.ArrayOfStr != nil {
		return utils.MarshalJSON(u.ArrayOfStr, "", true)
	}

	return nil, errors.New("could not marshal union type CreateCompletionRequestStop: all fields are null")
}

type CreateCompletionRequest struct {
	// Generates `best_of` completions server-side and returns the "best" (the one with the highest log probability per token). Results cannot be streamed.
	//
	// When used with `n`, `best_of` controls the number of candidate completions and `n` specifies how many to return – `best_of` must be greater than `n`.
	//
	// **Note:** Because this parameter generates many completions, it can quickly consume your token quota. Use carefully and ensure that you have reasonable settings for `max_tokens` and `stop`.
	//
	BestOf *int64 `default:"1" json:"best_of"`
	// Echo back the prompt in addition to the completion
	//
	Echo *bool `default:"false" json:"echo"`
	// Number between -2.0 and 2.0. Positive values penalize new tokens based on their existing frequency in the text so far, decreasing the model's likelihood to repeat the same line verbatim.
	//
	// [See more information about frequency and presence penalties.](/docs/api-reference/parameter-details)
	//
	FrequencyPenalty *float64 `default:"0" json:"frequency_penalty"`
	// Modify the likelihood of specified tokens appearing in the completion.
	//
	// Accepts a json object that maps tokens (specified by their token ID in the GPT tokenizer) to an associated bias value from -100 to 100. You can use this [tokenizer tool](/tokenizer?view=bpe) (which works for both GPT-2 and GPT-3) to convert text to token IDs. Mathematically, the bias is added to the logits generated by the model prior to sampling. The exact effect will vary per model, but values between -1 and 1 should decrease or increase likelihood of selection; values like -100 or 100 should result in a ban or exclusive selection of the relevant token.
	//
	// As an example, you can pass `{"50256": -100}` to prevent the <|endoftext|> token from being generated.
	//
	LogitBias map[string]int64 `json:"logit_bias,omitempty"`
	// Include the log probabilities on the `logprobs` most likely tokens, as well the chosen tokens. For example, if `logprobs` is 5, the API will return a list of the 5 most likely tokens. The API will always return the `logprob` of the sampled token, so there may be up to `logprobs+1` elements in the response.
	//
	// The maximum value for `logprobs` is 5.
	//
	Logprobs *int64 `default:"null" json:"logprobs"`
	// The maximum number of [tokens](/tokenizer) to generate in the completion.
	//
	// The token count of your prompt plus `max_tokens` cannot exceed the model's context length. [Example Python code](https://github.com/openai/openai-cookbook/blob/main/examples/How_to_count_tokens_with_tiktoken.ipynb) for counting tokens.
	//
	MaxTokens *int64 `default:"16" json:"max_tokens"`
	// ID of the model to use. You can use the [List models](/docs/api-reference/models/list) API to see all of your available models, or see our [Model overview](/docs/models/overview) for descriptions of them.
	//
	Model CreateCompletionRequestModel `json:"model"`
	// How many completions to generate for each prompt.
	//
	// **Note:** Because this parameter generates many completions, it can quickly consume your token quota. Use carefully and ensure that you have reasonable settings for `max_tokens` and `stop`.
	//
	N *int64 `default:"1" json:"n"`
	// Number between -2.0 and 2.0. Positive values penalize new tokens based on whether they appear in the text so far, increasing the model's likelihood to talk about new topics.
	//
	// [See more information about frequency and presence penalties.](/docs/api-reference/parameter-details)
	//
	PresencePenalty *float64 `default:"0" json:"presence_penalty"`
	// The prompt(s) to generate completions for, encoded as a string, array of strings, array of tokens, or array of token arrays.
	//
	// Note that <|endoftext|> is the document separator that the model sees during training, so if a prompt is not specified the model will generate as if from the beginning of a new document.
	//
	Prompt *Prompt `json:"prompt"`
	// Up to 4 sequences where the API will stop generating further tokens. The returned text will not contain the stop sequence.
	//
	Stop *CreateCompletionRequestStop `json:"stop,omitempty"`
	// Whether to stream back partial progress. If set, tokens will be sent as data-only [server-sent events](https://developer.mozilla.org/en-US/docs/Web/API/Server-sent_events/Using_server-sent_events#Event_stream_format) as they become available, with the stream terminated by a `data: [DONE]` message. [Example Python code](https://github.com/openai/openai-cookbook/blob/main/examples/How_to_stream_completions.ipynb).
	//
	Stream *bool `default:"false" json:"stream"`
	// The suffix that comes after a completion of inserted text.
	Suffix *string `default:"null" json:"suffix"`
	// What sampling temperature to use, between 0 and 2. Higher values like 0.8 will make the output more random, while lower values like 0.2 will make it more focused and deterministic.
	//
	// We generally recommend altering this or `top_p` but not both.
	//
	Temperature *float64 `default:"1" json:"temperature"`
	// An alternative to sampling with temperature, called nucleus sampling, where the model considers the results of the tokens with top_p probability mass. So 0.1 means only the tokens comprising the top 10% probability mass are considered.
	//
	// We generally recommend altering this or `temperature` but not both.
	//
	TopP *float64 `default:"1" json:"top_p"`
	// A unique identifier representing your end-user, which can help OpenAI to monitor and detect abuse. [Learn more](/docs/guides/safety-best-practices/end-user-ids).
	//
	User *string `json:"user,omitempty"`
}

func (c CreateCompletionRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateCompletionRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateCompletionRequest) GetBestOf() *int64 {
	if o == nil {
		return nil
	}
	return o.BestOf
}

func (o *CreateCompletionRequest) GetEcho() *bool {
	if o == nil {
		return nil
	}
	return o.Echo
}

func (o *CreateCompletionRequest) GetFrequencyPenalty() *float64 {
	if o == nil {
		return nil
	}
	return o.FrequencyPenalty
}

func (o *CreateCompletionRequest) GetLogitBias() map[string]int64 {
	if o == nil {
		return nil
	}
	return o.LogitBias
}

func (o *CreateCompletionRequest) GetLogprobs() *int64 {
	if o == nil {
		return nil
	}
	return o.Logprobs
}

func (o *CreateCompletionRequest) GetMaxTokens() *int64 {
	if o == nil {
		return nil
	}
	return o.MaxTokens
}

func (o *CreateCompletionRequest) GetModel() CreateCompletionRequestModel {
	if o == nil {
		return CreateCompletionRequestModel("")
	}
	return o.Model
}

func (o *CreateCompletionRequest) GetN() *int64 {
	if o == nil {
		return nil
	}
	return o.N
}

func (o *CreateCompletionRequest) GetPresencePenalty() *float64 {
	if o == nil {
		return nil
	}
	return o.PresencePenalty
}

func (o *CreateCompletionRequest) GetPrompt() *Prompt {
	if o == nil {
		return nil
	}
	return o.Prompt
}

func (o *CreateCompletionRequest) GetStop() *CreateCompletionRequestStop {
	if o == nil {
		return nil
	}
	return o.Stop
}

func (o *CreateCompletionRequest) GetStream() *bool {
	if o == nil {
		return nil
	}
	return o.Stream
}

func (o *CreateCompletionRequest) GetSuffix() *string {
	if o == nil {
		return nil
	}
	return o.Suffix
}

func (o *CreateCompletionRequest) GetTemperature() *float64 {
	if o == nil {
		return nil
	}
	return o.Temperature
}

func (o *CreateCompletionRequest) GetTopP() *float64 {
	if o == nil {
		return nil
	}
	return o.TopP
}

func (o *CreateCompletionRequest) GetUser() *string {
	if o == nil {
		return nil
	}
	return o.User
}