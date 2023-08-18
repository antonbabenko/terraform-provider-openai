resource "openai_completion" "my_completion" {
    best_of = 0
            echo = true
            frequency_penalty = 2.02
            logit_bias = {
        "ipsam" = 9
        "sapiente" = 8
    }
            logprobs = 1
            max_tokens = 16
            model = "text-ada-001"
            n = 1
            presence_penalty = 47.36
            prompt = {
        str = "This is a test."
    }
            stop = {
        str = "\n"
    }
            stream = false
            suffix = "test."
            temperature = 1
            top_p = 1
            user = "user-1234"
        }