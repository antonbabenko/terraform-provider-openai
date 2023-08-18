resource "openai_embedding" "my_embedding" {
    input = {
        str = "This is a test."
    }
            model = "text-embedding-ada-002"
            user = "user-1234"
        }