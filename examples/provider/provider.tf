terraform {
  required_providers {
    openai = {
      source  = "antonbabenko/openai"
      version = "1.12.6"
    }
  }
}

provider "openai" {
  # Configuration options
}