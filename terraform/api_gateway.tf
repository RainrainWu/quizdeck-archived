resource "aws_apigatewayv2_api" "quizdeck" {
  name          = "quizdeck"
  protocol_type = "HTTP"
}

resource "aws_apigatewayv2_integration" "example" {
  api_id             = aws_apigatewayv2_api.quizdeck.id
  integration_type   = "HTTP_PROXY"
  integration_method = "POST"
  integration_uri    = aws_lambda_function.quizdeck.invoke_arn
}

resource "aws_apigatewayv2_route" "quizdeck" {
  api_id    = aws_apigatewayv2_api.quizdeck.id
  route_key = "ANY /discord/{proxy+}"
  target    = "integrations/${aws_apigatewayv2_integration.quizdeck.id}"
}

resource "aws_apigatewayv2_stage" "quizdeck" {
  api_id = aws_apigatewayv2_api.quizdeck.id
  name   = "develop"
}
