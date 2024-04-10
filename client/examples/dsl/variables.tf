variable "name" {
  type    = string
  default = "apps"
}

variable "database" {
  type = object({
    engine          = optional(string, "aurora-postgresql")
    engine_version  = optional(string, "14.5")
    master_username = optional(string, "administrator")
  })
  default     = {}
  description = "Configuration for database."
}

variable "eks" {
  type = object({
    // control plane access configuration
    access = optional(object({
      public  = optional(bool, false)
      private = optional(bool, true)
    }), {})
  })
  default     = {}
  description = "Configuration for EKS cluster."
}

variable "network_cidr" {
  type        = string
  default     = "10.0.0.0/16"
  description = "CIDR for core VPC"
}
