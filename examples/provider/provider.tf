terraform {
  required_providers {
    openai = {
      source  = "antonbabenko/openai"
      version = "2.3.2"
    }
  }
}

provider "openai" {
  # Configuration options
}