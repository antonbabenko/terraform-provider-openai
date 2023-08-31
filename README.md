<div align="center">
    <picture>
        <img src="https://github.com/antonbabenko/terraform-provider-openai/assets/108416695/e1d2a17a-0831-41c4-804a-6e162bd33b60" width="500">
    </picture>
   <p>Programmatically control resources in the OpenAI API.</p>
   <a href="https://platform.openai.com/docs/api-reference"><img src="https://img.shields.io/static/v1?label=Docs&message=API Ref&color=000000&style=for-the-badge" /></a>
  <a href="https://opensource.org/licenses/MIT"><img src="https://img.shields.io/badge/License-MIT-blue.svg?style=for-the-badge" /></a>
</div>

## Authentication

Developers will need to create an API Key within their OpenAI [organisation](https://platform.openai.com/) to make API requests.

<!-- Start SDK Installation -->
## SDK Installation

To install this provider, copy and paste this code into your Terraform configuration. Then, run `terraform init`.

```hcl
terraform {
  required_providers {
    openai = {
      source  = "antonbabenko/openai"
      version = "1.15.0"
    }
  }
}

provider "openai" {
  # Configuration options
}
```
<!-- End SDK Installation -->

<!-- Start SDK Example Usage -->
## Testing the provider locally

Should you want to validate a change locally, the `--debug` flag allows you to execute the provider against a terraform instance locally.

This also allows for debuggers (e.g. delve) to be attached to the provider.

```sh
go run main.go --debug
# Copy the TF_REATTACH_PROVIDERS env var
# In a new terminal
cd examples/your-example
TF_REATTACH_PROVIDERS=... terraform init
TF_REATTACH_PROVIDERS=... terraform apply
```
<!-- End SDK Example Usage -->

<!-- Start SDK Available Operations -->

<!-- End SDK Available Operations -->

Terraform allows you to use local provider builds by setting a `dev_overrides` block in a configuration file called `.terraformrc`. This block overrides all other configured installation methods.

Terraform searches for the `.terraformrc` file in your home directory and applies any configuration settings you set.

```
provider_installation {

  dev_overrides {
      "registry.terraform.io/openai/scaffolding" = "<PATH>"
  }

  # For all other providers, install them directly from their origin provider
  # registries as normal. If you omit this, Terraform will _only_ use
  # the dev_overrides block, and so no other providers will be available.
  direct {}
}
```

Your `<PATH>` may vary depending on how your Go environment variables are configured. Execute `go env GOBIN` to set it, then set the `<PATH>` to the value returned. If nothing is returned, set it to the default location, `$HOME/go/bin`.

### Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release!

### Provider Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
