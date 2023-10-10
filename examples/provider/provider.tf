terraform {
  required_providers {
    openai = {
      source  = "antonbabenko/openai"
      version = "1.17.2"
    }
  }
}

provider "openai" {
  # Configuration options
}