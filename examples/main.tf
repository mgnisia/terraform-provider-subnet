terraform {
  required_providers {
    subnet = {
      version = "0.1"
      source  = "github.com/mgnisia/subnet"
    }
  }
}

provider "subnet" {

}

module "example" {
  source = "./subnet"
}
output "included" {
  value = module.example.included
}
output "included_subnet" {
  value = module.example.included_subnet
}
