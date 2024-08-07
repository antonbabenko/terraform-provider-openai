resource "openai_chat_completion" "my_chatcompletion" {
  frequency_penalty = 5
  function_call = {
    one = "none"
  }
  functions = [
    {
      description = "...my_description..."
      name        = "Carlos Halvorson"
      parameters = {
        "Granite" = "{ \"see\": \"documentation\" }"
        "DNS"     = "{ \"see\": \"documentation\" }"
      }
    },
  ]
  logit_bias = {
    "copying"      = 2
    "transmitting" = 9
  }
  max_tokens = 9
  messages = [
    {
      content = "...my_content..."
      function_call = {
        arguments = "...my_arguments..."
        name      = "Mae VonRueden"
      }
      name = "Rick Gottlieb"
      role = "user"
    },
  ]
  model            = "gpt-3.5-turbo"
  n                = 1
  presence_penalty = 48.05
  stop = {
    str = "..."
  }
  stream      = false
  temperature = 1
  top_p       = 1
  user        = "user-1234"
}