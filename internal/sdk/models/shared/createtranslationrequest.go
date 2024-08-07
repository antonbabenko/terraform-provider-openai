// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/antonbabenko/terraform-provider-openai/v2/internal/sdk/internal/utils"
)

type CreateTranslationRequestFile struct {
	Content  []byte `multipartForm:"content"`
	FileName string `multipartForm:"name=file"`
}

func (o *CreateTranslationRequestFile) GetContent() []byte {
	if o == nil {
		return []byte{}
	}
	return o.Content
}

func (o *CreateTranslationRequestFile) GetFileName() string {
	if o == nil {
		return ""
	}
	return o.FileName
}

// CreateTranslationRequestModel - ID of the model to use. Only `whisper-1` is currently available.
type CreateTranslationRequestModel string

const (
	CreateTranslationRequestModelWhisper1 CreateTranslationRequestModel = "whisper-1"
)

func (e CreateTranslationRequestModel) ToPointer() *CreateTranslationRequestModel {
	return &e
}
func (e *CreateTranslationRequestModel) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "whisper-1":
		*e = CreateTranslationRequestModel(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateTranslationRequestModel: %v", v)
	}
}

type CreateTranslationRequest struct {
	// The audio file object (not file name) translate, in one of these formats: mp3, mp4, mpeg, mpga, m4a, wav, or webm.
	//
	File CreateTranslationRequestFile `multipartForm:"file"`
	// ID of the model to use. Only `whisper-1` is currently available.
	//
	Model CreateTranslationRequestModel `multipartForm:"name=model"`
	// An optional text to guide the model's style or continue a previous audio segment. The [prompt](/docs/guides/speech-to-text/prompting) should be in English.
	//
	Prompt *string `multipartForm:"name=prompt"`
	// The format of the transcript output, in one of these options: json, text, srt, verbose_json, or vtt.
	//
	ResponseFormat *string `default:"json" multipartForm:"name=response_format"`
	// The sampling temperature, between 0 and 1. Higher values like 0.8 will make the output more random, while lower values like 0.2 will make it more focused and deterministic. If set to 0, the model will use [log probability](https://en.wikipedia.org/wiki/Log_probability) to automatically increase the temperature until certain thresholds are hit.
	//
	Temperature *float64 `default:"0" multipartForm:"name=temperature"`
}

func (c CreateTranslationRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateTranslationRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateTranslationRequest) GetFile() CreateTranslationRequestFile {
	if o == nil {
		return CreateTranslationRequestFile{}
	}
	return o.File
}

func (o *CreateTranslationRequest) GetModel() CreateTranslationRequestModel {
	if o == nil {
		return CreateTranslationRequestModel("")
	}
	return o.Model
}

func (o *CreateTranslationRequest) GetPrompt() *string {
	if o == nil {
		return nil
	}
	return o.Prompt
}

func (o *CreateTranslationRequest) GetResponseFormat() *string {
	if o == nil {
		return nil
	}
	return o.ResponseFormat
}

func (o *CreateTranslationRequest) GetTemperature() *float64 {
	if o == nil {
		return nil
	}
	return o.Temperature
}
