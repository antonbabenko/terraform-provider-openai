terraform {
  required_providers {
    openai = {
      source  = "antonbabenko/openai"
      version = "2.2.0"
    }
  }
}

provider "openai" {
  # Configuration options
}