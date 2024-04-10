locals {
  subnet_bits = 4
  subnet_cidrs = {
    database = [
      cidrsubnet(var.network_cidr, local.subnet_bits, 0),
      cidrsubnet(var.network_cidr, local.subnet_bits, 1),
      cidrsubnet(var.network_cidr, local.subnet_bits, 2),
    ]
    private = [
      cidrsubnet(var.network_cidr, local.subnet_bits, 4),
      cidrsubnet(var.network_cidr, local.subnet_bits, 5),
      cidrsubnet(var.network_cidr, local.subnet_bits, 6),
    ]
    public = [
      cidrsubnet(var.network_cidr, local.subnet_bits, 8),
      cidrsubnet(var.network_cidr, local.subnet_bits, 9),
      cidrsubnet(var.network_cidr, local.subnet_bits, 10),
    ]
  }
}

module "network" {
  source = "vpc"

  name             = var.name
  azs              = data.aws_availability_zones.current.names
  cidr             = var.network_cidr
  database_subnets = local.subnet_cidrs.database
  private_subnets  = local.subnet_cidrs.private
  public_subnets   = local.subnet_cidrs.public
}
