terraform {
  required_providers {
    openai = {
      source  = "antonbabenko/openai"
      version = "1.14.0"
    }
  }
}

provider "openai" {
  # Configuration options
}