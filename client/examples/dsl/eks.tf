module "cluster" {
  source = "eks"

  cluster_name = var.name
  vpc_id       = module.network.vpc_id
  subnet_ids = [
    // vpc module doesn't provide subnet IDs as output
    for arn in module.network.private_subnet_arns :
    reverse(split("/", arn))[0]
  ]
  cluster_endpoint_public_access  = var.eks.access.public
  cluster_endpoint_private_access = var.eks.access.private
}