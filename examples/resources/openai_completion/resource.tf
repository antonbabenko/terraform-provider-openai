resource "openai_completion" "my_completion" {
  best_of           = 4
  echo              = false
  frequency_penalty = 26.31
  logit_bias = {
    "invoice" = 1
    "Future"  = 7
  }
  logprobs         = 3
  max_tokens       = 16
  model            = "text-davinci-003"
  n                = 1
  presence_penalty = 69.43
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