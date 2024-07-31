resource "openai_embedding" "my_embedding" {
  input = {
    str = "The quick brown fox jumped over the lazy dog"
  }
  model = "text-embedding-ada-002"
  user  = "user-1234"
}