// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type Choices struct {
	FinishReason types.String                  `tfsdk:"finish_reason"`
	Index        types.Int64                   `tfsdk:"index"`
	Message      ChatCompletionResponseMessage `tfsdk:"message"`
}
