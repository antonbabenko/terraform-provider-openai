resource "openai_completion" "my_completion" {
  best_of           = 4
  echo              = false
  frequency_penalty = 48.05
  logit_bias = {
    "totam" = 2
    "fuga"  = 2
  }
  logprobs         = 8
  max_tokens       = 16
  model            = "text-davinci-001"
  n                = 1
  presence_penalty = 26.31
  prompt = {
    str = "This is a test."
  }
  stop = {
    str = "\n"
  }
  stream      = true
  suffix      = "test."
  temperature = 1
  top_p       = 1
  user        = "user-1234"
}