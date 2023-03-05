//data "aws_ami" "ubuntu" {
//  most_recent = true
//
//  filter {
//    name   = "name"
//    values = ["ubuntu/images/hvm-ssd/ubuntu-focal-20.04-amd64-server-*"]
//  }
//
//  filter {
//    name   = "virtualization-type"
//    values = ["hvm"]
//  }
//
//  owners = ["099720109477"] # Canonical
//}
//
//resource "aws_instance" "web" {
//  ami           = data.aws_ami.ubuntu.id
//  instance_type = "t3.micro"
//  tags = {
//    Name = "HelloWorld"
//  }
//
//  user_data = <<-EOF
//              #!/bin/bash
//              apt-get update
//              apt-get install -y nginx
//              systemctl enable nginx
//              systemctl start nginx
//              EOF
//
//  vpc_security_group_ids = [aws_security_group.web.id]
//
//  # Add a public IP address to the instance
//  # Note: This is not recommended for production use
//  associate_public_ip_address = true
//}
//
//resource "aws_security_group" "web" {
//  name_prefix = "web-"
//
//  ingress {
//    from_port   = 80
//    to_port     = 80
//    protocol    = "tcp"
//    cidr_blocks = ["0.0.0.0/0"]
//  }
//
//  # Allow all outbound traffic
//  egress {
//    from_port   = 0
//    to_port     = 0
//    protocol    = "-1"
//    cidr_blocks = ["0.0.0.0/0"]
//  }
//}
//
//
//
////data "aws_ami" "ubuntu" {
////  most_recent = true
////
////  filter {
////    name   = "name"
////    values = ["ubuntu/images/hvm-ssd/ubuntu-focal-20.04-amd64-server-*"]
////  }
////
////  filter {
////    name   = "virtualization-type"
////    values = ["hvm"]
////  }
////
////  owners = ["099720109477"] # Canonical
////}
////
////resource "aws_instance" "web" {
////  ami           = data.aws_ami.ubuntu.id
////  instance_type = "t3.micro"
////  tags = {
////    Name = "HelloWorld"
////  }
////}