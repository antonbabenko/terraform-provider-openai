// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type ChatCompletionRequestMessage struct {
	Content      types.String                               `tfsdk:"content"`
	FunctionCall *ChatCompletionResponseMessageFunctionCall `tfsdk:"function_call"`
	Name         types.String                               `tfsdk:"name"`
	Role         types.String                               `tfsdk:"role"`
}