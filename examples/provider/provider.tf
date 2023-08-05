terraform {
  required_providers {
    openai = {
      source  = "antonbabenko/openai"
      version = "1.9.0"
    }
  }
}

provider "openai" {
  # Configuration options
}