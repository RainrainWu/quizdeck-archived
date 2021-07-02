resource "aws_lambda_function" "quizdeck_lab" {

  depends_on = [
    aws_iam_role_policy_attachment.lambda_logs_quizdeck,
    aws_cloudwatch_log_group.quizdeck_lab,
  ]

  function_name = var.LAMBDA_QUIZDECK_LAB_FUNCTION_NAME
  role          = aws_iam_role.lambda_quizdeck.arn
  package_type  = "Image"
  image_uri     = "${var.AWS_ACCOUNT_ID}.dkr.ecr.${var.AWS_REGION}.amazonaws.com/${var.ECR_QUIZDECK_LAB_REGISTRY_NAME}:latest"
  timeout       = 60
}

resource "aws_lambda_permission" "quizdeck_lab" {

  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.quizdeck_lab.function_name
  principal     = "apigateway.amazonaws.com"
  source_arn    = "${aws_apigatewayv2_api.quizdeck.execution_arn}/*/*/lab"
}

resource "aws_lambda_function" "quizdeck_router_discord" {

  depends_on = [
    aws_iam_role_policy_attachment.lambda_logs_quizdeck,
    aws_cloudwatch_log_group.quizdeck_router_discord,
  ]

  function_name = var.LAMBDA_QUIZDECK_ROUTER_DISCORD_FUNCTION_NAME
  role          = aws_iam_role.lambda_quizdeck.arn
  package_type  = "Image"
  image_uri     = "${var.AWS_ACCOUNT_ID}.dkr.ecr.${var.AWS_REGION}.amazonaws.com/${var.ECR_QUIZDECK_ROUTER_DISCORD_REGISTRY_NAME}:latest"
  timeout       = 60
}

resource "aws_lambda_permission" "quizdeck_router_discord" {

  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.quizdeck_router_discord.function_name
  principal     = "apigateway.amazonaws.com"
  source_arn    = "${aws_apigatewayv2_api.quizdeck.execution_arn}/*/*/router_discord"
}
