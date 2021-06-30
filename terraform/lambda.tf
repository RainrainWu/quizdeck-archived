resource "aws_lambda_function" "quizdeck_lab" {

  depends_on = [
    aws_iam_role_policy_attachment.lambda_logs_quizdeck_lab,
    aws_cloudwatch_log_group.quizdeck_lab,
  ]

  function_name = "quizdeck_lab"
  role          = aws_iam_role.quizdeck_lab.arn
  image_uri     = "${var.aws_account_id}.dkr.ecr.${var.aws_region}.amazonaws.com/quizdeck_lab:latest"
  runtime       = "go1.x"
  timeout       = 60
}
