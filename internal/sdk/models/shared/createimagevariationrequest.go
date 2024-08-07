// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/antonbabenko/terraform-provider-openai/v2/internal/sdk/internal/utils"
)

type CreateImageVariationRequestImage struct {
	Content  []byte `multipartForm:"content"`
	FileName string `multipartForm:"name=image"`
}

func (o *CreateImageVariationRequestImage) GetContent() []byte {
	if o == nil {
		return []byte{}
	}
	return o.Content
}

func (o *CreateImageVariationRequestImage) GetFileName() string {
	if o == nil {
		return ""
	}
	return o.FileName
}

// CreateImageVariationRequestResponseFormat - The format in which the generated images are returned. Must be one of `url` or `b64_json`.
type CreateImageVariationRequestResponseFormat string

const (
	CreateImageVariationRequestResponseFormatURL     CreateImageVariationRequestResponseFormat = "url"
	CreateImageVariationRequestResponseFormatB64JSON CreateImageVariationRequestResponseFormat = "b64_json"
)

func (e CreateImageVariationRequestResponseFormat) ToPointer() *CreateImageVariationRequestResponseFormat {
	return &e
}
func (e *CreateImageVariationRequestResponseFormat) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "url":
		fallthrough
	case "b64_json":
		*e = CreateImageVariationRequestResponseFormat(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateImageVariationRequestResponseFormat: %v", v)
	}
}

// CreateImageVariationRequestSize - The size of the generated images. Must be one of `256x256`, `512x512`, or `1024x1024`.
type CreateImageVariationRequestSize string

const (
	CreateImageVariationRequestSizeTwoHundredAndFiftySixx256     CreateImageVariationRequestSize = "256x256"
	CreateImageVariationRequestSizeFiveHundredAndTwelvex512      CreateImageVariationRequestSize = "512x512"
	CreateImageVariationRequestSizeOneThousandAndTwentyFourx1024 CreateImageVariationRequestSize = "1024x1024"
)

func (e CreateImageVariationRequestSize) ToPointer() *CreateImageVariationRequestSize {
	return &e
}
func (e *CreateImageVariationRequestSize) UnmarshalJSON(data []byte) error {
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
		*e = CreateImageVariationRequestSize(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateImageVariationRequestSize: %v", v)
	}
}

type CreateImageVariationRequest struct {
	// The image to use as the basis for the variation(s). Must be a valid PNG file, less than 4MB, and square.
	Image CreateImageVariationRequestImage `multipartForm:"file"`
	// The number of images to generate. Must be between 1 and 10.
	N *int64 `default:"1" multipartForm:"name=n"`
	// The format in which the generated images are returned. Must be one of `url` or `b64_json`.
	ResponseFormat *CreateImageVariationRequestResponseFormat `default:"url" multipartForm:"name=response_format"`
	// The size of the generated images. Must be one of `256x256`, `512x512`, or `1024x1024`.
	Size *CreateImageVariationRequestSize `default:"1024x1024" multipartForm:"name=size"`
	// A unique identifier representing your end-user, which can help OpenAI to monitor and detect abuse. [Learn more](/docs/guides/safety-best-practices/end-user-ids).
	//
	User *string `multipartForm:"name=user"`
}

func (c CreateImageVariationRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateImageVariationRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateImageVariationRequest) GetImage() CreateImageVariationRequestImage {
	if o == nil {
		return CreateImageVariationRequestImage{}
	}
	return o.Image
}

func (o *CreateImageVariationRequest) GetN() *int64 {
	if o == nil {
		return nil
	}
	return o.N
}

func (o *CreateImageVariationRequest) GetResponseFormat() *CreateImageVariationRequestResponseFormat {
	if o == nil {
		return nil
	}
	return o.ResponseFormat
}

func (o *CreateImageVariationRequest) GetSize() *CreateImageVariationRequestSize {
	if o == nil {
		return nil
	}
	return o.Size
}

func (o *CreateImageVariationRequest) GetUser() *string {
	if o == nil {
		return nil
	}
	return o.User
}
