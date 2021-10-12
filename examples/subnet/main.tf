terraform {
  required_providers {
    subnet = {
      version = "0.1"
      source  = "github.com/mgnisia/subnet"
    }
  }
}
data "subnet_single" "test" {
  cidr = "10.69.32.0/20"
  ip   = "10.69.36.88"
}

data "subnet_list" "test" {
  cidr_list = ["10.69.32.0/20", "10.75.32.0/20"]
  ip        = "10.69.36.88"
}

output "included" {
  value = data.subnet_list.test.included
}

output "included_subnet" {
  value = data.subnet_list.test.included_subnet_cidr
}