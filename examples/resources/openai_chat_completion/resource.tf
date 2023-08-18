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