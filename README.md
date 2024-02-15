# Fork of Terraform Provider for AWS

Implementing support for aws_msk_scram_secret_association that work in multiple different terraform modules that are currently not available in official AWS terraform Provider

# New Resources

### MSK

- aws_msk_single_scram_secret_association (allows to assign a single secret to a cluster)

## Usage
```hcl
terraform {
  required_providers {
    awscust = {
      version = "x.x.x"
      source  = "jandillenkofer/aws"
    }
  }
}
provider "awscust" {
    region = "eu-west-1"
    profile = "my-account"
}

# cluster and secrets must be defined...

resource "aws_msk_single_scram_secret_association" "secret-association" {
  cluster_arn = data.aws_msk_cluster.cluster.arn
  secret_arn  = aws_secretsmanager_secret.secret.arn
  provider    = awscust
}
```

# Terraform AWS Provider

[![Forums][discuss-badge]][discuss]

[discuss-badge]: https://img.shields.io/badge/discuss-terraform--aws-623CE4.svg?style=flat
[discuss]: https://discuss.hashicorp.com/c/terraform-providers/tf-aws/

The [AWS Provider](https://registry.terraform.io/providers/hashicorp/aws/latest/docs) allows [Terraform](https://terraform.io) to manage [AWS](https://aws.amazon.com) resources.

- [Contributing guide](https://hashicorp.github.io/terraform-provider-aws/)
- [Quarterly development roadmap](ROADMAP.md)
- [FAQ](https://hashicorp.github.io/terraform-provider-aws/faq/)
- [Tutorials](https://learn.hashicorp.com/collections/terraform/aws-get-started)
- [discuss.hashicorp.com](https://discuss.hashicorp.com/c/terraform-providers/tf-aws/)
- [Google Groups](http://groups.google.com/group/terraform-tool)

_**Please note:** We take Terraform's security and our users' trust very seriously. If you believe you have found a security issue in the Terraform AWS Provider, please responsibly disclose it by contacting us at security@hashicorp.com._
