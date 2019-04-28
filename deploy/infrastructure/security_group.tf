resource "aws_security_group" "allow_ssh_anywhere" {
  name        = "allow_ssh_anywhere"
  description = "Allow all inbound traffic to ssh"
  vpc_id      = "vpc-7763ea0d" #Aca tomar de AWS directamente
  ingress {
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = [ "0.0.0.0/0" ]
    }
  ingress {
    from_port   = 8080
    to_port     = 8080
    protocol    = "tcp"
    cidr_blocks = [ "0.0.0.0/0"]
    description = "Allow access from any host"
    }
  egress {
    from_port       = 0
    to_port         = 0
    protocol        = "-1"
    cidr_blocks     = ["0.0.0.0/0"]
  }
}
