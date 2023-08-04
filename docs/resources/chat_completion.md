---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "openai_chat_completion Resource - terraform-provider-openai"
subcategory: ""
description: |-
  ChatCompletion Resource
---

# openai_chat_completion (Resource)

ChatCompletion Resource

## Example Usage

```terraform
resource "openai_chat_completion" "my_chatcompletion" {
  frequency_penalty = 54.88
  function_call = {
    create_chat_completion_request_function_call_1 = "auto"
  }
  functions = [
    {
      description = "...my_description..."
      name        = "Ellis Mitchell"
      parameters = {
        "illum" = "{ \"see\": \"documentation\" }"
        "vel"   = "{ \"see\": \"documentation\" }"
      }
    },
  ]
  logit_bias = {
    "error"    = 7
    "suscipit" = 4
  }
  max_tokens = 3
  messages = [
    {
      content = "...my_content..."
      function_call = {
        arguments = "...my_arguments..."
        name      = "Larry Windler"
      }
      name = "Alexandra Schulist"
      role = "assistant"
    },
  ]
  model            = "gpt-3.5-turbo"
  n                = 1
  presence_penalty = 92.56
  stop = {
    str = "..."
  }
  stream      = false
  temperature = 1
  top_p       = 1
  user        = "user-1234"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `messages` (Attributes List) A list of messages comprising the conversation so far. [Example Python code](https://github.com/openai/openai-cookbook/blob/main/examples/How_to_format_inputs_to_ChatGPT_models.ipynb). (see [below for nested schema](#nestedatt--messages))
- `model` (String) must be one of [gpt-4, gpt-4-0314, gpt-4-0613, gpt-4-32k, gpt-4-32k-0314, gpt-4-32k-0613, gpt-3.5-turbo, gpt-3.5-turbo-16k, gpt-3.5-turbo-0301, gpt-3.5-turbo-0613, gpt-3.5-turbo-16k-0613]
ID of the model to use. See the [model endpoint compatibility](/docs/models/model-endpoint-compatibility) table for details on which models work with the Chat API.

### Optional

- `frequency_penalty` (Number) Number between -2.0 and 2.0. Positive values penalize new tokens based on their existing frequency in the text so far, decreasing the model's likelihood to repeat the same line verbatim.

[See more information about frequency and presence penalties.](/docs/api-reference/parameter-details)
- `function_call` (Attributes) Controls how the model responds to function calls. "none" means the model does not call a function, and responds to the end-user. "auto" means the model can pick between an end-user or calling a function.  Specifying a particular function via `{"name":\ "my_function"}` forces the model to call that function. "none" is the default when no functions are present. "auto" is the default if functions are present. (see [below for nested schema](#nestedatt--function_call))
- `functions` (Attributes List) A list of functions the model may generate JSON inputs for. (see [below for nested schema](#nestedatt--functions))
- `logit_bias` (Map of Number) Modify the likelihood of specified tokens appearing in the completion.

Accepts a json object that maps tokens (specified by their token ID in the tokenizer) to an associated bias value from -100 to 100. Mathematically, the bias is added to the logits generated by the model prior to sampling. The exact effect will vary per model, but values between -1 and 1 should decrease or increase likelihood of selection; values like -100 or 100 should result in a ban or exclusive selection of the relevant token.
- `max_tokens` (Number) The maximum number of [tokens](/tokenizer) to generate in the chat completion.

The total length of input tokens and generated tokens is limited by the model's context length. [Example Python code](https://github.com/openai/openai-cookbook/blob/main/examples/How_to_count_tokens_with_tiktoken.ipynb) for counting tokens.
- `n` (Number) How many chat completion choices to generate for each input message.
- `presence_penalty` (Number) Number between -2.0 and 2.0. Positive values penalize new tokens based on whether they appear in the text so far, increasing the model's likelihood to talk about new topics.

[See more information about frequency and presence penalties.](/docs/api-reference/parameter-details)
- `stop` (Attributes) Up to 4 sequences where the API will stop generating further tokens. (see [below for nested schema](#nestedatt--stop))
- `stream` (Boolean) If set, partial message deltas will be sent, like in ChatGPT. Tokens will be sent as data-only [server-sent events](https://developer.mozilla.org/en-US/docs/Web/API/Server-sent_events/Using_server-sent_events#Event_stream_format) as they become available, with the stream terminated by a `data: [DONE]` message. [Example Python code](https://github.com/openai/openai-cookbook/blob/main/examples/How_to_stream_completions.ipynb).
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

<a id="nestedatt--messages"></a>
### Nested Schema for `messages`

Required:

- `content` (String) The contents of the message. `content` is required for all messages, and may be null for assistant messages with function calls.
- `role` (String) must be one of [system, user, assistant, function]
The role of the messages author. One of `system`, `user`, `assistant`, or `function`.

Optional:

- `function_call` (Attributes) The name and arguments of a function that should be called, as generated by the model. (see [below for nested schema](#nestedatt--messages--function_call))
- `name` (String) The name of the author of this message. `name` is required if role is `function`, and it should be the name of the function whose response is in the `content`. May contain a-z, A-Z, 0-9, and underscores, with a maximum length of 64 characters.

<a id="nestedatt--messages--function_call"></a>
### Nested Schema for `messages.function_call`

Required:

- `arguments` (String) The arguments to call the function with, as generated by the model in JSON format. Note that the model does not always generate valid JSON, and may hallucinate parameters not defined by your function schema. Validate the arguments in your code before calling your function.
- `name` (String) The name of the function to call.



<a id="nestedatt--function_call"></a>
### Nested Schema for `function_call`

Optional:

- `create_chat_completion_request_function_call_1` (String) must be one of [none, auto]
Controls how the model responds to function calls. "none" means the model does not call a function, and responds to the end-user. "auto" means the model can pick between an end-user or calling a function.  Specifying a particular function via `{"name":\ "my_function"}` forces the model to call that function. "none" is the default when no functions are present. "auto" is the default if functions are present.
- `create_chat_completion_request_function_call_2` (Attributes) Controls how the model responds to function calls. "none" means the model does not call a function, and responds to the end-user. "auto" means the model can pick between an end-user or calling a function.  Specifying a particular function via `{"name":\ "my_function"}` forces the model to call that function. "none" is the default when no functions are present. "auto" is the default if functions are present. (see [below for nested schema](#nestedatt--function_call--create_chat_completion_request_function_call_2))

<a id="nestedatt--function_call--create_chat_completion_request_function_call_2"></a>
### Nested Schema for `function_call.create_chat_completion_request_function_call_2`

Required:

- `name` (String) The name of the function to call.



<a id="nestedatt--functions"></a>
### Nested Schema for `functions`

Required:

- `name` (String) The name of the function to be called. Must be a-z, A-Z, 0-9, or contain underscores and dashes, with a maximum length of 64.
- `parameters` (Map of String) The parameters the functions accepts, described as a JSON Schema object. See the [guide](/docs/guides/gpt/function-calling) for examples, and the [JSON Schema reference](https://json-schema.org/understanding-json-schema/) for documentation about the format.

To describe a function that accepts no parameters, provide the value `{"type": "object", "properties": {}}`.

Optional:

- `description` (String) A description of what the function does, used by the model to choose when and how to call the function.


<a id="nestedatt--stop"></a>
### Nested Schema for `stop`

Optional:

- `array_ofstr` (List of String)
- `str` (String)


<a id="nestedatt--choices"></a>
### Nested Schema for `choices`

Read-Only:

- `finish_reason` (String) must be one of [stop, length, function_call]
- `index` (Number)
- `message` (Attributes) (see [below for nested schema](#nestedatt--choices--message))

<a id="nestedatt--choices--message"></a>
### Nested Schema for `choices.message`

Read-Only:

- `content` (String) The contents of the message.
- `function_call` (Attributes) The name and arguments of a function that should be called, as generated by the model. (see [below for nested schema](#nestedatt--choices--message--function_call))
- `role` (String) must be one of [system, user, assistant, function]
The role of the author of this message.

<a id="nestedatt--choices--message--function_call"></a>
### Nested Schema for `choices.message.function_call`

Read-Only:

- `arguments` (String) The arguments to call the function with, as generated by the model in JSON format. Note that the model does not always generate valid JSON, and may hallucinate parameters not defined by your function schema. Validate the arguments in your code before calling your function.
- `name` (String) The name of the function to call.




<a id="nestedatt--usage"></a>
### Nested Schema for `usage`

Read-Only:

- `completion_tokens` (Number)
- `prompt_tokens` (Number)
- `total_tokens` (Number)

