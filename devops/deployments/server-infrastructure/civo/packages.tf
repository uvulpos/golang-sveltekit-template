variable "civo_token" {}
variable "civo_region" {}

terraform {
  required_providers {
    civo = {
      source  = "civo/civo"
      version = "1.0.39"
    }
  }
}

provider "civo" {
  token  = var.civo_token
  region = var.civo_region
}
