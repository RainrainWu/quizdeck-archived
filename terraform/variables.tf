variable "AWS_REGION" {

  type    = string
  default = "us-east-1"
}

variable "AWS_ACCOUNT_ID" {}

variable "ECR_QUIZDECK_LAB_REGISTRY_NAME" {

  type    = string
  default = "quizdeck/lab"
}

variable "LAMBDA_QUIZDECK_LAB_FUNCTION_NAME" {

  type    = string
  default = "quizdeck_lab"
}
