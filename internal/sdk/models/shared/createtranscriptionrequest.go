// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/antonbabenko/terraform-provider-openai/v2/internal/sdk/internal/utils"
)

type CreateTranscriptionRequestFile struct {
	Content  []byte `multipartForm:"content"`
	FileName string `multipartForm:"name=file"`
}

func (o *CreateTranscriptionRequestFile) GetContent() []byte {
	if o == nil {
		return []byte{}
	}
	return o.Content
}

func (o *CreateTranscriptionRequestFile) GetFileName() string {
	if o == nil {
		return ""
	}
	return o.FileName
}

// CreateTranscriptionRequestModel - ID of the model to use. Only `whisper-1` is currently available.
type CreateTranscriptionRequestModel string

const (
	CreateTranscriptionRequestModelWhisper1 CreateTranscriptionRequestModel = "whisper-1"
)

func (e CreateTranscriptionRequestModel) ToPointer() *CreateTranscriptionRequestModel {
	return &e
}
func (e *CreateTranscriptionRequestModel) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "whisper-1":
		*e = CreateTranscriptionRequestModel(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateTranscriptionRequestModel: %v", v)
	}
}

// CreateTranscriptionRequestResponseFormat - The format of the transcript output, in one of these options: json, text, srt, verbose_json, or vtt.
type CreateTranscriptionRequestResponseFormat string

const (
	CreateTranscriptionRequestResponseFormatJSON        CreateTranscriptionRequestResponseFormat = "json"
	CreateTranscriptionRequestResponseFormatText        CreateTranscriptionRequestResponseFormat = "text"
	CreateTranscriptionRequestResponseFormatSrt         CreateTranscriptionRequestResponseFormat = "srt"
	CreateTranscriptionRequestResponseFormatVerboseJSON CreateTranscriptionRequestResponseFormat = "verbose_json"
	CreateTranscriptionRequestResponseFormatVtt         CreateTranscriptionRequestResponseFormat = "vtt"
)

func (e CreateTranscriptionRequestResponseFormat) ToPointer() *CreateTranscriptionRequestResponseFormat {
	return &e
}
func (e *CreateTranscriptionRequestResponseFormat) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "json":
		fallthrough
	case "text":
		fallthrough
	case "srt":
		fallthrough
	case "verbose_json":
		fallthrough
	case "vtt":
		*e = CreateTranscriptionRequestResponseFormat(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateTranscriptionRequestResponseFormat: %v", v)
	}
}

type CreateTranscriptionRequest struct {
	// The audio file object (not file name) to transcribe, in one of these formats: mp3, mp4, mpeg, mpga, m4a, wav, or webm.
	//
	File CreateTranscriptionRequestFile `multipartForm:"file"`
	// The language of the input audio. Supplying the input language in [ISO-639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes) format will improve accuracy and latency.
	//
	Language *string `multipartForm:"name=language"`
	// ID of the model to use. Only `whisper-1` is currently available.
	//
	Model CreateTranscriptionRequestModel `multipartForm:"name=model"`
	// An optional text to guide the model's style or continue a previous audio segment. The [prompt](/docs/guides/speech-to-text/prompting) should match the audio language.
	//
	Prompt *string `multipartForm:"name=prompt"`
	// The format of the transcript output, in one of these options: json, text, srt, verbose_json, or vtt.
	//
	ResponseFormat *CreateTranscriptionRequestResponseFormat `default:"json" multipartForm:"name=response_format"`
	// The sampling temperature, between 0 and 1. Higher values like 0.8 will make the output more random, while lower values like 0.2 will make it more focused and deterministic. If set to 0, the model will use [log probability](https://en.wikipedia.org/wiki/Log_probability) to automatically increase the temperature until certain thresholds are hit.
	//
	Temperature *float64 `default:"0" multipartForm:"name=temperature"`
}

func (c CreateTranscriptionRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateTranscriptionRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateTranscriptionRequest) GetFile() CreateTranscriptionRequestFile {
	if o == nil {
		return CreateTranscriptionRequestFile{}
	}
	return o.File
}

func (o *CreateTranscriptionRequest) GetLanguage() *string {
	if o == nil {
		return nil
	}
	return o.Language
}

func (o *CreateTranscriptionRequest) GetModel() CreateTranscriptionRequestModel {
	if o == nil {
		return CreateTranscriptionRequestModel("")
	}
	return o.Model
}

func (o *CreateTranscriptionRequest) GetPrompt() *string {
	if o == nil {
		return nil
	}
	return o.Prompt
}

func (o *CreateTranscriptionRequest) GetResponseFormat() *CreateTranscriptionRequestResponseFormat {
	if o == nil {
		return nil
	}
	return o.ResponseFormat
}

func (o *CreateTranscriptionRequest) GetTemperature() *float64 {
	if o == nil {
		return nil
	}
	return o.Temperature
}