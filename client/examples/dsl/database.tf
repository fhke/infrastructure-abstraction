module "database" {
  source = "aurora"

  name                 = format("%s-entities", var.name)
  vpc_id               = module.network.vpc_id
  db_subnet_group_name = module.network.database_subnet_group_name
  engine               = var.database.engine
  engine_version       = var.database.engine_version
  master_username      = var.database.master_username
}