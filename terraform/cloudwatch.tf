resource "aws_cloudwatch_log_group" "quizdeck_lab" {
  name              = "/aws/lambda/${var.LAMBDA_QUIZDECK_LAB_FUNCTION_NAME}"
  retention_in_days = 14
}
