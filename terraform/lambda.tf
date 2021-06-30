resource "aws_lambda_function" "quizdeck_lab" {

  depends_on = [
    aws_iam_role_policy_attachment.lambda_logs_quizdeck_lab,
    aws_cloudwatch_log_group.quizdeck_lab,
  ]

  function_name = "quizdeck_lab"
  role          = aws_iam_role.quizdeck_lab.name
  image_uri     = "${var.AWS_ACCOUNT_ID}.dkr.ecr.${var.AWS_REGION}.amazonaws.com/${var.ECR_QUIZDECK_LAB_REGISTRY_NAME}:latest"
  runtime       = "go1.x"
  timeout       = 60
}
