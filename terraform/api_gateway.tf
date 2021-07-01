resource "aws_apigatewayv2_api" "quizdeck" {
  name          = "quizdeck"
  protocol_type = "HTTP"
}

resource "aws_apigatewayv2_integration" "quizdeck" {
  api_id             = aws_apigatewayv2_api.quizdeck.id
  integration_type   = "AWS_PROXY"
  integration_method = "POST"
  integration_uri    = aws_lambda_function.quizdeck_lab.invoke_arn
}

resource "aws_apigatewayv2_route" "quizdeck" {
  api_id    = aws_apigatewayv2_api.quizdeck.id
  route_key = "POST /lab"
  target    = "integrations/${aws_apigatewayv2_integration.quizdeck.id}"
}

resource "aws_apigatewayv2_stage" "quizdeck" {
  api_id = aws_apigatewayv2_api.quizdeck.id
  name   = "develop"
}
