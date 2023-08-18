resource "openai_image" "my_image" {
    n = 1
            prompt = "A cute baby sea otter"
            response_format = "url"
            size = "1024x1024"
            user = "user-1234"
        }