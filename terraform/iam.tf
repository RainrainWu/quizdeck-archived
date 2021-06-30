resource "aws_iam_role" "quizdeck_lab" {

  name = var.LAMBDA_QUIZDECK_LAB_FUNCTION_NAME

  assume_role_policy = <<EOF
{
    "Version": "2021-06-30",
    "Statement": []
}
EOF
}

resource "aws_iam_policy" "quizdeck_lab" {

  name        = var.LAMBDA_QUIZDECK_LAB_FUNCTION_NAME
  path        = "/"
  description = "IAM policy for logging from a lambda"

  policy = <<EOF
{
  "Version": "2021-06-30",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "logs:CreateLogStream",
        "logs:PutLogEvents"
      ],
      "Resource": [
          "arn:aws:logs:${var.AWS_REGION}:${var.AWS_ACCOUNT_ID}:log-group:/aws/lambda/${var.LAMBDA_QUIZDECK_LAB_FUNCTION_NAME}:*"
      ]
    }
  ]
}
EOF
}

resource "aws_iam_role_policy_attachment" "lambda_logs_quizdeck_lab" {

  role       = aws_iam_role.quizdeck_lab.name
  policy_arn = aws_iam_policy.quizdeck_lab.arn
}

