resource "aws_apigatewayv2_api" "quizdeck" {

  depends_on = [
    aws_iam_role_policy_attachment.apigateway_invoke_lambda,
  ]

  name          = "quizdeck"
  protocol_type = "HTTP"
}

resource "aws_apigatewayv2_integration" "quizdeck" {

  api_id                 = aws_apigatewayv2_api.quizdeck.id
  integration_type       = "AWS_PROXY"
  integration_method     = "POST"
  integration_uri        = aws_lambda_function.quizdeck_lab.invoke_arn
  payload_format_version = "2.0"
}

resource "aws_apigatewayv2_route" "quizdeck" {

  api_id    = aws_apigatewayv2_api.quizdeck.id
  route_key = "POST /lab"
  target    = "integrations/${aws_apigatewayv2_integration.quizdeck.id}"
}

resource "aws_apigatewayv2_stage" "quizdeck" {

  api_id      = aws_apigatewayv2_api.quizdeck.id
  name        = "$default"
  auto_deploy = true
}
