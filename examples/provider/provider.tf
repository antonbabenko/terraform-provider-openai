terraform {
  required_providers {
    openai = {
      source  = "antonbabenko/openai"
      version = "1.11.1"
    }
  }
}

provider "openai" {
  # Configuration options
}