terraform {
  required_providers {
    openai = {
      source  = "antonbabenko/openai"
      version = "2.2.1"
    }
  }
}

provider "openai" {
  # Configuration options
}