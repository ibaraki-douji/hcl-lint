resource "aws_instance" "example" {
  ami           = "ami-0abcdef1234567890"
  instance_type = "t3.micro"

  tags = {
    Name = "example-instance"
  }
}