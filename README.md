# Terraform Provider subnet

Run the following command to build the provider

```shell
go build -o terraform-provider-subnet
```

## Examples

```
data "subnet_single" "test" {
  cidr = "10.69.32.0/20"
  ip   = "10.69.36.88"
}
```

with `data.subnet_single_test.included` you get a boolean which tells you whether the given IP is included in the given subnet

```
data "subnet_list" "test" {
  cidr_list = ["10.69.32.0/20", "10.75.32.0/20"]
  ip        = "10.69.36.88"
}
```

- with `data.subnet_list.included` you get a boolean which tells you whether the given IP is included in the given subnet
- with `data.subnet_list.included_subnet_index` you get the index (int) which of the given subnets in the cidr_list includes the given IP
- with `data.subnet_list.included_subnet_cidr` you get the cidr (string) which of the given subnets in the cidr_list includes the given IP