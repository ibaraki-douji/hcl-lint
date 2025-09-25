locals {
  common_tags = {
    Environment = var.environment
    Project     = "hcl-lint-test"
  }
}

# This file has an intentional syntax error
resource "aws_s3_bucket" "test" {
  bucket = "test-bucket-${var.environment}"
  
  # Missing closing brace - this should cause a parse error