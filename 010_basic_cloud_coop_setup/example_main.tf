provider "aws" {
  region = local.region
}

terraform {
  backend "s3" {
    bucket = "jug-hannover-iac-example-bucket"
    key    = "dev/"
    region = "eu-central-1"
    dynamodb_table = "jug-hannover-iac-terraform-state-lock-dynamo"
  }
}

locals {
  region = "eu-central-1"
}


