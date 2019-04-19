resource "aws_instance" "web" {
  ami           = "ami-0a313d6098716f372" #"ami-0de53d8956e8dcf80"  ubuntu : "ami-0a313d6098716f372"
  instance_type = "t2.micro"
  vpc_security_group_ids =  [ "${aws_security_group.allow_ssh_anywhere.id}" ] #This will be static getting permissions
  key_name = "${aws_key_pair.keypair.key_name}" #to connect ssh use ssh -i ../watchdog ec2-user@ec2-3-91-94-206.compute-1.amazonaws.com
  tags = {
    Name = "HelloWorld2"
  }
}
