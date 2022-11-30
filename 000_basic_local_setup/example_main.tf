provider "aws" {
  region = "eu-central-1"
}

resource "aws_s3_bucket" "jug-iac-example-bucket" {
  bucket = "jug-hannover-iac-example-bucket"
}

