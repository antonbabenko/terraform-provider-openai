terraform {
  required_providers {
    openai = {
      source  = "antonbabenko/openai"
      version = "1.8.1"
    }
  }
}

provider "openai" {
  # Configuration options
}