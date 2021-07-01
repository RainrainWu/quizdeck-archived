resource "aws_iam_role" "lambda_quizdeck" {

  name = "lambda_quizdeck"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "lambda.amazonaws.com"
      },
      "Sid": ""
    }
  ]
}
EOF
}

resource "aws_iam_policy" "lambda_quizdeck" {

  name        = "lambda_quizdeck"
  path        = "/"
  description = "IAM policy for logging from a lambda"

  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "logs:CreateLogStream",
        "logs:PutLogEvents"
      ],
      "Resource": [
        "arn:aws:logs:${var.AWS_REGION}:${var.AWS_ACCOUNT_ID}:log-group:/aws/lambda/*:*"
      ]
    }
  ]
}
EOF
}

resource "aws_iam_role_policy_attachment" "lambda_logs_quizdeck" {

  role       = aws_iam_role.lambda_quizdeck.name
  policy_arn = aws_iam_policy.lambda_quizdeck.arn
}
