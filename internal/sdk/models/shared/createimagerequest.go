// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/antonbabenko/terraform-provider-openai/v2/internal/sdk/internal/utils"
)

// CreateImageRequestResponseFormat - The format in which the generated images are returned. Must be one of `url` or `b64_json`.
type CreateImageRequestResponseFormat string

const (
	CreateImageRequestResponseFormatURL     CreateImageRequestResponseFormat = "url"
	CreateImageRequestResponseFormatB64JSON CreateImageRequestResponseFormat = "b64_json"
)

func (e CreateImageRequestResponseFormat) ToPointer() *CreateImageRequestResponseFormat {
	return &e
}
func (e *CreateImageRequestResponseFormat) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "url":
		fallthrough
	case "b64_json":
		*e = CreateImageRequestResponseFormat(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateImageRequestResponseFormat: %v", v)
	}
}

// CreateImageRequestSize - The size of the generated images. Must be one of `256x256`, `512x512`, or `1024x1024`.
type CreateImageRequestSize string

const (
	CreateImageRequestSizeTwoHundredAndFiftySixx256     CreateImageRequestSize = "256x256"
	CreateImageRequestSizeFiveHundredAndTwelvex512      CreateImageRequestSize = "512x512"
	CreateImageRequestSizeOneThousandAndTwentyFourx1024 CreateImageRequestSize = "1024x1024"
)

func (e CreateImageRequestSize) ToPointer() *CreateImageRequestSize {
	return &e
}
func (e *CreateImageRequestSize) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "256x256":
		fallthrough
	case "512x512":
		fallthrough
	case "1024x1024":
		*e = CreateImageRequestSize(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateImageRequestSize: %v", v)
	}
}

type CreateImageRequest struct {
	// The number of images to generate. Must be between 1 and 10.
	N *int64 `default:"1" json:"n"`
	// A text description of the desired image(s). The maximum length is 1000 characters.
	Prompt string `json:"prompt"`
	// The format in which the generated images are returned. Must be one of `url` or `b64_json`.
	ResponseFormat *CreateImageRequestResponseFormat `default:"url" json:"response_format"`
	// The size of the generated images. Must be one of `256x256`, `512x512`, or `1024x1024`.
	Size *CreateImageRequestSize `default:"1024x1024" json:"size"`
	// A unique identifier representing your end-user, which can help OpenAI to monitor and detect abuse. [Learn more](/docs/guides/safety-best-practices/end-user-ids).
	//
	User *string `json:"user,omitempty"`
}

func (c CreateImageRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateImageRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateImageRequest) GetN() *int64 {
	if o == nil {
		return nil
	}
	return o.N
}

func (o *CreateImageRequest) GetPrompt() string {
	if o == nil {
		return ""
	}
	return o.Prompt
}

func (o *CreateImageRequest) GetResponseFormat() *CreateImageRequestResponseFormat {
	if o == nil {
		return nil
	}
	return o.ResponseFormat
}

func (o *CreateImageRequest) GetSize() *CreateImageRequestSize {
	if o == nil {
		return nil
	}
	return o.Size
}

func (o *CreateImageRequest) GetUser() *string {
	if o == nil {
		return nil
	}
	return o.User
}
