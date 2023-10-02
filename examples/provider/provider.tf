terraform {
  required_providers {
    openai = {
      source  = "antonbabenko/openai"
      version = "1.16.1"
    }
  }
}

provider "openai" {
  # Configuration options
}