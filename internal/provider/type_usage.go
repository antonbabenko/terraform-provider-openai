// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type Usage struct {
	CompletionTokens types.Int64 `tfsdk:"completion_tokens"`
	PromptTokens     types.Int64 `tfsdk:"prompt_tokens"`
	TotalTokens      types.Int64 `tfsdk:"total_tokens"`
}
