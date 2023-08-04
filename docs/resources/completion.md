---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "openai_completion Resource - terraform-provider-openai"
subcategory: ""
description: |-
  Completion Resource
---

# openai_completion (Resource)

Completion Resource

## Example Usage

```terraform
resource "openai_completion" "my_completion" {
  best_of           = 0
  echo              = true
  frequency_penalty = 2.02
  logit_bias = {
    "ipsam"    = 9
    "sapiente" = 8
  }
  logprobs         = 1
  max_tokens       = 16
  model            = "text-ada-001"
  n                = 1
  presence_penalty = 47.36
  prompt = {
    str = "This is a test."
  }
  stop = {
    str = "\n"
  }
  stream      = false
  suffix      = "test."
  temperature = 1
  top_p       = 1
  user        = "user-1234"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `model` (String) must be one of [text-davinci-003, text-davinci-002, text-davinci-001, code-davinci-002, text-curie-001, text-babbage-001, text-ada-001]
ID of the model to use. You can use the [List models](/docs/api-reference/models/list) API to see all of your available models, or see our [Model overview](/docs/models/overview) for descriptions of them.
- `prompt` (Attributes) The prompt(s) to generate completions for, encoded as a string, array of strings, array of tokens, or array of token arrays.

Note that <|endoftext|> is the document separator that the model sees during training, so if a prompt is not specified the model will generate as if from the beginning of a new document. (see [below for nested schema](#nestedatt--prompt))

### Optional

- `best_of` (Number) Generates `best_of` completions server-side and returns the "best" (the one with the highest log probability per token). Results cannot be streamed.

When used with `n`, `best_of` controls the number of candidate completions and `n` specifies how many to return – `best_of` must be greater than `n`.

**Note:** Because this parameter generates many completions, it can quickly consume your token quota. Use carefully and ensure that you have reasonable settings for `max_tokens` and `stop`.
- `echo` (Boolean) Echo back the prompt in addition to the completion
- `frequency_penalty` (Number) Number between -2.0 and 2.0. Positive values penalize new tokens based on their existing frequency in the text so far, decreasing the model's likelihood to repeat the same line verbatim.

[See more information about frequency and presence penalties.](/docs/api-reference/parameter-details)
- `logit_bias` (Map of Number) Modify the likelihood of specified tokens appearing in the completion.

Accepts a json object that maps tokens (specified by their token ID in the GPT tokenizer) to an associated bias value from -100 to 100. You can use this [tokenizer tool](/tokenizer?view=bpe) (which works for both GPT-2 and GPT-3) to convert text to token IDs. Mathematically, the bias is added to the logits generated by the model prior to sampling. The exact effect will vary per model, but values between -1 and 1 should decrease or increase likelihood of selection; values like -100 or 100 should result in a ban or exclusive selection of the relevant token.

As an example, you can pass `{"50256": -100}` to prevent the <|endoftext|> token from being generated.
- `logprobs` (Number) Include the log probabilities on the `logprobs` most likely tokens, as well the chosen tokens. For example, if `logprobs` is 5, the API will return a list of the 5 most likely tokens. The API will always return the `logprob` of the sampled token, so there may be up to `logprobs+1` elements in the response.

The maximum value for `logprobs` is 5.
- `max_tokens` (Number) The maximum number of [tokens](/tokenizer) to generate in the completion.

The token count of your prompt plus `max_tokens` cannot exceed the model's context length. [Example Python code](https://github.com/openai/openai-cookbook/blob/main/examples/How_to_count_tokens_with_tiktoken.ipynb) for counting tokens.
- `n` (Number) How many completions to generate for each prompt.

**Note:** Because this parameter generates many completions, it can quickly consume your token quota. Use carefully and ensure that you have reasonable settings for `max_tokens` and `stop`.
- `presence_penalty` (Number) Number between -2.0 and 2.0. Positive values penalize new tokens based on whether they appear in the text so far, increasing the model's likelihood to talk about new topics.

[See more information about frequency and presence penalties.](/docs/api-reference/parameter-details)
- `stop` (Attributes) Up to 4 sequences where the API will stop generating further tokens. The returned text will not contain the stop sequence. (see [below for nested schema](#nestedatt--stop))
- `stream` (Boolean) Whether to stream back partial progress. If set, tokens will be sent as data-only [server-sent events](https://developer.mozilla.org/en-US/docs/Web/API/Server-sent_events/Using_server-sent_events#Event_stream_format) as they become available, with the stream terminated by a `data: [DONE]` message. [Example Python code](https://github.com/openai/openai-cookbook/blob/main/examples/How_to_stream_completions.ipynb).
- `suffix` (String) The suffix that comes after a completion of inserted text.
- `temperature` (Number) What sampling temperature to use, between 0 and 2. Higher values like 0.8 will make the output more random, while lower values like 0.2 will make it more focused and deterministic.

We generally recommend altering this or `top_p` but not both.
- `top_p` (Number) An alternative to sampling with temperature, called nucleus sampling, where the model considers the results of the tokens with top_p probability mass. So 0.1 means only the tokens comprising the top 10% probability mass are considered.

We generally recommend altering this or `temperature` but not both.
- `user` (String) A unique identifier representing your end-user, which can help OpenAI to monitor and detect abuse. [Learn more](/docs/guides/safety-best-practices/end-user-ids).

### Read-Only

- `choices` (Attributes List) (see [below for nested schema](#nestedatt--choices))
- `created` (Number)
- `id` (String) The ID of this resource.
- `object` (String)
- `usage` (Attributes) (see [below for nested schema](#nestedatt--usage))

<a id="nestedatt--prompt"></a>
### Nested Schema for `prompt`

Optional:

- `array_ofarray_ofinteger` (List of List of Number)
- `array_ofinteger` (List of Number)
- `array_ofstr` (List of String)
- `str` (String)


<a id="nestedatt--stop"></a>
### Nested Schema for `stop`

Optional:

- `array_ofstr` (List of String)
- `str` (String)


<a id="nestedatt--choices"></a>
### Nested Schema for `choices`

Read-Only:

- `finish_reason` (String) must be one of [stop, length]
- `index` (Number)
- `logprobs` (Attributes) (see [below for nested schema](#nestedatt--choices--logprobs))
- `text` (String)

<a id="nestedatt--choices--logprobs"></a>
### Nested Schema for `choices.logprobs`

Read-Only:

- `text_offset` (List of Number)
- `token_logprobs` (List of Number)
- `tokens` (List of String)
- `top_logprobs` (String) Parsed as JSON.



<a id="nestedatt--usage"></a>
### Nested Schema for `usage`

Read-Only:

- `completion_tokens` (Number)
- `prompt_tokens` (Number)
- `total_tokens` (Number)

