variable "environment" {
  type        = string
  description = "Environment name"
  default     = "dev"
}

variable "instance_count" {
  type    = number
  default = 1
}