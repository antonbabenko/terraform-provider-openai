// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/antonbabenko/terraform-provider-openai/v2/internal/sdk/internal/utils"
)

// CreateEditRequestModel - ID of the model to use. You can use the `text-davinci-edit-001` or `code-davinci-edit-001` model with this endpoint.
type CreateEditRequestModel string

const (
	CreateEditRequestModelTextDavinciEdit001 CreateEditRequestModel = "text-davinci-edit-001"
	CreateEditRequestModelCodeDavinciEdit001 CreateEditRequestModel = "code-davinci-edit-001"
)

func (e CreateEditRequestModel) ToPointer() *CreateEditRequestModel {
	return &e
}
func (e *CreateEditRequestModel) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "text-davinci-edit-001":
		fallthrough
	case "code-davinci-edit-001":
		*e = CreateEditRequestModel(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateEditRequestModel: %v", v)
	}
}

type CreateEditRequest struct {
	// The input text to use as a starting point for the edit.
	Input *string `default:"" json:"input"`
	// The instruction that tells the model how to edit the prompt.
	Instruction string `json:"instruction"`
	// ID of the model to use. You can use the `text-davinci-edit-001` or `code-davinci-edit-001` model with this endpoint.
	Model CreateEditRequestModel `json:"model"`
	// How many edits to generate for the input and instruction.
	N *int64 `default:"1" json:"n"`
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
}

func (c CreateEditRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateEditRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateEditRequest) GetInput() *string {
	if o == nil {
		return nil
	}
	return o.Input
}

func (o *CreateEditRequest) GetInstruction() string {
	if o == nil {
		return ""
	}
	return o.Instruction
}

func (o *CreateEditRequest) GetModel() CreateEditRequestModel {
	if o == nil {
		return CreateEditRequestModel("")
	}
	return o.Model
}

func (o *CreateEditRequest) GetN() *int64 {
	if o == nil {
		return nil
	}
	return o.N
}

func (o *CreateEditRequest) GetTemperature() *float64 {
	if o == nil {
		return nil
	}
	return o.Temperature
}

func (o *CreateEditRequest) GetTopP() *float64 {
	if o == nil {
		return nil
	}
	return o.TopP
}