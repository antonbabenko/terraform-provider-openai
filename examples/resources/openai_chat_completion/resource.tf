resource "openai_chat_completion" "my_chatcompletion" {
  frequency_penalty = 5
  function_call = {
    create_chat_completion_request_function_call_1 = "none"
  }
  functions = [
    {
      description = "...my_description..."
      name        = "Carlos Halvorson"
      parameters = {
        "non"  = "{ \"see\": \"documentation\" }"
        "quas" = "{ \"see\": \"documentation\" }"
      }
    },
  ]
  logit_bias = {
    "esse"       = 2
    "blanditiis" = 3
  }
  max_tokens = 2
  messages = [
    {
      content = "...my_content..."
      function_call = {
        arguments = "...my_arguments..."
        name      = "Marcos Smith"
      }
      name = "Phil Robel"
      role = "user"
    },
  ]
  model            = "gpt-3.5-turbo"
  n                = 1
  presence_penalty = 48.94
  stop = {
    str = "..."
  }
  stream      = true
  temperature = 1
  top_p       = 1
  user        = "user-1234"
}