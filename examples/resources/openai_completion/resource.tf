resource "openai_completion" "my_completion" {
  best_of           = 2
  echo              = true
  frequency_penalty = 26.2
  logit_bias = {
    "tesla"      = 4
    "Incredible" = 7
  }
  logprobs         = 2
  max_tokens       = 16
  model            = "text-davinci-001"
  n                = 1
  presence_penalty = 48.85
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