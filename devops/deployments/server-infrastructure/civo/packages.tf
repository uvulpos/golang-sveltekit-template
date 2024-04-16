variable "civo_token" { type = string }
variable "civo_region" { type = string }

terraform {
  backend "s3" {
    endpoints {
      s3 = "https://objectstore.fra1.civo.com/"
    }
    bucket                      = ""
    key                         = ""
    region                      = "FRA1"
    skip_region_validation      = true
    skip_credentials_validation = true
    skip_requesting_account_id  = true
    skip_metadata_api_check     = true
    use_path_style              = true
    # access_key                  = var.tfbackend_objectstorage_accesskey
    # secret_key                  = var.tfbackend_objectstorage_secretkey
  }

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
