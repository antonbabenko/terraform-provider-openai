// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	speakeasy_stringplanmodifier "github.com/antonbabenko/terraform-provider-openai/v2/internal/planmodifiers/stringplanmodifier"
	tfTypes "github.com/antonbabenko/terraform-provider-openai/v2/internal/provider/types"
	"github.com/antonbabenko/terraform-provider-openai/v2/internal/sdk"
	"github.com/antonbabenko/terraform-provider-openai/v2/internal/validators"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/numberdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/numberplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"math/big"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &CompletionResource{}
var _ resource.ResourceWithImportState = &CompletionResource{}

func NewCompletionResource() resource.Resource {
	return &CompletionResource{}
}

// CompletionResource defines the resource implementation.
type CompletionResource struct {
	client *sdk.Oai
}

// CompletionResourceModel describes the resource data model.
type CompletionResourceModel struct {
	BestOf           types.Int64                               `tfsdk:"best_of"`
	Choices          []tfTypes.CreateCompletionResponseChoices `tfsdk:"choices"`
	Created          types.Int64                               `tfsdk:"created"`
	Echo             types.Bool                                `tfsdk:"echo"`
	FrequencyPenalty types.Number                              `tfsdk:"frequency_penalty"`
	ID               types.String                              `tfsdk:"id"`
	LogitBias        map[string]types.Int64                    `tfsdk:"logit_bias"`
	Logprobs         types.Int64                               `tfsdk:"logprobs"`
	MaxTokens        types.Int64                               `tfsdk:"max_tokens"`
	Model            types.String                              `tfsdk:"model"`
	N                types.Int64                               `tfsdk:"n"`
	Object           types.String                              `tfsdk:"object"`
	PresencePenalty  types.Number                              `tfsdk:"presence_penalty"`
	Prompt           *tfTypes.Prompt                           `tfsdk:"prompt"`
	Stop             *tfTypes.Stop                             `tfsdk:"stop"`
	Stream           types.Bool                                `tfsdk:"stream"`
	Suffix           types.String                              `tfsdk:"suffix"`
	Temperature      types.Number                              `tfsdk:"temperature"`
	TopP             types.Number                              `tfsdk:"top_p"`
	Usage            *tfTypes.Usage                            `tfsdk:"usage"`
	User             types.String                              `tfsdk:"user"`
}

func (r *CompletionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_completion"
}

func (r *CompletionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Completion Resource",
		Attributes: map[string]schema.Attribute{
			"best_of": schema.Int64Attribute{
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplaceIfConfigured(),
				},
				Optional: true,
				Default:  int64default.StaticInt64(1),
				MarkdownDescription: `Generates ` + "`" + `best_of` + "`" + ` completions server-side and returns the "best" (the one with the highest log probability per token). Results cannot be streamed.` + "\n" +
					`` + "\n" +
					`When used with ` + "`" + `n` + "`" + `, ` + "`" + `best_of` + "`" + ` controls the number of candidate completions and ` + "`" + `n` + "`" + ` specifies how many to return – ` + "`" + `best_of` + "`" + ` must be greater than ` + "`" + `n` + "`" + `.` + "\n" +
					`` + "\n" +
					`**Note:** Because this parameter generates many completions, it can quickly consume your token quota. Use carefully and ensure that you have reasonable settings for ` + "`" + `max_tokens` + "`" + ` and ` + "`" + `stop` + "`" + `.` + "\n" +
					`` + "\n" +
					`Requires replacement if changed. ; Default: 1`,
				Validators: []validator.Int64{
					int64validator.AtMost(20),
				},
			},
			"choices": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"finish_reason": schema.StringAttribute{
							Computed:    true,
							Description: `must be one of ["stop", "length"]`,
							Validators: []validator.String{
								stringvalidator.OneOf(
									"stop",
									"length",
								),
							},
						},
						"index": schema.Int64Attribute{
							Computed: true,
						},
						"logprobs": schema.SingleNestedAttribute{
							Computed: true,
							Attributes: map[string]schema.Attribute{
								"text_offset": schema.ListAttribute{
									Computed:    true,
									ElementType: types.Int64Type,
								},
								"token_logprobs": schema.ListAttribute{
									Computed:    true,
									ElementType: types.NumberType,
								},
								"tokens": schema.ListAttribute{
									Computed:    true,
									ElementType: types.StringType,
								},
								"top_logprobs": schema.StringAttribute{
									Computed:    true,
									Description: `Parsed as JSON.`,
									Validators: []validator.String{
										validators.IsValidJSON(),
									},
								},
							},
						},
						"text": schema.StringAttribute{
							Computed: true,
						},
					},
				},
			},
			"created": schema.Int64Attribute{
				Computed: true,
			},
			"echo": schema.BoolAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplaceIfConfigured(),
				},
				Optional: true,
				Default:  booldefault.StaticBool(false),
				MarkdownDescription: `Echo back the prompt in addition to the completion` + "\n" +
					`` + "\n" +
					`Requires replacement if changed. ; Default: false`,
			},
			"frequency_penalty": schema.NumberAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.Number{
					numberplanmodifier.RequiresReplaceIfConfigured(),
				},
				Optional: true,
				Default:  numberdefault.StaticBigFloat(big.NewFloat(0)),
				MarkdownDescription: `Number between -2.0 and 2.0. Positive values penalize new tokens based on their existing frequency in the text so far, decreasing the model's likelihood to repeat the same line verbatim.` + "\n" +
					`` + "\n" +
					`[See more information about frequency and presence penalties.](/docs/api-reference/parameter-details)` + "\n" +
					`` + "\n" +
					`Requires replacement if changed. ; Default: 0`,
			},
			"id": schema.StringAttribute{
				Computed: true,
			},
			"logit_bias": schema.MapAttribute{
				PlanModifiers: []planmodifier.Map{
					mapplanmodifier.RequiresReplaceIfConfigured(),
				},
				Optional:    true,
				ElementType: types.Int64Type,
				MarkdownDescription: `Modify the likelihood of specified tokens appearing in the completion.` + "\n" +
					`` + "\n" +
					`Accepts a json object that maps tokens (specified by their token ID in the GPT tokenizer) to an associated bias value from -100 to 100. You can use this [tokenizer tool](/tokenizer?view=bpe) (which works for both GPT-2 and GPT-3) to convert text to token IDs. Mathematically, the bias is added to the logits generated by the model prior to sampling. The exact effect will vary per model, but values between -1 and 1 should decrease or increase likelihood of selection; values like -100 or 100 should result in a ban or exclusive selection of the relevant token.` + "\n" +
					`` + "\n" +
					`As an example, you can pass ` + "`" + `{"50256": -100}` + "`" + ` to prevent the <|endoftext|> token from being generated.` + "\n" +
					`` + "\n" +
					`Requires replacement if changed. `,
			},
			"logprobs": schema.Int64Attribute{
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplaceIfConfigured(),
				},
				Optional: true,
				MarkdownDescription: `Include the log probabilities on the ` + "`" + `logprobs` + "`" + ` most likely tokens, as well the chosen tokens. For example, if ` + "`" + `logprobs` + "`" + ` is 5, the API will return a list of the 5 most likely tokens. The API will always return the ` + "`" + `logprob` + "`" + ` of the sampled token, so there may be up to ` + "`" + `logprobs+1` + "`" + ` elements in the response.` + "\n" +
					`` + "\n" +
					`The maximum value for ` + "`" + `logprobs` + "`" + ` is 5.` + "\n" +
					`` + "\n" +
					`Requires replacement if changed. ; Default: null`,
				Validators: []validator.Int64{
					int64validator.AtMost(5),
				},
			},
			"max_tokens": schema.Int64Attribute{
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplaceIfConfigured(),
				},
				Optional: true,
				Default:  int64default.StaticInt64(16),
				MarkdownDescription: `The maximum number of [tokens](/tokenizer) to generate in the completion.` + "\n" +
					`` + "\n" +
					`The token count of your prompt plus ` + "`" + `max_tokens` + "`" + ` cannot exceed the model's context length. [Example Python code](https://github.com/openai/openai-cookbook/blob/main/examples/How_to_count_tokens_with_tiktoken.ipynb) for counting tokens.` + "\n" +
					`` + "\n" +
					`Requires replacement if changed. ; Default: 16`,
			},
			"model": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
					speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
				},
				Required: true,
				MarkdownDescription: `ID of the model to use. You can use the [List models](/docs/api-reference/models/list) API to see all of your available models, or see our [Model overview](/docs/models/overview) for descriptions of them.` + "\n" +
					`` + "\n" +
					`Requires replacement if changed. ; must be one of ["text-davinci-003", "text-davinci-002", "text-davinci-001", "code-davinci-002", "text-curie-001", "text-babbage-001", "text-ada-001"]`,
				Validators: []validator.String{
					stringvalidator.OneOf(
						"text-davinci-003",
						"text-davinci-002",
						"text-davinci-001",
						"code-davinci-002",
						"text-curie-001",
						"text-babbage-001",
						"text-ada-001",
					),
				},
			},
			"n": schema.Int64Attribute{
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplaceIfConfigured(),
				},
				Optional: true,
				Default:  int64default.StaticInt64(1),
				MarkdownDescription: `How many completions to generate for each prompt.` + "\n" +
					`` + "\n" +
					`**Note:** Because this parameter generates many completions, it can quickly consume your token quota. Use carefully and ensure that you have reasonable settings for ` + "`" + `max_tokens` + "`" + ` and ` + "`" + `stop` + "`" + `.` + "\n" +
					`` + "\n" +
					`Requires replacement if changed. ; Default: 1`,
				Validators: []validator.Int64{
					int64validator.Between(1, 128),
				},
			},
			"object": schema.StringAttribute{
				Computed: true,
			},
			"presence_penalty": schema.NumberAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.Number{
					numberplanmodifier.RequiresReplaceIfConfigured(),
				},
				Optional: true,
				Default:  numberdefault.StaticBigFloat(big.NewFloat(0)),
				MarkdownDescription: `Number between -2.0 and 2.0. Positive values penalize new tokens based on whether they appear in the text so far, increasing the model's likelihood to talk about new topics.` + "\n" +
					`` + "\n" +
					`[See more information about frequency and presence penalties.](/docs/api-reference/parameter-details)` + "\n" +
					`` + "\n" +
					`Requires replacement if changed. ; Default: 0`,
			},
			"prompt": schema.SingleNestedAttribute{
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.RequiresReplaceIfConfigured(),
				},
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"str": schema.StringAttribute{
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplaceIfConfigured(),
						},
						Optional:    true,
						Description: `Requires replacement if changed. `,
						Validators: []validator.String{
							stringvalidator.ConflictsWith(path.Expressions{
								path.MatchRelative().AtParent().AtName("array_of_str"),
								path.MatchRelative().AtParent().AtName("array_of_integer"),
								path.MatchRelative().AtParent().AtName("array_of_array_of_integer"),
							}...),
						},
					},
					"array_of_str": schema.ListAttribute{
						PlanModifiers: []planmodifier.List{
							listplanmodifier.RequiresReplaceIfConfigured(),
						},
						Optional:    true,
						ElementType: types.StringType,
						Description: `Requires replacement if changed. `,
						Validators: []validator.List{
							listvalidator.ConflictsWith(path.Expressions{
								path.MatchRelative().AtParent().AtName("str"),
								path.MatchRelative().AtParent().AtName("array_of_integer"),
								path.MatchRelative().AtParent().AtName("array_of_array_of_integer"),
							}...),
						},
					},
					"array_of_integer": schema.ListAttribute{
						PlanModifiers: []planmodifier.List{
							listplanmodifier.RequiresReplaceIfConfigured(),
						},
						Optional:    true,
						ElementType: types.Int64Type,
						Description: `Requires replacement if changed. `,
						Validators: []validator.List{
							listvalidator.ConflictsWith(path.Expressions{
								path.MatchRelative().AtParent().AtName("str"),
								path.MatchRelative().AtParent().AtName("array_of_str"),
								path.MatchRelative().AtParent().AtName("array_of_array_of_integer"),
							}...),
							listvalidator.SizeAtLeast(1),
						},
					},
					"array_of_array_of_integer": schema.ListAttribute{
						PlanModifiers: []planmodifier.List{
							listplanmodifier.RequiresReplaceIfConfigured(),
						},
						Optional: true,
						ElementType: types.ListType{
							ElemType: types.Int64Type,
						},
						Description: `Requires replacement if changed. `,
						Validators: []validator.List{
							listvalidator.ConflictsWith(path.Expressions{
								path.MatchRelative().AtParent().AtName("str"),
								path.MatchRelative().AtParent().AtName("array_of_str"),
								path.MatchRelative().AtParent().AtName("array_of_integer"),
							}...),
							listvalidator.SizeAtLeast(1),
						},
					},
				},
				MarkdownDescription: `The prompt(s) to generate completions for, encoded as a string, array of strings, array of tokens, or array of token arrays.` + "\n" +
					`` + "\n" +
					`Note that <|endoftext|> is the document separator that the model sees during training, so if a prompt is not specified the model will generate as if from the beginning of a new document.` + "\n" +
					`` + "\n" +
					`Requires replacement if changed. `,
				Validators: []validator.Object{
					validators.ExactlyOneChild(),
				},
			},
			"stop": schema.SingleNestedAttribute{
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.RequiresReplaceIfConfigured(),
				},
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"str": schema.StringAttribute{
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplaceIfConfigured(),
						},
						Optional:    true,
						Description: `Requires replacement if changed. `,
						Validators: []validator.String{
							stringvalidator.ConflictsWith(path.Expressions{
								path.MatchRelative().AtParent().AtName("array_of_str"),
							}...),
						},
					},
					"array_of_str": schema.ListAttribute{
						PlanModifiers: []planmodifier.List{
							listplanmodifier.RequiresReplaceIfConfigured(),
						},
						Optional:    true,
						ElementType: types.StringType,
						Description: `Requires replacement if changed. `,
						Validators: []validator.List{
							listvalidator.ConflictsWith(path.Expressions{
								path.MatchRelative().AtParent().AtName("str"),
							}...),
							listvalidator.SizeAtLeast(1),
							listvalidator.SizeAtMost(4),
						},
					},
				},
				MarkdownDescription: `Up to 4 sequences where the API will stop generating further tokens. The returned text will not contain the stop sequence.` + "\n" +
					`` + "\n" +
					`Requires replacement if changed. `,
				Validators: []validator.Object{
					validators.ExactlyOneChild(),
				},
			},
			"stream": schema.BoolAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplaceIfConfigured(),
				},
				Optional: true,
				Default:  booldefault.StaticBool(false),
				MarkdownDescription: `Whether to stream back partial progress. If set, tokens will be sent as data-only [server-sent events](https://developer.mozilla.org/en-US/docs/Web/API/Server-sent_events/Using_server-sent_events#Event_stream_format) as they become available, with the stream terminated by a ` + "`" + `data: [DONE]` + "`" + ` message. [Example Python code](https://github.com/openai/openai-cookbook/blob/main/examples/How_to_stream_completions.ipynb).` + "\n" +
					`` + "\n" +
					`Requires replacement if changed. ; Default: false`,
			},
			"suffix": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Optional:    true,
				Description: `The suffix that comes after a completion of inserted text. Requires replacement if changed. ; Default: null`,
			},
			"temperature": schema.NumberAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.Number{
					numberplanmodifier.RequiresReplaceIfConfigured(),
				},
				Optional: true,
				Default:  numberdefault.StaticBigFloat(big.NewFloat(1)),
				MarkdownDescription: `What sampling temperature to use, between 0 and 2. Higher values like 0.8 will make the output more random, while lower values like 0.2 will make it more focused and deterministic.` + "\n" +
					`` + "\n" +
					`We generally recommend altering this or ` + "`" + `top_p` + "`" + ` but not both.` + "\n" +
					`` + "\n" +
					`Requires replacement if changed. ; Default: 1`,
			},
			"top_p": schema.NumberAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.Number{
					numberplanmodifier.RequiresReplaceIfConfigured(),
				},
				Optional: true,
				Default:  numberdefault.StaticBigFloat(big.NewFloat(1)),
				MarkdownDescription: `An alternative to sampling with temperature, called nucleus sampling, where the model considers the results of the tokens with top_p probability mass. So 0.1 means only the tokens comprising the top 10% probability mass are considered.` + "\n" +
					`` + "\n" +
					`We generally recommend altering this or ` + "`" + `temperature` + "`" + ` but not both.` + "\n" +
					`` + "\n" +
					`Requires replacement if changed. ; Default: 1`,
			},
			"usage": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"completion_tokens": schema.Int64Attribute{
						Computed: true,
					},
					"prompt_tokens": schema.Int64Attribute{
						Computed: true,
					},
					"total_tokens": schema.Int64Attribute{
						Computed: true,
					},
				},
			},
			"user": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
				},
				Optional: true,
				MarkdownDescription: `A unique identifier representing your end-user, which can help OpenAI to monitor and detect abuse. [Learn more](/docs/guides/safety-best-practices/end-user-ids).` + "\n" +
					`` + "\n" +
					`Requires replacement if changed. `,
			},
		},
	}
}

func (r *CompletionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.Oai)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *sdk.Oai, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *CompletionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *CompletionResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(plan.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	request := *data.ToSharedCreateCompletionRequest()
	res, err := r.client.OpenAI.CreateCompletion(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if !(res.CreateCompletionResponse != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedCreateCompletionResponse(res.CreateCompletionResponse)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CompletionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *CompletionResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Not Implemented; we rely entirely on CREATE API request response

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CompletionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *CompletionResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	// Not Implemented; all attributes marked as RequiresReplace

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *CompletionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *CompletionResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Not Implemented; entity does not have a configured DELETE operation
}

func (r *CompletionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.AddError("Not Implemented", "No available import state operation is available for resource completion.")
}
