// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type ChatCompletionFunctions struct {
	Description types.String            `tfsdk:"description"`
	Name        types.String            `tfsdk:"name"`
	Parameters  map[string]types.String `tfsdk:"parameters"`
}
